//nolint:wrapcheck
package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/docker/go-units"
	"github.com/dustin/go-humanize"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pachyderm/pachyderm/v2/src/client"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/http2"
)

var (
	size   = flag.String("size", "50MB", "size of the file to upload")
	kind   = flag.String("kind", "grpc", "type of benchmark to perform; grpc, http (requires experimental pachd), s3, s3-multipart, console")
	scheme = flag.String("scheme", "http", "url scheme for http-based protocols")
	addr   = flag.String("addr", "localhost", "for http, console, and s3; the address to connect to (always requires a valid pach context, though)")
	h2c    = flag.Bool("h2c", false, "if true, use http2 over cleartext (only works against proxy)")
	iters  = flag.Int("n", 10, "number of times to run the benchmark; reusing connections between tests where allowed in the protocol")
)

type R struct {
	b     byte
	Len   uint64
	EOFAt time.Time
}

func (r *R) Read(p []byte) (int, error) {
	var n int
	if r.Len == 0 {
		fmt.Printf("0\n")
		r.EOFAt = time.Now()
		return 0, io.EOF
	}
	for i := range p {
		if r.Len > 0 {
			p[i] = r.b
			r.Len--
			n++
		} else {
			fmt.Printf(".\n")
			r.EOFAt = time.Now()
			return n, io.EOF
		}
	}
	r.b++
	fmt.Printf(".")
	return n, nil
}

var _ io.Reader = new(R)

func bench(f func(name string, r io.Reader, length uint64) error) error {
	for i := 0; i < *iters; i++ {
		r := new(R)
		var length uint64
		if x, err := humanize.ParseBytes(*size); err != nil {
			log.Fatalf("parse size: %v", err)
		} else {
			r.Len = x
			length = x
		}

		n, err := uuid.NewV4()
		if err != nil {
			panic(err)
		}
		name := n.String()
		start := time.Now()
		if err := f(name, r, length); err != nil {
			return err
		}
		log.Printf("total time: %v", time.Since(start).String())
		log.Printf(" = %s/s", humanize.Bytes(uint64(float64(length)/float64(time.Since(start).Seconds()))))
		log.Printf("read time: %v", time.Since(start)-time.Since(r.EOFAt))
		log.Printf("flush time: %v", time.Since(r.EOFAt).String())
	}
	return nil
}

func main() {
	flag.Parse()

	// Cancel context on first SIGINT.
	ctx, cancel := context.WithCancel(context.Background())
	cancelCh := make(chan os.Signal, 1)
	signal.Notify(cancelCh, os.Interrupt)
	go func() {
		<-cancelCh
		cancel()
		signal.Stop(cancelCh)
		close(cancelCh)
	}()

	// Setup pach context for creating destination repo, etc.
	c, err := client.NewOnUserMachine("")
	if err != nil {
		log.Fatal(err)
	}
	c = c.WithCtx(ctx)
	if err := c.CreateProjectRepo("default", "benchmark-upload"); err != nil {
		log.Printf("create repo: %v", err)
	}

	hc := &http.Client{}
	if *h2c {
		hc = &http.Client{
			Transport: &http2.Transport{
				AllowHTTP: true,
				DialTLSContext: func(ctx context.Context, network, addr string, cfg *tls.Config) (net.Conn, error) {
					return (&net.Dialer{}).DialContext(ctx, network, addr)
				},
			},
		}
	}

	// Run benchmark.
	var benchErr error
	switch *kind {
	case "grpc":
		benchErr = bench(func(name string, r io.Reader, _ uint64) error {
			commit := client.NewProjectCommit("default", "benchmark-upload", "master", "")
			return c.PutFile(commit, name, r)
		})
	case "http":
		benchErr = bench(func(name string, r io.Reader, _ uint64) error {
			req, err := http.NewRequestWithContext(ctx, "PUT", *scheme+"://"+*addr+"/upload/"+name, r)
			if err != nil {
				return err
			}
			res, err := hc.Do(req)
			if err != nil {
				return err
			}
			defer res.Body.Close()

			d, err := httputil.DumpResponse(res, true)
			if err != nil {
				return err
			}
			log.Printf("%s", d)
			return nil
		})
	case "console":
		benchErr = bench(func(name string, r io.Reader, length uint64) error {
			// Start
			start, err := json.Marshal(map[string]any{
				"branch": "master",
				"path":   "/",
				"repo":   "benchmark-upload",
			})
			if err != nil {
				return errors.Wrap(err, "marshal start")
			}
			req, err := http.NewRequestWithContext(ctx, "POST", *scheme+"://"+*addr+"/upload/start", bytes.NewReader(start))
			if err != nil {
				return err
			}
			req.Header.Add("content-type", "application/json")
			req.AddCookie(&http.Cookie{
				Name:  "dashAuthToken",
				Value: c.AuthToken(),
			})
			res, err := hc.Do(req)
			if err != nil {
				return errors.Wrap(err, "start")
			}
			body, err := io.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				return errors.Wrap(err, "read start body")
			}
			if got, want := res.StatusCode, http.StatusOK; got != want {
				return fmt.Errorf("start: unexpected status code: got %d want %d %v", got, want, res.Status)
			}
			var idInfo struct {
				UploadId string `json:"uploadId"`
			}
			if err := json.Unmarshal(body, &idInfo); err != nil {
				return errors.Wrap(err, "unmarshal start reply")
			}

			// Upload
			ckSize := 50 * units.MB
			nCk := int(math.Ceil(float64(length) / float64(ckSize)))
			buf := make([]byte, units.MB)
			for i := 1; i <= nCk; i++ { // chunks start at 1, not 0
				uploadR, uploadW := io.Pipe()
				w := multipart.NewWriter(uploadW)
				go func() {
					defer uploadW.Close()
					defer w.Close()
					_ = w.WriteField("uploadId", idInfo.UploadId)
					_ = w.WriteField("fileName", name)
					_ = w.WriteField("chunkTotal", strconv.FormatInt(int64(nCk), 10))
					_ = w.WriteField("currentChunk", strconv.FormatInt(int64(i), 10))
					fw, _ := w.CreateFormFile("file", name)
					if _, err := io.CopyBuffer(fw, io.LimitReader(r, int64(ckSize)), buf); err != nil {
						uploadW.CloseWithError(fmt.Errorf("copy: %w", err))
					}
				}()

				req, err := http.NewRequestWithContext(ctx, "POST", *scheme+"://"+*addr+"/upload", uploadR)
				if err != nil {
					return err
				}
				req.Header.Add("content-type", w.FormDataContentType())
				req.AddCookie(&http.Cookie{
					Name:  "dashAuthToken",
					Value: c.AuthToken(),
				})

				res, err := hc.Do(req)
				if err != nil {
					return errors.Wrapf(err, "chunk %d/%d", i, nCk)
				}
				res.Body.Close()
				if got, want := res.StatusCode, http.StatusOK; got != want {
					return fmt.Errorf("chunk %d/%d: unexpected status code: got %d want %d %v", i, nCk, got, want, res.Status)
				}
				fmt.Printf(",")
			}
			n, err := r.Read(buf)
			if n != 0 {
				return fmt.Errorf("%d bytes remain unexpectedly", n)
			}
			if err == nil || err != io.EOF {
				return fmt.Errorf("reader finished with unexpected error: %w", err)
			}
			fmt.Printf("\n")

			// Finish
			finish, err := json.Marshal(idInfo) // could reuse "body" from start section
			if err != nil {
				return errors.Wrap(err, "marshal finish request")
			}
			req, err = http.NewRequestWithContext(ctx, "POST", *scheme+"://"+*addr+"/upload/finish", bytes.NewReader(finish))
			if err != nil {
				return err
			}
			req.Header.Add("content-type", "application/json")
			req.AddCookie(&http.Cookie{
				Name:  "dashAuthToken",
				Value: c.AuthToken(),
			})
			res, err = hc.Do(req)
			if err != nil {
				return errors.Wrap(err, "finish")
			}
			res.Body.Close()
			if got, want := res.StatusCode, http.StatusOK; got != want {
				return fmt.Errorf("finish: unexpected status code: got %d want %d %v", got, want, res.Status)
			}
			return nil
		})
	case "s3":
		mc, err := minio.New(*addr, &minio.Options{
			Creds:  credentials.NewStaticV4(c.AuthToken(), c.AuthToken(), ""),
			Secure: *scheme == "https",
		})
		if err != nil {
			log.Fatalf("make minio client: %v", err)
		}
		benchErr = bench(func(name string, r io.Reader, length uint64) error {
			res, err := mc.PutObject(ctx, "master.benchmark-upload", name, r, int64(length), minio.PutObjectOptions{
				DisableMultipart: true,
			})
			log.Printf("minio reply: %#v", res)
			return err
		})
	case "s3-multipart":
		mc, err := minio.New(*addr, &minio.Options{
			Creds:  credentials.NewStaticV4(c.AuthToken(), c.AuthToken(), ""),
			Secure: *scheme == "https",
		})
		if err != nil {
			log.Fatalf("minio: %v", err)
		}
		benchErr = bench(func(name string, r io.Reader, length uint64) error {
			res, err := mc.PutObject(ctx, "master.benchmark-upload", name, r, int64(length), minio.PutObjectOptions{
				DisableMultipart: false,
			})
			log.Printf("minio response: %#v", res)
			return err
		})
	default:
		log.Fatalf("unknown benchmark kind %v", *kind)
	}
	// Exit if benchmark returned an error.
	if benchErr != nil {
		log.Fatal(benchErr)
	}

	// Cleanup.
	if err := c.Close(); err != nil {
		log.Fatalf("close: %v", err)
	}
	cancel()
}

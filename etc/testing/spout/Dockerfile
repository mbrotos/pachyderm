ARG GOVERSION=golang:1.20.6
FROM $GOVERSION
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]

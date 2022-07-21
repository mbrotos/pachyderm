"use strict";(self.webpackChunkjupyterlab_pachyderm=self.webpackChunkjupyterlab_pachyderm||[]).push([[283],{2283:(e,t,r)=>{r.r(t),r.d(t,{Controller:()=>L,FormProvider:()=>F,appendErrors:()=>W,get:()=>g,set:()=>N,useController:()=>R,useFieldArray:()=>ee,useForm:()=>Te,useFormContext:()=>x,useFormState:()=>C,useWatch:()=>M});var s=r(6271),a=e=>"checkbox"===e.type,n=e=>e instanceof Date,i=e=>null==e;const o=e=>"object"==typeof e;var u=e=>!i(e)&&!Array.isArray(e)&&o(e)&&!n(e),l=e=>u(e)&&e.target?a(e.target)?e.target.checked:e.target.value:e,c=(e,t)=>[...e].some((e=>(e=>e.substring(0,e.search(/.\d/))||e)(t)===e)),d=e=>e.filter(Boolean),f=e=>void 0===e,g=(e,t,r)=>{if(u(e)&&t){const s=d(t.split(/[,[\].]+?/)).reduce(((e,t)=>i(e)?e:e[t]),e);return f(s)||s===e?f(e[t])?r:e[t]:s}};const m="blur",y="onChange",b="onSubmit",h="all",_="pattern",p="required";var v=(e,t)=>{const r=Object.assign({},e);return delete r[t],r};const j=s.createContext(null),x=()=>s.useContext(j),F=e=>s.createElement(j.Provider,{value:v(e,"children")},e.children);var V=(e,t,r,s=!0)=>{function a(a){return()=>{if(a in e)return t[a]!==h&&(t[a]=!s||h),r&&(r[a]=!0),e[a]}}const n={};for(const t in e)Object.defineProperty(n,t,{get:a(t)});return n},A=e=>u(e)&&!Object.keys(e).length,O=(e,t,r)=>{const s=v(e,"name");return A(s)||Object.keys(s).length>=Object.keys(t).length||Object.keys(s).find((e=>t[e]===(!r||h)))},w=e=>Array.isArray(e)?e:[e],S=(e,t)=>!e||!t||e===t||w(e).some((e=>e&&(e.startsWith(t)||t.startsWith(e))));const k=e=>{e.current&&(e.current.unsubscribe(),e.current=void 0)};function D(e){const t=s.useRef();(({_subscription:e,props:t})=>{t.disabled?k(e):e.current||(e.current=t.subject.subscribe({next:t.callback}))})({_subscription:t,props:e}),s.useEffect((()=>()=>k(t)),[])}function C(e){const t=x(),{control:r=t.control,disabled:a,name:n}=e||{},[i,o]=s.useState(r._formState),u=s.useRef({isDirty:!1,dirtyFields:!1,touchedFields:!1,isValidating:!1,isValid:!1,errors:!1}),l=s.useRef(n);return l.current=n,D({disabled:a,callback:e=>S(l.current,e.name)&&O(e,u.current)&&o(Object.assign(Object.assign({},r._formState),e)),subject:r._subjects.state}),V(i,r._proxyFormState,u.current,!1)}var E=e=>"string"==typeof e;function B(e,t,r,s){const a=Array.isArray(e);return E(e)?(s&&t.watch.add(e),g(r,e)):a?e.map((e=>(s&&t.watch.add(e),g(r,e)))):(s&&(t.watchAll=!0),r)}var U=e=>/^\w*$/.test(e),T=e=>d(e.replace(/["|']|\]/g,"").split(/\.|\[/));function N(e,t,r){let s=-1;const a=U(t)?[t]:T(t),n=a.length,i=n-1;for(;++s<n;){const t=a[s];let n=r;if(s!==i){const r=e[t];n=u(r)||Array.isArray(r)?r:isNaN(+a[s+1])?{}:[]}e[t]=n,e=e[t]}return e}function M(e){const t=x(),{control:r=t.control,name:a,defaultValue:n,disabled:i}=e||{},o=s.useRef(a);o.current=a,D({disabled:i,subject:r._subjects.watch,callback:e=>{if(S(o.current,e.name)){const e=B(o.current,r._names,r._formValues);c(!u(e)||E(o.current)&&g(r._fields,o.current,{})._f?Array.isArray(e)?[...e]:e:Object.assign({},e))}}});const[l,c]=s.useState(f(n)?r._getWatch(a):n);return s.useEffect((()=>{r._removeUnmounted()})),l}function R(e){const t=x(),{name:r,control:a=t.control,shouldUnregister:n}=e,i=M({control:a,name:r,defaultValue:g(a._formValues,r,g(a._defaultValues,r,e.defaultValue))}),o=C({control:a,name:r});s.useRef(r).current=r;const u=a.register(r,Object.assign(Object.assign({},e.rules),{value:i}));return s.useEffect((()=>{const e=(e,t)=>{const r=g(a._fields,e);r&&(r._f.mount=t)};return e(r,!0),()=>{const t=a._options.shouldUnregister||n;(c(a._names.array,r)?t&&!a._stateFlags.action:t)?a.unregister(r):e(r,!1)}}),[r,a,n]),{field:{onChange:e=>{u.onChange({target:{value:l(e),name:r},type:"change"})},onBlur:()=>{u.onBlur({target:{value:i,name:r},type:m})},name:r,value:i,ref:e=>{const t=g(a._fields,r);e&&t&&e.focus&&(t._f.ref={focus:()=>e.focus(),setCustomValidity:t=>e.setCustomValidity(t),reportValidity:()=>e.reportValidity()})}},formState:o,fieldState:{invalid:!!g(o.errors,r),isDirty:!!g(o.dirtyFields,r),isTouched:!!g(o.touchedFields,r),error:g(o.errors,r)}}}const L=e=>e.render(R(e));var W=(e,t,r,s,a)=>t?Object.assign(Object.assign({},r[e]),{types:Object.assign(Object.assign({},r[e]&&r[e].types?r[e].types:{}),{[s]:a||!0})}):{};const q=(e,t,r)=>{for(const s of r||Object.keys(e)){const r=g(e,s);if(r){const e=r._f,s=v(r,"_f");if(e&&t(e.name)){if(e.ref.focus&&f(e.ref.focus()))break;if(e.refs){e.refs[0].focus();break}}else u(s)&&q(s,t)}}};var I=(e,t,r={})=>r.shouldFocus||f(r.shouldFocus)?r.focusName||`${e}.${f(r.focusIndex)?t:r.focusIndex}.`:"",$=(e,t,r)=>e.map(((e,s)=>{const a=t.current[s];return Object.assign(Object.assign({},e),a?{[r]:a[r]}:{})})),P=()=>{const e="undefined"==typeof performance?Date.now():1e3*performance.now();return"xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g,(t=>{const r=(16*Math.random()+e)%16|0;return("x"==t?r:3&r|8).toString(16)}))},H=(e=[],t)=>e.map((e=>Object.assign(Object.assign({},e[t]?{}:{[t]:P()}),e)));function z(e,t){return[...w(e),...w(t)]}var G=e=>Array.isArray(e)?e.map((()=>{})):void 0;function J(e,t,r){return[...e.slice(0,t),...w(r),...e.slice(t)]}var K=(e,t,r)=>Array.isArray(e)?(f(e[r])&&(e[r]=void 0),e.splice(r,0,e.splice(t,1)[0]),e):[];function Q(e,t){return[...w(t),...w(e)]}var X=(e,t)=>f(t)?[]:function(e,t){let r=0;const s=[...e];for(const e of t)s.splice(e-r,1),r++;return d(s).length?s:[]}(e,w(t).sort(((e,t)=>e-t))),Y=(e,t,r)=>{e[t]=[e[r],e[r]=e[t]][0]},Z=(e,t,r)=>(e[t]=r,e);const ee=e=>{const t=x(),{control:r=t.control,name:a,keyName:n="id",shouldUnregister:i}=e,[o,u]=s.useState(H(r._getFieldArray(a),n)),l=s.useRef(o),c=s.useRef(a),d=s.useRef(!1);c.current=a,l.current=o,r._names.array.add(a),D({callback:({values:e,name:t})=>{t!==c.current&&t||u(H(g(e,c.current),n))},subject:r._subjects.array});const f=s.useCallback((e=>{const t=((e,t)=>e.map(((e={})=>v(e,t))))(e,n);return d.current=!0,N(r._formValues,a,t),u(e),t}),[r,a,n]);return s.useEffect((()=>{if(r._stateFlags.action=!1,r._names.watchAll)r._subjects.state.next({});else for(const e of r._names.watch)if(a.startsWith(e)){r._subjects.state.next({});break}d.current&&r._executeSchema([a]).then((e=>{const t=g(e.errors,a);t&&t.type&&!g(r._formState.errors,a)&&(N(r._formState.errors,a,t),r._subjects.state.next({errors:r._formState.errors}))})),r._subjects.watch.next({name:a,values:r._formValues}),r._names.focus&&q(r._fields,(e=>e.startsWith(r._names.focus))),r._names.focus="",r._proxyFormState.isValid&&r._updateValid()}),[o,a,r,n]),s.useEffect((()=>(!g(r._formValues,a)&&N(r._formValues,a,[]),()=>{(r._options.shouldUnregister||i)&&r.unregister(a)})),[a,r,n,i]),{swap:s.useCallback(((e,t)=>{const s=$(r._getFieldArray(a),l,n);Y(s,e,t),r._updateFieldArray(a,Y,{argA:e,argB:t},f(s),!1)}),[f,a,r,n]),move:s.useCallback(((e,t)=>{const s=$(r._getFieldArray(a),l,n);K(s,e,t),r._updateFieldArray(a,K,{argA:e,argB:t},f(s),!1)}),[f,a,r,n]),prepend:s.useCallback(((e,t)=>{const s=Q($(r._getFieldArray(a),l,n),H(w(e),n));r._updateFieldArray(a,Q,{argA:G(e)},f(s)),r._names.focus=I(a,0,t)}),[f,a,r,n]),append:s.useCallback(((e,t)=>{const s=w(e),i=z($(r._getFieldArray(a),l,n),H(s,n));r._updateFieldArray(a,z,{argA:G(e)},f(i)),r._names.focus=I(a,i.length-s.length,t)}),[f,a,r,n]),remove:s.useCallback((e=>{const t=X($(r._getFieldArray(a),l,n),e);r._updateFieldArray(a,X,{argA:e},f(t))}),[f,a,r,n]),insert:s.useCallback(((e,t,s)=>{const i=J($(r._getFieldArray(a),l,n),e,H(w(t),n));r._updateFieldArray(a,J,{argA:e,argB:G(t)},f(i)),r._names.focus=I(a,e,s)}),[f,a,r,n]),update:s.useCallback(((e,t)=>{const s=$(r._getFieldArray(a),l,n),i=Z(s,e,t);l.current=H(i,n),r._updateFieldArray(a,Z,{argA:e,argB:t},f(l.current),!0,!1)}),[f,a,r,n]),replace:s.useCallback((e=>{const t=H(w(e),n);r._updateFieldArray(a,(()=>t),{},f(t),!0,!1)}),[f,a,r,n]),fields:o}};var te=e=>"function"==typeof e;function re(e){let t;const r=Array.isArray(e);if(e instanceof Date)t=new Date(e);else if(e instanceof Set)t=new Set(e);else{if(!r&&!u(e))return e;t=r?[]:{};for(const r in e){if(te(e[r])){t=e;break}t[r]=re(e[r])}}return t}function se(){let e=[];return{get observers(){return e},next:t=>{for(const r of e)r.next(t)},subscribe:t=>(e.push(t),{unsubscribe:()=>{e=e.filter((e=>e!==t))}}),unsubscribe:()=>{e=[]}}}var ae=e=>i(e)||!o(e);function ne(e,t){if(ae(e)||ae(t))return e===t;if(n(e)&&n(t))return e.getTime()===t.getTime();const r=Object.keys(e),s=Object.keys(t);if(r.length!==s.length)return!1;for(const a of r){const r=e[a];if(!s.includes(a))return!1;if("ref"!==a){const e=t[a];if(n(r)&&n(e)||u(r)&&u(e)||Array.isArray(r)&&Array.isArray(e)?!ne(r,e):r!==e)return!1}}return!0}var ie=e=>({isOnSubmit:!e||e===b,isOnBlur:"onBlur"===e,isOnChange:e===y,isOnAll:e===h,isOnTouch:"onTouched"===e}),oe=e=>"boolean"==typeof e,ue=e=>"file"===e.type,le=e=>e instanceof HTMLElement,ce=e=>"select-multiple"===e.type,de=e=>"radio"===e.type,fe="undefined"!=typeof window&&void 0!==window.HTMLElement&&"undefined"!=typeof document,ge=e=>le(e)&&document.contains(e);function me(e,t){const r=U(t)?[t]:T(t),s=1==r.length?e:function(e,t){const r=t.slice(0,-1).length;let s=0;for(;s<r;)e=f(e)?s++:e[t[s++]];return e}(e,r),a=r[r.length-1];let n;s&&delete s[a];for(let t=0;t<r.slice(0,-1).length;t++){let s,a=-1;const i=r.slice(0,-(t+1)),o=i.length-1;for(t>0&&(n=e);++a<i.length;){const t=i[a];s=s?s[t]:e[t],o===a&&(u(s)&&A(s)||Array.isArray(s)&&!s.filter((e=>u(e)&&!A(e)||oe(e))).length)&&(n?delete n[t]:delete e[t]),n=s}}return e}const ye={value:!1,isValid:!1},be={value:!0,isValid:!0};var he=e=>{if(Array.isArray(e)){if(e.length>1){const t=e.filter((e=>e&&e.checked&&!e.disabled)).map((e=>e.value));return{value:t,isValid:!!t.length}}return e[0].checked&&!e[0].disabled?e[0].attributes&&!f(e[0].attributes.value)?f(e[0].value)||""===e[0].value?be:{value:e[0].value,isValid:!0}:be:ye}return ye},_e=(e,{valueAsNumber:t,valueAsDate:r,setValueAs:s})=>f(e)?e:t?""===e?NaN:+e:r?new Date(e):s?s(e):e;const pe={isValid:!1,value:null};var ve=e=>Array.isArray(e)?e.reduce(((e,t)=>t&&t.checked&&!t.disabled?{isValid:!0,value:t.value}:e),pe):pe;function je(e){const t=e.ref;if(!(e.refs?e.refs.every((e=>e.disabled)):t.disabled))return ue(t)?t.files:de(t)?ve(e.refs).value:ce(t)?[...t.selectedOptions].map((({value:e})=>e)):a(t)?he(e.refs).value:_e(f(t.value)?e.ref.value:t.value,e)}function xe(e,t,r){const s=g(e,r);if(s||U(r))return{error:s,name:r};const a=r.split(".");for(;a.length;){const s=a.join("."),n=g(t,s),i=g(e,s);if(n&&!Array.isArray(n)&&r!==s)return{name:r};if(i&&i.type)return{name:s,error:i};a.pop()}return{name:r}}function Fe(e,t){if(ae(e)||ae(t))return t;for(const r in t){const s=e[r],a=t[r];try{e[r]=u(s)&&u(a)||Array.isArray(s)&&Array.isArray(a)?Fe(s,a):a}catch(e){}}return e}function Ve(e,t,r,s,a){let n=-1;for(;++n<e.length;){for(const s in e[n])Array.isArray(e[n][s])?(!r[n]&&(r[n]={}),r[n][s]=[],Ve(e[n][s],g(t[n]||{},s,[]),r[n][s],r[n],s)):!i(t)&&ne(g(t[n]||{},s),e[n][s])?N(r[n]||{},s):r[n]=Object.assign(Object.assign({},r[n]),{[s]:!0});s&&!r.length&&delete s[a]}return r}var Ae=(e,t,r)=>Fe(Ve(e,t,r.slice(0,e.length)),Ve(t,e,r.slice(0,e.length))),Oe=(e,t)=>!d(g(e,t,[])).length&&me(e,t),we=e=>E(e)||s.isValidElement(e),Se=e=>e instanceof RegExp;function ke(e,t,r="validate"){if(we(e)||Array.isArray(e)&&e.every(we)||oe(e)&&!e)return{type:r,message:we(e)?e:"",ref:t}}var De=e=>u(e)&&!Se(e)?e:{value:e,message:""},Ce=async(e,t,r,s)=>{const{ref:n,refs:o,required:l,maxLength:c,minLength:d,min:f,max:g,pattern:m,validate:y,name:b,valueAsNumber:h,mount:v,disabled:j}=e._f;if(!v||j)return{};const x=o?o[0]:n,F=e=>{s&&x.reportValidity&&(x.setCustomValidity(oe(e)?"":e||" "),x.reportValidity())},V={},O=de(n),w=a(n),S=O||w,k=(h||ue(n))&&!n.value||""===t||Array.isArray(t)&&!t.length,D=W.bind(null,b,r,V),C=(e,t,r,s="maxLength",a="minLength")=>{const i=e?t:r;V[b]=Object.assign({type:e?s:a,message:i,ref:n},D(e?s:a,i))};if(l&&(!S&&(k||i(t))||oe(t)&&!t||w&&!he(o).isValid||O&&!ve(o).isValid)){const{value:e,message:t}=we(l)?{value:!!l,message:l}:De(l);if(e&&(V[b]=Object.assign({type:p,message:t,ref:x},D(p,t)),!r))return F(t),V}if(!(k||i(f)&&i(g))){let e,s;const a=De(g),o=De(f);if(isNaN(t)){const r=n.valueAsDate||new Date(t);E(a.value)&&(e=r>new Date(a.value)),E(o.value)&&(s=r<new Date(o.value))}else{const r=n.valueAsNumber||parseFloat(t);i(a.value)||(e=r>a.value),i(o.value)||(s=r<o.value)}if((e||s)&&(C(!!e,a.message,o.message,"max","min"),!r))return F(V[b].message),V}if((c||d)&&!k&&E(t)){const e=De(c),s=De(d),a=!i(e.value)&&t.length>e.value,n=!i(s.value)&&t.length<s.value;if((a||n)&&(C(a,e.message,s.message),!r))return F(V[b].message),V}if(m&&!k&&E(t)){const{value:e,message:s}=De(m);if(Se(e)&&!t.match(e)&&(V[b]=Object.assign({type:_,message:s,ref:n},D(_,s)),!r))return F(s),V}if(y)if(te(y)){const e=ke(await y(t),x);if(e&&(V[b]=Object.assign(Object.assign({},e),D("validate",e.message)),!r))return F(e.message),V}else if(u(y)){let e={};for(const s in y){if(!A(e)&&!r)break;const a=ke(await y[s](t),x,s);a&&(e=Object.assign(Object.assign({},a),D(s,a.message)),F(a.message),r&&(V[b]=e))}if(!A(e)&&(V[b]=Object.assign({ref:x},e),!r))return V}return F(!0),V};const Ee={mode:b,reValidateMode:y,shouldFocusError:!0},Be="undefined"==typeof window;function Ue(e={}){let t,r=Object.assign(Object.assign({},Ee),e),s={isDirty:!1,isValidating:!1,dirtyFields:{},isSubmitted:!1,submitCount:0,touchedFields:{},isSubmitting:!1,isSubmitSuccessful:!1,isValid:!1,errors:{}},o={},u=r.defaultValues||{},l=r.shouldUnregister?{}:re(u),y={action:!1,mount:!1,watch:!1},b={mount:new Set,unMount:new Set,array:new Set,watch:new Set},_=0,p={};const j={isDirty:!1,dirtyFields:!1,touchedFields:!1,isValidating:!1,isValid:!1,errors:!1},x={watch:se(),array:se(),state:se()},F=ie(r.mode),V=ie(r.reValidateMode),O=r.criteriaMode===h,S=(e,t)=>!t&&(b.watchAll||b.watch.has(e)||b.watch.has((e.match(/\w+/)||[])[0])),k=async e=>{let t=!1;return j.isValid&&(t=r.resolver?A((await M()).errors):await R(o,!0),e||t===s.isValid||(s.isValid=t,x.state.next({isValid:t}))),t},D=(e,t)=>(N(s.errors,e,t),x.state.next({errors:s.errors})),C=(e,t,r)=>{const s=g(o,e);if(s){const a=g(l,e,g(u,e));f(a)||r&&r.defaultChecked||t?N(l,e,t?a:je(s._f)):I(e,a)}y.mount&&k()},U=(e,t,r,a=!0)=>{let n=!1;const i={name:e},o=g(s.touchedFields,e);if(j.isDirty){const e=s.isDirty;s.isDirty=i.isDirty=L(),n=e!==i.isDirty}if(j.dirtyFields&&!r){const r=g(s.dirtyFields,e);ne(g(u,e),t)?me(s.dirtyFields,e):N(s.dirtyFields,e,!0),i.dirtyFields=s.dirtyFields,n=n||r!==g(s.dirtyFields,e)}return r&&!o&&(N(s.touchedFields,e,r),i.touchedFields=s.touchedFields,n=n||j.touchedFields&&o!==r),n&&a&&x.state.next(i),n?i:{}},T=(e,t)=>(N(s.dirtyFields,e,Ae(t,g(u,e,[]),g(s.dirtyFields,e,[]))),Oe(s.dirtyFields,e)),M=async e=>r.resolver?await r.resolver(Object.assign({},l),r.context,((e,t,r,s)=>{const a={};for(const r of e){const e=g(t,r);e&&N(a,r,e._f)}return{criteriaMode:r,names:[...e],fields:a,shouldUseNativeValidation:s}})(e||b.mount,o,r.criteriaMode,r.shouldUseNativeValidation)):{},R=async(e,t,a={valid:!0})=>{for(const n in e){const i=e[n];if(i){const e=i._f,n=v(i,"_f");if(e){const n=await Ce(i,g(l,e.name),O,r.shouldUseNativeValidation);if(n[e.name]&&(a.valid=!1,t))break;t||(n[e.name]?N(s.errors,e.name,n[e.name]):me(s.errors,e.name))}n&&await R(n,t,a)}}return a.valid},L=(e,t)=>(e&&t&&N(l,e,t),!ne(G(),u)),W=(e,t,r)=>{const s=Object.assign({},y.mount?l:f(t)?u:E(e)?{[e]:t}:t);return B(e,b,s,r)},I=(e,t,r={})=>{const s=g(o,e);let n=t;if(s){const r=s._f;r&&(N(l,e,_e(t,r)),n=fe&&le(r.ref)&&i(t)?"":t,ue(r.ref)&&!E(n)?r.ref.files=n:ce(r.ref)?[...r.ref.options].forEach((e=>e.selected=n.includes(e.value))):r.refs?a(r.ref)?r.refs.length>1?r.refs.forEach((e=>e.checked=Array.isArray(n)?!!n.find((t=>t===e.value)):n===e.value)):r.refs[0].checked=!!n:r.refs.forEach((e=>e.checked=e.value===n)):r.ref.value=n)}(r.shouldDirty||r.shouldTouch)&&U(e,n,r.shouldTouch),r.shouldValidate&&z(e)},$=(e,t,r)=>{for(const s in t){const a=t[s],i=`${e}.${s}`,u=g(o,i);!b.array.has(e)&&ae(a)&&(!u||u._f)||n(a)?I(i,a,r):$(i,a,r)}},P=(e,t,r={})=>{const a=g(o,e),n=b.array.has(e);N(l,e,t),n?(x.array.next({name:e,values:l}),(j.isDirty||j.dirtyFields)&&r.shouldDirty&&(T(e,t),x.state.next({name:e,dirtyFields:s.dirtyFields,isDirty:L(e,t)}))):!a||a._f||i(t)?I(e,t,r):$(e,t,r),S(e)&&x.state.next({}),x.watch.next({name:e})},H=async a=>{const n=a.target;let i=n.name;const u=g(o,i);if(u){let d,f;const y=n.type?je(u._f):n.value,b=a.type===m,h=!((c=u._f).mount&&(c.required||c.min||c.max||c.maxLength||c.minLength||c.pattern||c.validate)||r.resolver||g(s.errors,i)||u._f.deps)||((e,t,r,s,a)=>!a.isOnAll&&(!r&&a.isOnTouch?!(t||e):(r?s.isOnBlur:a.isOnBlur)?!e:!(r?s.isOnChange:a.isOnChange)||e))(b,g(s.touchedFields,i),s.isSubmitted,V,F),v=S(i,b);b?u._f.onBlur&&u._f.onBlur(a):u._f.onChange&&u._f.onChange(a),N(l,i,y);const w=U(i,y,b,!1),C=!A(w)||v;if(!b&&x.watch.next({name:i,type:a.type}),h)return C&&x.state.next(Object.assign({name:i},v?{}:w));if(!b&&v&&x.state.next({}),p[i]=(p[i],1),j.isValidating&&x.state.next({isValidating:!0}),r.resolver){const{errors:e}=await M([i]),t=xe(s.errors,o,i),r=xe(e,o,t.name||i);d=r.error,i=r.name,f=A(e)}else d=(await Ce(u,g(l,i),O,r.shouldUseNativeValidation))[i],f=await k(!0);u._f.deps&&z(u._f.deps),(async(r,a,n,i,o)=>{const u=g(s.errors,a),l=j.isValid&&s.isValid!==n;var c,d;if(e.delayError&&i?(t=t||(c=D,d=e.delayError,(...e)=>{clearTimeout(_),_=window.setTimeout((()=>c(...e)),d)}),t(a,i)):(clearTimeout(_),i?N(s.errors,a,i):me(s.errors,a)),((i?!ne(u,i):u)||!A(o)||l)&&!r){const e=Object.assign(Object.assign(Object.assign({},o),l?{isValid:n}:{}),{errors:s.errors,name:a});s=Object.assign(Object.assign({},s),e),x.state.next(e)}p[a]--,j.isValidating&&!p[a]&&(x.state.next({isValidating:!1}),p={})})(!1,i,f,d,w)}var c},z=async(e,t={})=>{let a,n;const i=w(e);if(x.state.next({isValidating:!0}),r.resolver){const t=await(async e=>{const{errors:t}=await M();if(e)for(const r of e){const e=g(t,r);e?N(s.errors,r,e):me(s.errors,r)}else s.errors=t;return t})(f(e)?e:i);a=A(t),n=e?!i.some((e=>g(t,e))):a}else e?(n=(await Promise.all(i.map((async e=>{const t=g(o,e);return await R(t&&t._f?{[e]:t}:t)})))).every(Boolean),(n||s.isValid)&&k()):n=a=await R(o);return x.state.next(Object.assign(Object.assign(Object.assign({},!E(e)||j.isValid&&a!==s.isValid?{}:{name:e}),r.resolver?{isValid:a}:{}),{errors:s.errors,isValidating:!1})),t.shouldFocus&&!n&&q(o,(e=>g(s.errors,e)),e?i:b.mount),n},G=e=>{const t=Object.assign(Object.assign({},u),y.mount?l:{});return f(e)?t:E(e)?g(t,e):e.map((e=>g(t,e)))},J=(e,t={})=>{for(const a of e?w(e):b.mount)b.mount.delete(a),b.array.delete(a),g(o,a)&&(t.keepValue||(me(o,a),me(l,a)),!t.keepError&&me(s.errors,a),!t.keepDirty&&me(s.dirtyFields,a),!t.keepTouched&&me(s.touchedFields,a),!r.shouldUnregister&&!t.keepDefaultValue&&me(u,a));x.watch.next({}),x.state.next(Object.assign(Object.assign({},s),t.keepDirty?{isDirty:L()}:{})),!t.keepIsValid&&k()},K=(e,t={})=>{const s=g(o,e);return N(o,e,{_f:Object.assign(Object.assign(Object.assign({},s&&s._f?s._f:{ref:{name:e}}),{name:e,mount:!0}),t)}),b.mount.add(e),!f(t.value)&&!t.disabled&&N(l,e,g(l,e,t.value)),s?oe(t.disabled)&&N(l,e,t.disabled?void 0:g(l,e,je(s._f))):C(e,!0),Be?{name:e}:Object.assign(Object.assign({name:e},oe(t.disabled)?{disabled:t.disabled}:{}),{onChange:H,onBlur:H,ref:s=>{if(s){K(e,t);let r=g(o,e);const n=f(s.value)&&s.querySelectorAll&&s.querySelectorAll("input,select,textarea")[0]||s,i=(e=>de(e)||a(e))(n);if(n===r._f.ref||i&&d(r._f.refs||[]).find((e=>e===n)))return;r={_f:i?Object.assign(Object.assign({},r._f),{refs:[...d(r._f.refs||[]).filter(ge),n],ref:{type:n.type,name:e}}):Object.assign(Object.assign({},r._f),{ref:n})},N(o,e,r),(!t||!t.disabled)&&C(e,!1,n)}else{const s=g(o,e,{}),a=r.shouldUnregister||t.shouldUnregister;s._f&&(s._f.mount=!1),a&&(!c(b.array,e)||!y.action)&&b.unMount.add(e)}}})};return{control:{register:K,unregister:J,_executeSchema:M,_getWatch:W,_getDirty:L,_updateValid:k,_removeUnmounted:()=>{for(const e of b.unMount){const t=g(o,e);t&&(t._f.refs?t._f.refs.every((e=>!ge(e))):!ge(t._f.ref))&&J(e)}b.unMount=new Set},_updateFieldArray:(e,t,r,a=[],n=!0,i=!0)=>{if(y.action=!0,i&&g(o,e)){const s=t(g(o,e),r.argA,r.argB);n&&N(o,e,s)}if(Array.isArray(g(s.errors,e))){const a=t(g(s.errors,e),r.argA,r.argB);n&&N(s.errors,e,a),Oe(s.errors,e)}if(j.touchedFields&&g(s.touchedFields,e)){const a=t(g(s.touchedFields,e),r.argA,r.argB);n&&N(s.touchedFields,e,a),Oe(s.touchedFields,e)}(j.dirtyFields||j.isDirty)&&T(e,a),x.state.next({isDirty:L(e,a),dirtyFields:s.dirtyFields,errors:s.errors,isValid:s.isValid})},_getFieldArray:t=>g(y.mount?l:u,t,e.shouldUnregister?g(u,t,[]):[]),_subjects:x,_proxyFormState:j,get _fields(){return o},set _fields(e){o=e},get _formValues(){return l},set _formValues(e){l=e},get _stateFlags(){return y},set _stateFlags(e){y=e},get _defaultValues(){return u},set _defaultValues(e){u=e},get _names(){return b},set _names(e){b=e},get _formState(){return s},set _formState(e){s=e},get _options(){return r},set _options(e){r=Object.assign(Object.assign({},r),e)}},trigger:z,register:K,handleSubmit:(e,t)=>async a=>{a&&(a.preventDefault&&a.preventDefault(),a.persist&&a.persist());let n=!0,i=Object.assign({},l);x.state.next({isSubmitting:!0});try{if(r.resolver){const{errors:e,values:t}=await M();s.errors=e,i=t}else await R(o);A(s.errors)&&Object.keys(s.errors).every((e=>g(i,e)))?(x.state.next({errors:{},isSubmitting:!0}),await e(i,a)):(t&&await t(s.errors,a),r.shouldFocusError&&q(o,(e=>g(s.errors,e)),b.mount))}catch(e){throw n=!1,e}finally{s.isSubmitted=!0,x.state.next({isSubmitted:!0,isSubmitting:!1,isSubmitSuccessful:A(s.errors)&&n,submitCount:s.submitCount+1,errors:s.errors})}},watch:(e,t)=>te(e)?x.watch.subscribe({next:r=>e(W(void 0,t),r)}):W(e,t,!0),setValue:P,getValues:G,reset:(t,r={})=>{const a=t||u,n=re(a),i=A(t)?u:n;if(r.keepDefaultValues||(u=a),!r.keepValues){if(fe)for(const e of b.mount){const t=g(o,e);if(t&&t._f){const e=Array.isArray(t._f.refs)?t._f.refs[0]:t._f.ref;try{le(e)&&e.closest("form").reset();break}catch(e){}}}l=e.shouldUnregister?r.keepDefaultValues?re(u):{}:n,o={},x.watch.next({values:i}),x.array.next({values:i})}b={mount:new Set,unMount:new Set,array:new Set,watch:new Set,watchAll:!1,focus:""},x.state.next({submitCount:r.keepSubmitCount?s.submitCount:0,isDirty:r.keepDirty?s.isDirty:!!r.keepDefaultValues&&!ne(t,u),isSubmitted:!!r.keepIsSubmitted&&s.isSubmitted,dirtyFields:r.keepDirty?s.dirtyFields:r.keepDefaultValues&&t?Object.entries(t).reduce(((e,[t,r])=>Object.assign(Object.assign({},e),{[t]:r!==g(u,t)})),{}):{},touchedFields:r.keepTouched?s.touchedFields:{},errors:r.keepErrors?s.errors:{},isSubmitting:!1,isSubmitSuccessful:!1}),y.mount=!j.isValid||!!r.keepIsValid,y.watch=!!e.shouldUnregister},resetField:(e,t={})=>{f(t.defaultValue)?P(e,g(u,e)):(P(e,t.defaultValue),N(u,e,t.defaultValue)),t.keepTouched||me(s.touchedFields,e),t.keepDirty||(me(s.dirtyFields,e),s.isDirty=t.defaultValue?L(e,g(u,e)):L()),t.keepError||(me(s.errors,e),j.isValid&&k()),x.state.next(Object.assign({},s))},clearErrors:e=>{e?w(e).forEach((e=>me(s.errors,e))):s.errors={},x.state.next({errors:s.errors,isValid:!0})},unregister:J,setError:(e,t,r)=>{const a=(g(o,e,{_f:{}})._f||{}).ref;N(s.errors,e,Object.assign(Object.assign({},t),{ref:a})),x.state.next({name:e,errors:s.errors,isValid:!1}),r&&r.shouldFocus&&a&&a.focus&&a.focus()},setFocus:e=>{const t=g(o,e)._f;(t.ref.focus?t.ref:t.refs[0]).focus()}}}function Te(e={}){const t=s.useRef(),[r,a]=s.useState({isDirty:!1,isValidating:!1,dirtyFields:{},isSubmitted:!1,submitCount:0,touchedFields:{},isSubmitting:!1,isSubmitSuccessful:!1,isValid:!1,errors:{}});t.current?t.current.control._options=e:t.current=Object.assign(Object.assign({},Ue(e)),{formState:r});const n=t.current.control;return D({subject:n._subjects.state,callback:e=>{O(e,n._proxyFormState,!0)&&(n._formState=Object.assign(Object.assign({},n._formState),e),a(Object.assign({},n._formState)))}}),s.useEffect((()=>{n._stateFlags.mount||(n._proxyFormState.isValid&&n._updateValid(),n._stateFlags.mount=!0),n._stateFlags.watch&&(n._stateFlags.watch=!1,n._subjects.state.next({})),n._removeUnmounted()})),s.useEffect((()=>()=>Object.values(n._subjects).forEach((e=>e.unsubscribe()))),[n]),t.current.formState=V(r,n._proxyFormState),t.current}}}]);
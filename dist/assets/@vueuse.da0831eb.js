import{r as R,s as C,w as L,b as c,e as O,g as B,f as M,h as N,i as V,n as A,u as y}from"./@vue.1c15d20b.js";var H=Object.defineProperty,K=Object.defineProperties,q=Object.getOwnPropertyDescriptors,S=Object.getOwnPropertySymbols,G=Object.prototype.hasOwnProperty,J=Object.prototype.propertyIsEnumerable,j=(e,t,r)=>t in e?H(e,t,{enumerable:!0,configurable:!0,writable:!0,value:r}):e[t]=r,Q=(e,t)=>{for(var r in t||(t={}))G.call(t,r)&&j(e,r,t[r]);if(S)for(var r of S(t))J.call(t,r)&&j(e,r,t[r]);return e},U=(e,t)=>K(e,q(t));function se(e,t){var r;const n=C();return L(()=>{n.value=e()},U(Q({},t),{flush:(r=t==null?void 0:t.flush)!=null?r:"sync"})),R(n)}function E(e){return B()?(M(e),!0):!1}const P=typeof window!="undefined",ue=e=>typeof e=="boolean",ae=e=>typeof e=="number",X=e=>typeof e=="string",T=()=>{};function z(e,t){function r(...n){e(()=>t.apply(this,n),{fn:t,thisArg:this,args:n})}return r}function Y(e,t={}){let r,n;return i=>{const s=y(e),u=y(t.maxWait);if(r&&clearTimeout(r),s<=0||u!==void 0&&u<=0)return n&&(clearTimeout(n),n=null),i();u&&!n&&(n=setTimeout(()=>{r&&clearTimeout(r),n=null,i()},u)),r=setTimeout(()=>{n&&clearTimeout(n),n=null,i()},s)}}function Z(e,t=!0,r=!0){let n=0,o,i=!r;const s=()=>{o&&(clearTimeout(o),o=void 0)};return a=>{const f=y(e),d=Date.now()-n;if(s(),f<=0)return n=Date.now(),a();d>f&&(n=Date.now(),i?i=!1:a()),t&&(o=setTimeout(()=>{n=Date.now(),r||(i=!0),s(),a()},f)),!r&&!o&&(o=setTimeout(()=>i=!0,f))}}function k(e,t=200,r={}){return z(Y(t,r),e)}function le(e,t=200,r={}){if(t<=0)return e;const n=c(e.value),o=k(()=>{n.value=e.value},t,r);return O(e,()=>o()),n}function ce(e,t=200,r=!0,n=!0){return z(Z(t,r,n),e)}function I(e,t=!0){N()?V(e):t?e():A(e)}function fe(e,t,r={}){const{immediate:n=!0}=r,o=c(!1);let i=null;function s(){i&&(clearTimeout(i),i=null)}function u(){o.value=!1,s()}function a(...f){s(),o.value=!0,i=setTimeout(()=>{o.value=!1,i=null,e(...f)},y(t))}return n&&(o.value=!0,P&&a()),E(u),{isPending:o,start:a,stop:u}}function b(e){var t;const r=y(e);return(t=r==null?void 0:r.$el)!=null?t:r}const g=P?window:void 0,ee=P?window.document:void 0;function v(...e){let t,r,n,o;if(X(e[0])?([r,n,o]=e,t=g):[t,r,n,o]=e,!t)return T;let i=T;const s=O(()=>b(t),a=>{i(),a&&(a.addEventListener(r,n,o),i=()=>{a.removeEventListener(r,n,o),i=T})},{immediate:!0,flush:"post"}),u=()=>{s(),i()};return E(u),u}function pe(e,t,r={}){const{window:n=g,ignore:o,capture:i=!0}=r;if(!n)return;const s=c(!0);let u;const a=l=>{n.clearTimeout(u);const p=b(e),_=l.composedPath();!p||p===l.target||_.includes(p)||!s.value||o&&o.length>0&&o.some(m=>{const w=b(m);return w&&(l.target===w||_.includes(w))})||t(l)},f=[v(n,"click",a,{passive:!0,capture:i}),v(n,"pointerdown",l=>{const p=b(e);s.value=!!p&&!l.composedPath().includes(p)},{passive:!0}),v(n,"pointerup",l=>{u=n.setTimeout(()=>a(l),50)},{passive:!0})];return()=>f.forEach(l=>l())}const $=typeof globalThis!="undefined"?globalThis:typeof window!="undefined"?window:typeof global!="undefined"?global:typeof self!="undefined"?self:{},D="__vueuse_ssr_handlers__";$[D]=$[D]||{};$[D];function de({document:e=ee}={}){if(!e)return c("visible");const t=c(e.visibilityState);return v(e,"visibilitychange",()=>{t.value=e.visibilityState}),t}var x=Object.getOwnPropertySymbols,te=Object.prototype.hasOwnProperty,ne=Object.prototype.propertyIsEnumerable,re=(e,t)=>{var r={};for(var n in e)te.call(e,n)&&t.indexOf(n)<0&&(r[n]=e[n]);if(e!=null&&x)for(var n of x(e))t.indexOf(n)<0&&ne.call(e,n)&&(r[n]=e[n]);return r};function oe(e,t,r={}){const n=r,{window:o=g}=n,i=re(n,["window"]);let s;const u=o&&"ResizeObserver"in o,a=()=>{s&&(s.disconnect(),s=void 0)},f=O(()=>b(e),l=>{a(),u&&o&&l&&(s=new ResizeObserver(t),s.observe(l,i))},{immediate:!0,flush:"post"}),d=()=>{a(),f()};return E(d),{isSupported:u,stop:d}}function ve(e,t={}){const{reset:r=!0,windowResize:n=!0,windowScroll:o=!0,immediate:i=!0}=t,s=c(0),u=c(0),a=c(0),f=c(0),d=c(0),l=c(0),p=c(0),_=c(0);function m(){const w=b(e);if(!w){r&&(s.value=0,u.value=0,a.value=0,f.value=0,d.value=0,l.value=0,p.value=0,_.value=0);return}const h=w.getBoundingClientRect();s.value=h.height,u.value=h.bottom,a.value=h.left,f.value=h.right,d.value=h.top,l.value=h.width,p.value=h.x,_.value=h.y}return oe(e,m),O(()=>b(e),w=>!w&&m()),o&&v("scroll",m,{passive:!0}),n&&v("resize",m,{passive:!0}),I(()=>{i&&m()}),{height:s,bottom:u,left:a,right:f,top:d,width:l,x:p,y:_,update:m}}var F,W;P&&(window==null?void 0:window.navigator)&&((F=window==null?void 0:window.navigator)==null?void 0:F.platform)&&/iP(ad|hone|od)/.test((W=window==null?void 0:window.navigator)==null?void 0:W.platform);function me({window:e=g}={}){if(!e)return c(!1);const t=c(e.document.hasFocus());return v(e,"blur",()=>{t.value=!1}),v(e,"focus",()=>{t.value=!0}),t}function we({window:e=g,initialWidth:t=1/0,initialHeight:r=1/0}={}){const n=c(t),o=c(r),i=()=>{e&&(n.value=e.innerWidth,o.value=e.innerHeight)};return i(),I(i),v("resize",i,{passive:!0}),{width:n,height:o}}export{ae as a,ue as b,b as c,we as d,ve as e,oe as f,ce as g,fe as h,P as i,de as j,me as k,se as l,pe as o,le as r,E as t,v as u};

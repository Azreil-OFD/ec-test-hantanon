var A=Object.defineProperty;var E=(s,r,a)=>r in s?A(s,r,{enumerable:!0,configurable:!0,writable:!0,value:a}):s[r]=a;var $=(s,r,a)=>E(s,typeof r!="symbol"?r+"":r,a);import{d as k,r as W,v as z,x as d,J as w,K as g,L as m,M as D,t as I,_ as N}from"./fXgHqc7c.js";var P=/\{[^{}]+\}/g,x=({allowReserved:s,name:r,value:a})=>{if(a==null)return"";if(typeof a=="object")throw new Error("Deeply-nested arrays/objects aren’t supported. Provide your own `querySerializer()` to handle these.");return`${r}=${s?a:encodeURIComponent(a)}`},J=s=>{switch(s){case"label":return".";case"matrix":return";";case"simple":return",";default:return"&"}},V=s=>{switch(s){case"form":return",";case"pipeDelimited":return"|";case"spaceDelimited":return"%20";default:return","}},B=s=>{switch(s){case"label":return".";case"matrix":return";";case"simple":return",";default:return"&"}},U=({allowReserved:s,explode:r,name:a,style:o,value:e})=>{if(!r){let l=(s?e:e.map(i=>encodeURIComponent(i))).join(V(o));switch(o){case"label":return`.${l}`;case"matrix":return`;${a}=${l}`;case"simple":return l;default:return`${a}=${l}`}}let n=J(o),t=e.map(l=>o==="label"||o==="simple"?s?l:encodeURIComponent(l):x({allowReserved:s,name:a,value:l})).join(n);return o==="label"||o==="matrix"?n+t:t},q=({allowReserved:s,explode:r,name:a,style:o,value:e})=>{if(e instanceof Date)return`${a}=${e.toISOString()}`;if(o!=="deepObject"&&!r){let l=[];Object.entries(e).forEach(([c,f])=>{l=[...l,c,s?f:encodeURIComponent(f)]});let i=l.join(",");switch(o){case"form":return`${a}=${i}`;case"label":return`.${i}`;case"matrix":return`;${a}=${i}`;default:return i}}let n=B(o),t=Object.entries(e).map(([l,i])=>x({allowReserved:s,name:o==="deepObject"?`${a}[${l}]`:l,value:i})).join(n);return o==="label"||o==="matrix"?n+t:t},H=({path:s,url:r})=>{let a=r,o=r.match(P);if(o)for(let e of o){let n=!1,t=e.substring(1,e.length-1),l="simple";t.endsWith("*")&&(n=!0,t=t.substring(0,t.length-1)),t.startsWith(".")?(t=t.substring(1),l="label"):t.startsWith(";")&&(t=t.substring(1),l="matrix");let i=s[t];if(i==null)continue;if(Array.isArray(i)){a=a.replace(e,U({explode:n,name:t,style:l,value:i}));continue}if(typeof i=="object"){a=a.replace(e,q({explode:n,name:t,style:l,value:i}));continue}if(l==="matrix"){a=a.replace(e,`;${x({name:t,value:i})}`);continue}let c=encodeURIComponent(l==="label"?`.${i}`:i);a=a.replace(e,c)}return a},O=({allowReserved:s,array:r,object:a}={})=>o=>{let e=[];if(o&&typeof o=="object")for(let n in o){let t=o[n];if(t!=null){if(Array.isArray(t)){e=[...e,U({allowReserved:s,explode:!0,name:n,style:"form",value:t,...r})];continue}if(typeof t=="object"){e=[...e,q({allowReserved:s,explode:!0,name:n,style:"deepObject",value:t,...a})];continue}e=[...e,x({allowReserved:s,name:n,value:t})]}}return e.join("&")},L=s=>{if(!s)return;let r=s.split(";")[0].trim();if(r.startsWith("application/json")||r.endsWith("+json"))return"json";if(r==="multipart/form-data")return"formData";if(["application/","audio/","image/","video/"].some(a=>r.startsWith(a)))return"blob";if(r.startsWith("text/"))return"text"},M=({baseUrl:s,path:r,query:a,querySerializer:o,url:e})=>{let n=e.startsWith("/")?e:`/${e}`,t=s+n;r&&(t=H({path:r,url:t}));let l=a?o(a):"";return l.startsWith("?")&&(l=l.substring(1)),l&&(t+=`?${l}`),t},S=(s,r)=>{var o;let a={...s,...r};return(o=a.baseUrl)!=null&&o.endsWith("/")&&(a.baseUrl=a.baseUrl.substring(0,a.baseUrl.length-1)),a.headers=C(s.headers,r.headers),a},C=(...s)=>{let r=new Headers;for(let a of s){if(!a||typeof a!="object")continue;let o=a instanceof Headers?a.entries():Object.entries(a);for(let[e,n]of o)if(n===null)r.delete(e);else if(Array.isArray(n))for(let t of n)r.append(e,t);else n!==void 0&&r.set(e,typeof n=="object"?JSON.stringify(n):n)}return r},j=class{constructor(){$(this,"_fns");this._fns=[]}clear(){this._fns=[]}exists(s){return this._fns.indexOf(s)!==-1}eject(s){let r=this._fns.indexOf(s);r!==-1&&(this._fns=[...this._fns.slice(0,r),...this._fns.slice(r+1)])}use(s){this._fns=[...this._fns,s]}},G=()=>({error:new j,request:new j,response:new j}),K={bodySerializer:s=>JSON.stringify(s)},F=O({allowReserved:!1,array:{explode:!0,style:"form"},object:{explode:!0,style:"deepObject"}}),Q={"Content-Type":"application/json"},R=(s={})=>({...K,baseUrl:"",fetch:globalThis.fetch,headers:Q,parseAs:"auto",querySerializer:F,...s}),X=(s={})=>{let r=S(R(),s),a=()=>({...r}),o=t=>(r=S(r,t),a()),e=G(),n=async t=>{let l={...r,...t,headers:C(r.headers,t.headers)};l.body&&l.bodySerializer&&(l.body=l.bodySerializer(l.body)),l.body||l.headers.delete("Content-Type");let i=M({baseUrl:l.baseUrl??"",path:l.path,query:l.query,querySerializer:typeof l.querySerializer=="function"?l.querySerializer:O(l.querySerializer),url:l.url}),c={redirect:"follow",...l},f=new Request(i,c);for(let p of e.request._fns)f=await p(f,l);let T=l.fetch,u=await T(f);for(let p of e.response._fns)u=await p(u,f,l);let y={request:f,response:u};if(u.ok){if(u.status===204||u.headers.get("Content-Length")==="0")return{data:{},...y};if(l.parseAs==="stream")return{data:u.body,...y};let p=(l.parseAs==="auto"?L(u.headers.get("Content-Type")):l.parseAs)??"json",v=await u[p]();return p==="json"&&l.responseTransformer&&(v=await l.responseTransformer(v)),{data:v,...y}}let h=await u.text();try{h=JSON.parse(h)}catch{}let b=h;for(let p of e.error._fns)b=await p(h,u,f,l);if(b=b||{},l.throwOnError)throw b;return{error:b,...y}};return{connect:t=>n({...t,method:"CONNECT"}),delete:t=>n({...t,method:"DELETE"}),get:t=>n({...t,method:"GET"}),getConfig:a,head:t=>n({...t,method:"HEAD"}),interceptors:e,options:t=>n({...t,method:"OPTIONS"}),patch:t=>n({...t,method:"PATCH"}),post:t=>n({...t,method:"POST"}),put:t=>n({...t,method:"PUT"}),request:n,setConfig:o,trace:t=>n({...t,method:"TRACE"})}};const _=X(R()),Y=s=>((s==null?void 0:s.client)??_).post({...s,url:"/api/register"}),Z={class:"flex flex-col items-center justify-center min-h-screen bg-gray-100"},ee={class:"mb-4"},te={class:"mb-4"},re={class:"mb-4"},se={class:"mb-4"},le=k({__name:"index",setup(s){_.setConfig({baseUrl:"http://5.23.53.194:8000"});const r=W({login:"",password:"",email:"",full_name:""}),a=async()=>{try{const o=await Y({client:_,body:r.value});console.log(o)}catch(o){console.error("Ошибка регистрации:",o)}};return(o,e)=>(I(),z("div",Z,[e[9]||(e[9]=d("h1",{class:"text-3xl font-bold mb-6"},"Регистрация",-1)),d("form",{onSubmit:D(a,["prevent"]),class:"bg-white p-6 rounded shadow-md w-80"},[d("div",ee,[e[4]||(e[4]=d("label",{for:"login",class:"block text-sm font-medium text-gray-700"},"Логин",-1)),w(d("input",{"onUpdate:modelValue":e[0]||(e[0]=n=>m(r).login=n),type:"text",id:"login",required:"",class:"mt-1 block w-full p-2 border border-gray-300 rounded"},null,512),[[g,m(r).login]])]),d("div",te,[e[5]||(e[5]=d("label",{for:"email",class:"block text-sm font-medium text-gray-700"},"Email",-1)),w(d("input",{"onUpdate:modelValue":e[1]||(e[1]=n=>m(r).email=n),type:"email",id:"email",required:"",class:"mt-1 block w-full p-2 border border-gray-300 rounded"},null,512),[[g,m(r).email]])]),d("div",re,[e[6]||(e[6]=d("label",{for:"password",class:"block text-sm font-medium text-gray-700"},"Пароль",-1)),w(d("input",{"onUpdate:modelValue":e[2]||(e[2]=n=>m(r).password=n),type:"password",id:"password",required:"",class:"mt-1 block w-full p-2 border border-gray-300 rounded"},null,512),[[g,m(r).password]])]),d("div",se,[e[7]||(e[7]=d("label",{for:"full_name",class:"block text-sm font-medium text-gray-700"},"Полное имя",-1)),w(d("input",{"onUpdate:modelValue":e[3]||(e[3]=n=>m(r).full_name=n),type:"text",id:"full_name",required:"",class:"mt-1 block w-full p-2 border border-gray-300 rounded"},null,512),[[g,m(r).full_name]])]),e[8]||(e[8]=d("button",{type:"submit",class:"w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600"},"Зарегистрироваться",-1))],32)]))}}),oe=N(le,[["__scopeId","data-v-2d0a57d2"]]);export{oe as default};
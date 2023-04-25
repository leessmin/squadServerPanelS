import{d as L,k as y,w,G as I,o as S,e as v,f as i,g as d,h as _,F as N,B as O,q as b,H as D,c as x,i as h,t as q,s as V,v as G,b as U}from"./index-2df293b2.js";import{a as $,u as K}from"./alert-95b55714.js";import{_ as k}from"./_plugin-vue_export-helper-c27b6911.js";import"./http-9bd06147.js";import"./config-9c65774d.js";function C(p){const u=p.children;let t=new Array;for(let a=0;a<u.length;a++)t.push(u[a].children[0].textContent);return t}const T=p=>(V("data-v-42840746"),p=p(),G(),p),z={class:"btn-box"},E={class:"map"},F={class:"left"},R=T(()=>i("h3",null,"已选择的地图",-1)),j={class:"right"},P=T(()=>i("h3",null,"可选择的地图",-1)),Q=L({__name:"MapComponents",props:{mapArr:null},emits:["saveMap"],setup(p,{emit:u}){const t=p;let a=y(JSON.parse(JSON.stringify(t.mapArr.get("selected")))),r=y(JSON.parse(JSON.stringify(t.mapArr.get("select"))));w(t,async(e,o)=>{a.value=[],r.value=[],await D(),a.value=JSON.parse(JSON.stringify(e.mapArr.get("selected"))),r.value=JSON.parse(JSON.stringify(e.mapArr.get("select")))});const g=document.createElement("li");g.setAttribute("id","tempDom"),g.style.cssText="width: 100px; height: 30px; border: 3px dotted #ccc; box-sizing: border-box;";const m=e=>{var o;e.preventDefault(),(e.target.className.includes("square")||e.target.className.includes("box"))&&(e.target.className==="square"?(o=e.target.parentNode)==null||o.insertBefore(g,e.target):e.target.className==="box"&&e.target.appendChild(g))},n=e=>{e.target.classList.add("hide")},s=e=>{var c;const o=document.querySelector("#tempDom");(c=o==null?void 0:o.parentNode)==null||c.replaceChild(e.target,o),e.target.classList.remove("hide")},l=y(),f=()=>{u("saveMap",C(l.value))},M=async()=>{console.log("取消"),a.value=[],r.value=[],console.log(JSON.parse(JSON.stringify(t.mapArr.get("selected")))),await D(),a.value=JSON.parse(JSON.stringify(t.mapArr.get("selected"))),r.value=JSON.parse(JSON.stringify(t.mapArr.get("select")))};I(async(e,o,c)=>{if(!H()){c();return}if(!await $()){c(o);return}c()});function H(){return JSON.stringify(C(l.value))!=JSON.stringify(t.mapArr.get("selected"))}return(e,o)=>{const c=x("a-button"),J=x("a-tag");return S(),v(N,null,[i("div",z,[d(c,{type:"primary",onClick:f},{default:_(()=>[h("保存")]),_:1}),d(c,{type:"primary",danger:"",onClick:M},{default:_(()=>[h("取消")]),_:1})]),i("div",E,[i("div",F,[R,i("ul",{onDragover:m,class:"box",ref_key:"mapUlDom",ref:l},[(S(!0),v(N,null,O(b(a),(A,B)=>(S(),v("li",{onDragend:s,onDragstart:n,class:"square",draggable:"true"},[d(J,null,{default:_(()=>[h(q(A),1)]),_:2},1024)],32))),256))],544)]),i("div",j,[P,i("ul",{onDragover:m,class:"box"},[(S(!0),v(N,null,O(b(r),(A,B)=>(S(),v("li",{onDragend:s,onDragstart:n,class:"square",draggable:"true"},[d(J,null,{default:_(()=>[h(q(A),1)]),_:2},1024)],32))),256))],32)])])],64)}}});const W=k(Q,[["__scopeId","data-v-42840746"]]),X={class:"map-box"},Y=L({__name:"GSMap",setup(p){const u=U();let t=y(u.isTheme);u.$subscribe((n,s)=>{console.log(s.isTheme),t.value=s.isTheme});let a=y(new Map([["selected",[]],["select",[]]]));const r=y(["Level"]);w(r,(n,s)=>{m.getSquadMap(n[0])});const g=n=>{m.saveSquadMap(r.value[0],n)},m=K();return m.getSquadMap(r.value[0]),m.$subscribe((n,s)=>{var l,f;a.value.set("selected",(l=s.squadMap)==null?void 0:l.mapSelected),a.value.set("select",(f=s.squadMap)==null?void 0:f.mapSelect)}),(n,s)=>{const l=x("a-menu-item"),f=x("a-menu");return S(),v(N,null,[d(f,{selectedKeys:r.value,"onUpdate:selectedKeys":s[0]||(s[0]=M=>r.value=M),theme:b(t),mode:"horizontal"},{default:_(()=>[d(l,{key:"Level"},{default:_(()=>[h(" LevelList ")]),_:1}),d(l,{key:"Layer"},{default:_(()=>[h(" LayerList ")]),_:1})]),_:1},8,["selectedKeys","theme"]),i("div",X,[d(W,{onSaveMap:g,mapArr:b(a)},null,8,["mapArr"])])],64)}}});const re=k(Y,[["__scopeId","data-v-c17d1c6e"]]);export{re as default};

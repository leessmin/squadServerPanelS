import{n as y,k as u,D as A}from"./index-2df293b2.js";import{h as t}from"./http-9bd06147.js";import{W as m,M as c}from"./worker-e1b6861a.js";import{g as E,d as G}from"./getSteamName-10fa85f4.js";const S=y("admin",()=>{const r=u();async function d(){const e=await t().Require("/BA/adminGroup/get",{});if((e==null?void 0:e.code)!=200)return;if(e.data.adminGroup==null){r.value=[];return}const a=new m;new c(a,s=>{r.value=s.data}).sendMsg(e.data.adminGroup)}async function p(e){const a=await t().Require("BA/adminGroup/addEdit",{method:"POST",body:JSON.stringify(e)});if((a==null?void 0:a.code)!=200){d();return}}async function f(e){const a=await t().Require(`BA/adminGroup/del?groupName=${e}`,{method:"DELETE"});if((a==null?void 0:a.code)!=200){d();return}}const i=u([]),l=A(()=>{var a;let e=[];return(a=r.value)==null||a.forEach(n=>{e.push({value:n.groupName,label:n.groupName})}),e});async function o(){d();const e=await t().Require("/BA/adminUser/get",{});if((e==null?void 0:e.code)!=200)return;if(e.data.adminUser==null){i.value=[];return}const a=new m;new c(a,s=>{i.value=s.data,E(i.value)}).sendMsg(e.data.adminUser)}async function g(e){await t().Require("/BA/adminUser/addEdit",{method:"POST",body:JSON.stringify(e)}),o()}async function w(e){let a=G("steamIds",e);const n=await t().Require(`/BA/adminUser/del?${a}`,{method:"DELETE"});(n==null?void 0:n.code)!=200&&o()}return{adminGroup:r,getAdminGroup:d,addEditAdminGroup:p,delAdminGroup:f,adminUser:i,groupNameList:l,getAdminUser:o,addEditAdminUser:g,delAdminUser:w}});export{S as u};

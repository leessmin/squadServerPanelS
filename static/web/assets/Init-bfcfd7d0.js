import{u as b}from"./useLoginStore-9aa4a7ba.js";import{d as h,x as g,r as w,c as n,o as x,e as S,f as l,g as e,h as r,i as I,s as y,v as P}from"./index-2df293b2.js";import{_ as q}from"./_plugin-vue_export-helper-c27b6911.js";import"./http-9bd06147.js";import"./config-9c65774d.js";import"./formData-4b5941f0.js";const F=i=>(y("data-v-823b4e4f"),i=i(),P(),i),B={class:"setting"},C=F(()=>l("div",{class:"title-box"},[l("h1",null,"初始化服务器"),l("p",null,"第一次加载服务器需要进行初始化设置")],-1)),N={class:"confirm-box"},U=h({__name:"Init",setup(i){g.info("检测到第一次登录面板，需要进行初始化设置哦");const m=b(),o=w({username:"",password:"",serverPath:""}),_=t=>{console.log("Success:",t),m.initServer(t.username,t.password,t.serverPath)},d=t=>{console.log("Failed:",t)};return(t,s)=>{const u=n("a-input"),p=n("a-form-item"),c=n("a-input-password"),f=n("a-button"),v=n("a-form");return x(),S("main",null,[l("section",B,[C,e(v,{model:o,name:"basic","label-col":{style:{width:"160px"}},"wrapper-col":{span:16},autocomplete:"off",onFinish:_,onFinishFailed:d},{default:r(()=>[e(p,{label:"设置新账号",name:"username",rules:[{required:!0,message:"请输入新的账号！"}]},{default:r(()=>[e(u,{value:o.username,"onUpdate:value":s[0]||(s[0]=a=>o.username=a)},null,8,["value"])]),_:1}),e(p,{label:"设置新密码",name:"password",rules:[{required:!0,message:"请输入新的密码！"}]},{default:r(()=>[e(c,{value:o.password,"onUpdate:value":s[1]||(s[1]=a=>o.password=a)},null,8,["value"])]),_:1}),e(p,{label:"squad服务器安装路径",name:"serverPath",rules:[{required:!0,message:"请输入服务器安装路径"}]},{default:r(()=>[e(u,{value:o.serverPath,"onUpdate:value":s[2]||(s[2]=a=>o.serverPath=a)},null,8,["value"])]),_:1}),l("div",N,[e(f,{type:"primary",class:"confirm-btn","html-type":"submit"},{default:r(()=>[I("确认")]),_:1})])]),_:1},8,["model"])])])}}});const z=q(U,[["__scopeId","data-v-823b4e4f"]]);export{z as default};

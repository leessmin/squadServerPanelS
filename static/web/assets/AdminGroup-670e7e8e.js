import{S as F}from"./SectionPanel-bf0ce22b.js";import{d as H,k as U,r as B,w as D,o as b,e as N,g as e,h as t,q as w,x as L,c as g,y as M,z as P,f as c,i as u,F as G,A as q,t as E,B as J,j as z,C as $,s as K,v as Q}from"./index-2df293b2.js";import{u as T}from"./useAdminStore-4bb0703f.js";import{_ as R}from"./_plugin-vue_export-helper-c27b6911.js";import"./http-9bd06147.js";import"./config-9c65774d.js";import"./worker-e1b6861a.js";import"./getSteamName-10fa85f4.js";const W={style:{"font-size":"12px",color:"red"}},X={class:"form-gutter"},Y=c("h3",null,"权限：",-1),Z={class:"gutter-box"},ee={class:"gutter-box"},te={class:"gutter-box"},ae={class:"gutter-box"},oe={class:"gutter-box"},le={class:"gutter-box"},ne={class:"gutter-box"},se={class:"gutter-box"},re={class:"gutter-box"},ce={class:"gutter-box"},de={class:"gutter-box"},ie={class:"gutter-box"},ue={class:"gutter-box"},fe={class:"gutter-box"},pe={class:"gutter-box"},me={class:"gutter-box"},_e={class:"gutter-box"},ge=H({__name:"AdminGroupAddPanel",props:{visible:{type:Boolean},groupNameArr:null,keyValue:null,editObj:null},emits:["changeVisible","addData"],setup(x,{emit:C}){const i=x,h=U(i.visible);let k=U(!1),p=B({state:"",help:""});const a=B({groupName:"",info:"",changemap:!1,pause:!1,cheat:!1,private:!1,balance:!1,chat:!1,kick:!1,ban:!1,config:!1,cameraman:!1,immune:!1,manageserver:!1,reserve:!1,debug:!1,teamchange:!1,forceteamchange:!1,canseeadminchat:!1}),S=T(),O=async m=>{if(a.groupName==""){p.state="error",p.help="组名不能为空";return}if(i.groupNameArr.filter(d=>{var r;return d===a.groupName&&d!=((r=i.editObj)==null?void 0:r.groupName)}).length!==0){p.state="error",p.help="存在相同的组名";return}p.state="success",p.help="";let f={key:i.keyValue,groupName:a.groupName,info:a.info,auth:[]},n;for(n in a)n==="groupName"||n==="info"||a[n]&&f.auth.push(n);if(f.auth.length==0){L.warning("管理组的权限不能为空");return}console.log(f),k.value=!0,await S.addEditAdminGroup(f),h.value=!1,k.value=!1,C("addData",f)};return D(i,m=>{h.value=m.visible;let o;for(o in a)if(o==="groupName"||o==="info")a[o]="";else{if(!a[o])continue;a[o]=!1}m.editObj!=null&&(a.groupName=m.editObj.groupName,a.info=m.editObj.info,m.editObj.auth.forEach(f=>{f=="groupName"||f=="info"||(a[f]=!0)}))}),D(h,m=>{C("changeVisible",m)}),(m,o)=>{const f=g("a-input"),n=g("a-form-item"),d=g("a-checkbox"),r=g("a-col"),s=g("a-row"),_=g("a-form"),v=g("a-modal");return b(),N("div",null,[e(v,{visible:h.value,"onUpdate:visible":o[19]||(o[19]=l=>h.value=l),maskClosable:!1,title:"添加管理组",onOk:O,okText:"确认","confirm-loading":w(k),cancelText:"取消",width:"1000px"},{default:t(()=>[e(_,{model:a,name:"basic","label-col":{span:8},"wrapper-col":{span:10},autocomplete:"off"},{default:t(()=>[e(n,{label:"组名",name:"groupName",validateStatus:w(p).state,help:w(p).help},{default:t(()=>{var l,I;return[e(f,{value:a.groupName,"onUpdate:value":o[0]||(o[0]=V=>a.groupName=V),disabled:((l=x.editObj)==null?void 0:l.groupName.length)>0},null,8,["value","disabled"]),M(c("span",W,"编辑操作不能修改组名，如需修改组名，建议删除后重新添加",512),[[P,((I=x.editObj)==null?void 0:I.groupName.length)>0]])]}),_:1},8,["validateStatus","help"]),e(n,{label:"备注",name:"info"},{default:t(()=>[e(f,{value:a.info,"onUpdate:value":o[1]||(o[1]=l=>a.info=l)},null,8,["value"])]),_:1}),c("div",X,[Y,e(s,{gutter:16},{default:t(()=>[e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",Z,[e(n,{name:"changemap","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.changemap,"onUpdate:checked":o[2]||(o[2]=l=>a.changemap=l)},{default:t(()=>[u("更换/预设地图")]),_:1},8,["checked"])]),_:1})])]),_:1}),e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",ee,[e(n,{name:"pause","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.pause,"onUpdate:checked":o[3]||(o[3]=l=>a.pause=l)},{default:t(()=>[u("暂停游戏")]),_:1},8,["checked"])]),_:1})])]),_:1}),e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",te,[e(n,{name:"cheat","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.cheat,"onUpdate:checked":o[4]||(o[4]=l=>a.cheat=l)},{default:t(()=>[u("使用作弊命令")]),_:1},8,["checked"])]),_:1})])]),_:1}),e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",ae,[e(n,{name:"private","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.private,"onUpdate:checked":o[5]||(o[5]=l=>a.private=l)},{default:t(()=>[u("设置服务器密码")]),_:1},8,["checked"])]),_:1})])]),_:1})]),_:1}),e(s,{gutter:16},{default:t(()=>[e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",oe,[e(n,{name:"balance","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.balance,"onUpdate:checked":o[6]||(o[6]=l=>a.balance=l)},{default:t(()=>[u("忽略服务器阵营平衡")]),_:1},8,["checked"])]),_:1})])]),_:1}),e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",le,[e(n,{name:"chat","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.chat,"onUpdate:checked":o[7]||(o[7]=l=>a.chat=l)},{default:t(()=>[u("管理员聊天/服务器公告")]),_:1},8,["checked"])]),_:1})])]),_:1}),e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",ne,[e(n,{name:"kick","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.kick,"onUpdate:checked":o[8]||(o[8]=l=>a.kick=l)},{default:t(()=>[u("踢出玩家")]),_:1},8,["checked"])]),_:1})])]),_:1}),e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",se,[e(n,{name:"ban","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.ban,"onUpdate:checked":o[9]||(o[9]=l=>a.ban=l)},{default:t(()=>[u("封禁玩家")]),_:1},8,["checked"])]),_:1})])]),_:1})]),_:1}),e(s,{gutter:16},{default:t(()=>[e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",re,[e(n,{name:"config","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.config,"onUpdate:checked":o[10]||(o[10]=l=>a.config=l)},{default:t(()=>[u("更改服务器配置")]),_:1},8,["checked"])]),_:1})])]),_:1}),e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",ce,[e(n,{name:"cameraman","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.cameraman,"onUpdate:checked":o[11]||(o[11]=l=>a.cameraman=l)},{default:t(()=>[u("摄影机-管理员视角")]),_:1},8,["checked"])]),_:1})])]),_:1}),e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",de,[e(n,{name:"immune","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.immune,"onUpdate:checked":o[12]||(o[12]=l=>a.immune=l)},{default:t(()=>[u("无法被 踢出/封禁")]),_:1},8,["checked"])]),_:1})])]),_:1}),e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",ie,[e(n,{name:"manageserver","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.manageserver,"onUpdate:checked":o[13]||(o[13]=l=>a.manageserver=l)},{default:t(()=>[u("关闭服务器")]),_:1},8,["checked"])]),_:1})])]),_:1})]),_:1}),e(s,{gutter:16},{default:t(()=>[e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",ue,[e(n,{name:"reserve","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.reserve,"onUpdate:checked":o[14]||(o[14]=l=>a.reserve=l)},{default:t(()=>[u("预留位")]),_:1},8,["checked"])]),_:1})])]),_:1}),e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",fe,[e(n,{name:"debug","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.debug,"onUpdate:checked":o[15]||(o[15]=l=>a.debug=l)},{default:t(()=>[u("调试")]),_:1},8,["checked"])]),_:1})])]),_:1}),e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",pe,[e(n,{name:"teamchange","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.teamchange,"onUpdate:checked":o[16]||(o[16]=l=>a.teamchange=l)},{default:t(()=>[u("忽略更换阵营时间限制")]),_:1},8,["checked"])]),_:1})])]),_:1}),e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",me,[e(n,{name:"forceteamchange","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.forceteamchange,"onUpdate:checked":o[17]||(o[17]=l=>a.forceteamchange=l)},{default:t(()=>[u("允许执行强制更换阵营命令")]),_:1},8,["checked"])]),_:1})])]),_:1})]),_:1}),e(s,{gutter:16},{default:t(()=>[e(r,{class:"gutter-row",span:6},{default:t(()=>[c("div",_e,[e(n,{name:"canseeadminchat","wrapper-col":{offset:8,span:16}},{default:t(()=>[e(d,{checked:a.canseeadminchat,"onUpdate:checked":o[18]||(o[18]=l=>a.canseeadminchat=l)},{default:t(()=>[u("查看 管理员聊天/友军击杀/管理员加入消息")]),_:1},8,["checked"])]),_:1})])]),_:1})]),_:1})])]),_:1},8,["model"])]),_:1},8,["visible","confirm-loading"])])}}}),he=x=>(K("data-v-45b2cd6f"),x=x(),Q(),x),ke=he(()=>c("div",{class:"title-box"},[c("h3",null,"管理组")],-1)),ve={class:"btn-box"},be={key:1},we={key:0,style:{color:"#fff"}},xe=H({__name:"AdminGroup",setup(x){const C=[{title:"组名",dataIndex:"groupName",width:"150px",fixed:"left"},{title:"备注",dataIndex:"info"},{title:"权限",dataIndex:"auth"},{title:"操作",dataIndex:"operation",width:200,fixed:"right"}],i=U([]),h=new Map([["changemap",{title:"更换/预设地图",color:"#eccc68"}],["pause",{title:"暂停游戏",color:"#ff6b81"}],["cheat",{title:"使用作弊命令",color:"#ff7f50"}],["private",{title:"设置服务器密码",color:"#ffa502"}],["balance",{title:"忽略服务器阵营平衡",color:"#69c987"}],["chat",{title:"管理员聊天/服务器公告",color:"#c9b984"}],["kick",{title:"踢出玩家",color:"#70a1ff"}],["ban",{title:"封禁玩家",color:"#00a8ff"}],["config",{title:"更改服务器配置",color:"#9aecdb"}],["cameraman",{title:"摄影机-管理员视角",color:"#1e90ff"}],["immune",{title:"无法被 踢出/封禁",color:"#e056fd"}],["manageserver",{title:"关闭服务器",color:"#55e6c1"}],["reserve",{title:"预留位",color:"#ff9ff3"}],["debug",{title:"调试",color:"#f368e0"}],["teamchange",{title:"忽略更换阵营时间限制",color:"#feca57"}],["forceteamchange",{title:"允许执行强制更换阵营命令",color:"#ff9f43"}],["canseeadminchat",{title:"查看 管理员聊天/友军击杀/管理员加入消息",color:"#48dbfb"}]]),k=U(!1),p=U([]);d(i.value);let a=U(),S=U("");function O(){k.value=!0}function m(s){console.log(s),a.value=s,S.value=s.key,k.value=!0}function o(s){r.delAdminGroup(s.groupName),i.value=i.value.filter(_=>_.key!==s.key)}function f(s){k.value=s,s||(a.value=void 0,S.value=String(parseInt(i.value[i.value.length-1].key)+1))}function n(s){let _=!1;i.value.forEach((v,l)=>{v.key==s.key&&(i.value[l]=s,_=!0)}),_||i.value.push(s),k.value=!1}D(i.value,s=>{d(s)});function d(s){p.value=[],s.forEach(_=>{p.value.push(_.groupName)})}const r=T();return r.getAdminGroup(),r.$subscribe((s,_)=>{i.value=_.adminGroup;let v=i.value[i.value.length-1]!=null?i.value[i.value.length-1].key+1:"0";S.value=String(parseInt(v))}),(s,_)=>{const v=g("a-button"),l=g("a-tag"),I=g("a-popconfirm"),V=g("a-table");return b(),N(G,null,[c("main",null,[e(F,{class:"mt-zero"},{default:t(()=>[ke,c("div",ve,[e(v,{type:"primary",class:"editable-add-btn",onClick:O},{default:t(()=>[u("添加")]),_:1})]),e(V,{bordered:"","data-source":i.value,columns:C,pagination:!1,scroll:{x:1100}},{bodyCell:t(({column:j,text:ye,record:A})=>[j.dataIndex==="info"?(b(),N("span",{key:0,style:q({color:A.info==""?"var(--heading-color)":"var(--text-color)"})},E(A.info==""?"无备注":A.info),5)):j.dataIndex==="auth"?(b(),N("span",be,[(b(!0),N(G,null,J(A.auth,y=>(b(),z(l,{key:y,color:w(h).has(y)?w(h).get(y).color:"#fff",style:{"margin-bottom":"8px","font-size":"medium"}},{default:t(()=>[w(h).has(y)?(b(),N("span",we,E(w(h).get(y).title),1)):$("",!0)]),_:2},1032,["color"]))),128))])):j.dataIndex==="operation"?(b(),N(G,{key:2},[e(v,{type:"link",onClick:y=>m(A)},{default:t(()=>[u("编辑")]),_:2},1032,["onClick"]),i.value.length?(b(),z(I,{key:0,title:"删除会将组下的全部管理员一起删除，你确定要删除吗?",okText:"确定",cancelText:"取消",onConfirm:y=>o(A)},{default:t(()=>[e(v,{danger:"",type:"link"},{default:t(()=>[u("删除")]),_:1})]),_:2},1032,["onConfirm"])):$("",!0)],64)):$("",!0)]),_:1},8,["data-source"])]),_:1})]),e(ge,{visible:k.value,onAddData:n,onChangeVisible:f,groupNameArr:p.value,keyValue:w(S),editObj:w(a)},null,8,["visible","groupNameArr","keyValue","editObj"])],64)}}});const je=R(xe,[["__scopeId","data-v-45b2cd6f"]]);export{je as default};
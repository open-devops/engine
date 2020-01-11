(window.webpackJsonp=window.webpackJsonp||[]).push([[34],{158:function(e,r,t){"use strict";t.r(r),t.d(r,"frontMatter",(function(){return a})),t.d(r,"metadata",(function(){return i})),t.d(r,"rightToc",(function(){return p})),t.d(r,"default",(function(){return s}));var n=t(1),o=t(9),c=(t(0),t(177)),a={title:"Reference [string]"},i={id:"opspec/reference/structure/op-directory/op/reference",title:"Reference [string]",description:"A string used to reference the source of or destination for a value.",source:"@site/docs/opspec/reference/structure/op-directory/op/reference.md",permalink:"/docs/opspec/reference/structure/op-directory/op/reference",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/opspec/reference/structure/op-directory/op/reference.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1578700596,sidebar:"docs",previous:{title:"Markdown [string]",permalink:"/docs/opspec/reference/structure/op-directory/op/markdown"},next:{title:"Variable Name [string]",permalink:"/docs/opspec/reference/structure/op-directory/op/variable-name"}},p=[],u={rightToc:p},l="wrapper";function s(e){var r=e.components,t=Object(o.a)(e,["components"]);return Object(c.b)(l,Object(n.a)({},u,t,{components:r,mdxType:"MDXLayout"}),Object(c.b)("p",null,"A string used to reference the source of or destination for a value."),Object(c.b)("p",null,"The format is ",Object(c.b)("inlineCode",{parentName:"p"},"$(LOCATION)")," where ",Object(c.b)("inlineCode",{parentName:"p"},"LOCATION")," MUST be one of:"),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},"an absolute path to a file or directory included (a.k.a embedded) in the current operation e.g. the current op.yml can be referenced via ",Object(c.b)("inlineCode",{parentName:"li"},"$(/op.yml"),"."),Object(c.b)("li",{parentName:"ul"},"a ",Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"variable-name"}),"variable name")," representing a scoped variable.")))}s.isMDXComponent=!0},177:function(e,r,t){"use strict";t.d(r,"a",(function(){return s})),t.d(r,"b",(function(){return m}));var n=t(0),o=t.n(n);function c(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function a(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);r&&(n=n.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,n)}return t}function i(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?a(Object(t),!0).forEach((function(r){c(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):a(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function p(e,r){if(null==e)return{};var t,n,o=function(e,r){if(null==e)return{};var t,n,o={},c=Object.keys(e);for(n=0;n<c.length;n++)t=c[n],r.indexOf(t)>=0||(o[t]=e[t]);return o}(e,r);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);for(n=0;n<c.length;n++)t=c[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}var u=o.a.createContext({}),l=function(e){var r=o.a.useContext(u),t=r;return e&&(t="function"==typeof e?e(r):i({},r,{},e)),t},s=function(e){var r=l(e.components);return o.a.createElement(u.Provider,{value:r},e.children)},f="mdxType",d={inlineCode:"code",wrapper:function(e){var r=e.children;return o.a.createElement(o.a.Fragment,{},r)}},b=Object(n.forwardRef)((function(e,r){var t=e.components,n=e.mdxType,c=e.originalType,a=e.parentName,u=p(e,["components","mdxType","originalType","parentName"]),s=l(t),f=n,b=s["".concat(a,".").concat(f)]||s[f]||d[f]||c;return t?o.a.createElement(b,i({ref:r},u,{components:t})):o.a.createElement(b,i({ref:r},u))}));function m(e,r){var t=arguments,n=r&&r.mdxType;if("string"==typeof e||n){var c=t.length,a=new Array(c);a[0]=b;var i={};for(var p in r)hasOwnProperty.call(r,p)&&(i[p]=r[p]);i.originalType=e,i[f]="string"==typeof e?e:n,a[1]=i;for(var u=2;u<c;u++)a[u]=t[u];return o.a.createElement.apply(null,a)}return o.a.createElement.apply(null,t)}b.displayName="MDXCreateElement"}}]);
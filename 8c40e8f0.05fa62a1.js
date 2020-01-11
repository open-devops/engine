(window.webpackJsonp=window.webpackJsonp||[]).push([[30],{154:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return a})),r.d(t,"metadata",(function(){return i})),r.d(t,"rightToc",(function(){return p})),r.d(t,"default",(function(){return s}));var n=r(1),o=r(9),c=(r(0),r(177)),a={title:"Docker",sidebar_label:"Docker"},i={id:"setup/docker",title:"Docker",description:"An official docker opctl image is [maintained on docker hub](https://hub.docker.com/r/opctl/opctl/) (as of v0.1.15), which features a ready to use opctl node.",source:"@site/docs/setup/docker.md",permalink:"/docs/setup/docker",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/setup/docker.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1578700596,sidebar_label:"Docker",sidebar:"docs",previous:{title:"Azure Pipelines",permalink:"/docs/setup/azure-pipelines"},next:{title:"Gitlab",permalink:"/docs/setup/gitlab"}},p=[{value:"Examples",id:"examples",children:[]}],l={rightToc:p},u="wrapper";function s(e){var t=e.components,r=Object(o.a)(e,["components"]);return Object(c.b)(u,Object(n.a)({},l,r,{components:t,mdxType:"MDXLayout"}),Object(c.b)("p",null,"An official docker opctl image is ",Object(c.b)("a",Object(n.a)({parentName:"p"},{href:"https://hub.docker.com/r/opctl/opctl/"}),"maintained on docker hub")," (as of v0.1.15), which features a ready to use opctl node."),Object(c.b)("blockquote",null,Object(c.b)("p",{parentName:"blockquote"},"The container runtime in this case will be an embedded docker daemon, which leads to the ",Object(c.b)("inlineCode",{parentName:"p"},"--privileged")," flag being required")),Object(c.b)("h3",{id:"examples"},"Examples"),Object(c.b)("pre",null,Object(c.b)("code",Object(n.a)({parentName:"pre"},{className:"language-shell"}),"docker run --privileged opctl/opctl:0.1.25 opctl run github.com/opspec-pkgs/uuid.v4.generate#1.0.0\n")))}s.isMDXComponent=!0},177:function(e,t,r){"use strict";r.d(t,"a",(function(){return s})),r.d(t,"b",(function(){return m}));var n=r(0),o=r.n(n);function c(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function a(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?a(Object(r),!0).forEach((function(t){c(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):a(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function p(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},c=Object.keys(e);for(n=0;n<c.length;n++)r=c[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);for(n=0;n<c.length;n++)r=c[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var l=o.a.createContext({}),u=function(e){var t=o.a.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):i({},t,{},e)),r},s=function(e){var t=u(e.components);return o.a.createElement(l.Provider,{value:t},e.children)},d="mdxType",b={inlineCode:"code",wrapper:function(e){var t=e.children;return o.a.createElement(o.a.Fragment,{},t)}},f=Object(n.forwardRef)((function(e,t){var r=e.components,n=e.mdxType,c=e.originalType,a=e.parentName,l=p(e,["components","mdxType","originalType","parentName"]),s=u(r),d=n,f=s["".concat(a,".").concat(d)]||s[d]||b[d]||c;return r?o.a.createElement(f,i({ref:t},l,{components:r})):o.a.createElement(f,i({ref:t},l))}));function m(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var c=r.length,a=new Array(c);a[0]=f;var i={};for(var p in t)hasOwnProperty.call(t,p)&&(i[p]=t[p]);i.originalType=e,i[d]="string"==typeof e?e:n,a[1]=i;for(var l=2;l<c;l++)a[l]=r[l];return o.a.createElement.apply(null,a)}return o.a.createElement.apply(null,r)}f.displayName="MDXCreateElement"}}]);
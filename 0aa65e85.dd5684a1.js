(window.webpackJsonp=window.webpackJsonp||[]).push([[6],{144:function(e,t,r){"use strict";r.d(t,"a",(function(){return s})),r.d(t,"b",(function(){return d}));var a=r(0),n=r.n(a);function i(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function c(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,a)}return r}function l(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?c(Object(r),!0).forEach((function(t){i(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):c(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function o(e,t){if(null==e)return{};var r,a,n=function(e,t){if(null==e)return{};var r,a,n={},i=Object.keys(e);for(a=0;a<i.length;a++)r=i[a],t.indexOf(r)>=0||(n[r]=e[r]);return n}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(a=0;a<i.length;a++)r=i[a],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(n[r]=e[r])}return n}var b=n.a.createContext({}),p=function(e){var t=n.a.useContext(b),r=t;return e&&(r="function"==typeof e?e(t):l({},t,{},e)),r},s=function(e){var t=p(e.components);return n.a.createElement(b.Provider,{value:t},e.children)},m="mdxType",u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.a.createElement(n.a.Fragment,{},t)}},O=Object(a.forwardRef)((function(e,t){var r=e.components,a=e.mdxType,i=e.originalType,c=e.parentName,b=o(e,["components","mdxType","originalType","parentName"]),s=p(r),m=a,O=s["".concat(c,".").concat(m)]||s[m]||u[m]||i;return r?n.a.createElement(O,l({ref:t},b,{components:r})):n.a.createElement(O,l({ref:t},b))}));function d(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=r.length,c=new Array(i);c[0]=O;var l={};for(var o in t)hasOwnProperty.call(t,o)&&(l[o]=t[o]);l.originalType=e,l[m]="string"==typeof e?e:a,c[1]=l;for(var b=2;b<i;b++)c[b]=r[b];return n.a.createElement.apply(null,c)}return n.a.createElement.apply(null,r)}O.displayName="MDXCreateElement"},92:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return c})),r.d(t,"metadata",(function(){return l})),r.d(t,"rightToc",(function(){return o})),r.d(t,"default",(function(){return s}));var a=r(1),n=r(6),i=(r(0),r(144)),c={title:"Array"},l={id:"opspec/reference/types/array",title:"Array",description:"Array typed values are a container for numerically indexed values (referred to as items).",source:"@site/docs/opspec/reference/types/array.md",permalink:"/docs/opspec/reference/types/array",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/opspec/reference/types/array.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1579135864,sidebar:"docs",previous:{title:"Variable Name [string]",permalink:"/docs/opspec/reference/structure/op-directory/op/variable-name"},next:{title:"Boolean",permalink:"/docs/opspec/reference/types/boolean"}},o=[{value:"Initialization",id:"initialization",children:[]},{value:"Item Referencing",id:"item-referencing",children:[]},{value:"Coercion",id:"coercion",children:[]}],b={rightToc:o},p="wrapper";function s(e){var t=e.components,r=Object(n.a)(e,["components"]);return Object(i.b)(p,Object(a.a)({},b,r,{components:t,mdxType:"MDXLayout"}),Object(i.b)("p",null,"Array typed values are a container for numerically indexed values (referred to as items)."),Object(i.b)("p",null,"Arrays..."),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},"are immutable, i.e. assigning to an array results in a copy of the original array"),Object(i.b)("li",{parentName:"ul"},"can be passed in/out of ops via ",Object(i.b)("a",Object(a.a)({parentName:"li"},{href:"../structure/op-directory/op/parameter/array"}),"array parameters")),Object(i.b)("li",{parentName:"ul"},"can be initialized via ",Object(i.b)("a",Object(a.a)({parentName:"li"},{href:"#initialization"}),"array initialization")),Object(i.b)("li",{parentName:"ul"},"items can be referenced via ",Object(i.b)("a",Object(a.a)({parentName:"li"},{href:"#item-referencing"}),"array item referencing")),Object(i.b)("li",{parentName:"ul"},"are coerced according to ",Object(i.b)("a",Object(a.a)({parentName:"li"},{href:"#coercion"}),"array coercion"))),Object(i.b)("h3",{id:"initialization"},"Initialization"),Object(i.b)("p",null,"Array typed values can be constructed from a literal or templated array."),Object(i.b)("p",null,"A templated array is an array which includes one or more value reference.\nAt runtime, each reference gets evaluated and replaced with it's corresponding value. "),Object(i.b)("h4",{id:"initialization-example-literal"},"Initialization Example (literal)"),Object(i.b)("pre",null,Object(i.b)("code",Object(a.a)({parentName:"pre"},{className:"language-yaml"}),"- item1\n- item2\n")),Object(i.b)("h4",{id:"initialization-example-templated"},"Initialization Example (templated)"),Object(i.b)("p",null,"given:"),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},Object(i.b)("inlineCode",{parentName:"li"},"/someDir/file2.txt")," is embedded in op"),Object(i.b)("li",{parentName:"ul"},Object(i.b)("inlineCode",{parentName:"li"},"someObject")," ",Object(i.b)("ul",{parentName:"li"},Object(i.b)("li",{parentName:"ul"},"is in scope"),Object(i.b)("li",{parentName:"ul"},"is type coercible to object"),Object(i.b)("li",{parentName:"ul"},"has property ",Object(i.b)("inlineCode",{parentName:"li"},"someProperty"))))),Object(i.b)("pre",null,Object(i.b)("code",Object(a.a)({parentName:"pre"},{className:"language-yaml"}),"- string $(/someDir/file2.txt)\n- $(someObject.someProperty)\n- [ sub, array, 2]\n")),Object(i.b)("h3",{id:"item-referencing"},"Item Referencing"),Object(i.b)("p",null,"Array items can be referenced via ",Object(i.b)("inlineCode",{parentName:"p"},"$(ARRAY[index])")," syntax, where ",Object(i.b)("inlineCode",{parentName:"p"},"index")," is the zero based index of the item.\nIf ",Object(i.b)("inlineCode",{parentName:"p"},"index")," is negative, indexing will take place from the end of the array."),Object(i.b)("h4",{id:"item-referencing-example-first-item"},"Item Referencing Example (first item)"),Object(i.b)("p",null,"given:"),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},"someArray",Object(i.b)("ul",{parentName:"li"},Object(i.b)("li",{parentName:"ul"},"is in scope"),Object(i.b)("li",{parentName:"ul"},"is type coercible to array"),Object(i.b)("li",{parentName:"ul"},"has at least one item")))),Object(i.b)("pre",null,Object(i.b)("code",Object(a.a)({parentName:"pre"},{className:"language-yaml"}),"$(someArray[0])\n")),Object(i.b)("h4",{id:"item-referencing-example-last-item"},"Item Referencing Example (last item)"),Object(i.b)("p",null,"given:"),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},"someArray",Object(i.b)("ul",{parentName:"li"},Object(i.b)("li",{parentName:"ul"},"is in scope"),Object(i.b)("li",{parentName:"ul"},"is type coercible to array"),Object(i.b)("li",{parentName:"ul"},"has at least one item")))),Object(i.b)("pre",null,Object(i.b)("code",Object(a.a)({parentName:"pre"},{className:"language-yaml"}),"$(someArray[-1])\n")),Object(i.b)("h3",{id:"coercion"},"Coercion"),Object(i.b)("p",null,"Array typed values are coercible to:"),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},Object(i.b)("a",Object(a.a)({parentName:"li"},{href:"/docs/opspec/reference/types/boolean"}),"boolean")," (arrays which are null or empty coerce to false; all else coerce to true)"),Object(i.b)("li",{parentName:"ul"},Object(i.b)("a",Object(a.a)({parentName:"li"},{href:"/docs/opspec/reference/types/file"}),"file")," (will be serialized to JSON)"),Object(i.b)("li",{parentName:"ul"},Object(i.b)("a",Object(a.a)({parentName:"li"},{href:"/docs/opspec/reference/types/string"}),"string")," (will be serialized to JSON)")))}s.isMDXComponent=!0}}]);
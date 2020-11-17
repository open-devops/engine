(window.webpackJsonp=window.webpackJsonp||[]).push([[28],{130:function(e,t,n){"use strict";n.r(t),n.d(t,"frontMatter",(function(){return i})),n.d(t,"metadata",(function(){return s})),n.d(t,"rightToc",(function(){return c})),n.d(t,"default",(function(){return u}));var r=n(1),a=n(6),o=(n(0),n(176)),i={title:"Run a go service",sidebar_label:"run a go service"},s={id:"run-a-go-service",title:"Run a go service",description:"We'll now look at an op to run a sample Go application. Please see the code at [https://github.com/opctl/opctl/tree/master/examples/run-a-go-service](https://github.com/opctl/opctl/tree/master/examples/run-a-go-service)",source:"@site/docs/run-a-go-service.md",permalink:"/docs/run-a-go-service",editUrl:"https://github.com/opctl/opctl/edit/master/website/docs/run-a-go-service.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1605581772,sidebar_label:"run a go service",sidebar:"docs",previous:{title:"Inputs & Outputs",permalink:"/docs/training/inputs-outputs"},next:{title:"Run a react app",permalink:"/docs/run-a-react-app"}},c=[],l={rightToc:c},p="wrapper";function u(e){var t=e.components,n=Object(a.a)(e,["components"]);return Object(o.b)(p,Object(r.a)({},l,n,{components:t,mdxType:"MDXLayout"}),Object(o.b)("p",null,"We'll now look at an op to run a sample Go application. Please see the code at ",Object(o.b)("a",Object(r.a)({parentName:"p"},{href:"https://github.com/opctl/opctl/tree/master/examples/run-a-go-service"}),"https://github.com/opctl/opctl/tree/master/examples/run-a-go-service"),"\nThe sample application we have requires a mysql database, therefore to run it locally we need to:\n1. Start a MySQL database with some DB credentials\n2. Seed the DB with sample data\n3. Start the application and provide the MySQL DB credentials as inputs - our application will read those from environment variables"),Object(o.b)("p",null,"The ops to make that happen are explained below."),Object(o.b)("h4",{id:"mysql"},"mysql"),Object(o.b)("pre",null,Object(o.b)("code",Object(r.a)({parentName:"pre"},{className:"language-yaml"}),"name: mysql\ndescription: runs a mysql container, seeding it with sample data\ninputs:\n  MYSQL_USER:\n    string:\n      default: testuser\n      description: username for MySQL user to create\n  MYSQL_PASSWORD:\n    string:\n      default: testpassword\n      description: password for MySQL user to create\n  MYSQL_DATABASE:\n    string:\n      default: testapp\n      description: name of mysql database to create\n  MYSQL_HOST:\n    string:\n      default: run-a-go-service-mysql\nrun:\n  container:\n    dirs:\n      # mount our data seed scripts\n      /docker-entrypoint-initdb.d:\n    image:\n      ref: 'mysql:8'\n    # this sets an opctl overlay network DNS A record so the containers available via this name.\n    name: run-a-go-service-mysql\n    envVars:\n      MYSQL_USER:\n      MYSQL_PASSWORD:\n      MYSQL_DATABASE:\n      MYSQL_RANDOM_ROOT_PASSWORD: \"yes\"\n    ports:\n      3306: 3306\n")),Object(o.b)("p",null,"This op will run a MySQL database in a Docker container and seeds the MySQL database with a single table and a single row of data."),Object(o.b)("p",null,"Note that instead of writing the ",Object(o.b)("inlineCode",{parentName:"p"},"cmd")," inline, we could have also put the run script in a file in the op directory and referenced it.\ne.g."),Object(o.b)("pre",null,Object(o.b)("code",Object(r.a)({parentName:"pre"},{className:"language-yaml"}),'cmd: ["/run.sh"]\n')),Object(o.b)("p",null,"This becomes useful for readability if the script we're running in a container is large."),Object(o.b)("p",null,"Note that while we could have also built a custom docker image that contains that script and ran it, that's not a recommended practice. While the op remains portable with that approach, it is less transparent and harder to maintain, since we would have the extra step of building and pushing that image whenever we make a change to the script. Instead, we usually opt for the leanest Docker image that fulfills our use case, and include any scripts or binaries we need in the op directory."),Object(o.b)("h4",{id:"dev"},"dev"),Object(o.b)("pre",null,Object(o.b)("code",Object(r.a)({parentName:"pre"},{className:"language-yaml"}),"name: dev\ndescription: runs go-svc for development\ninputs:\n  MYSQL_USER:\n    string:\n      default: testuser\n      description: username for MySQL user to create\n  MYSQL_PASSWORD:\n    string:\n      default: testpassword\n      description: password for MySQL user to create\n  MYSQL_DATABASE:\n    string:\n      default: testapp\n      description: Database to create\n  MYSQL_HOST:\n    string:\n      default: mysql # the host for mysql is the container name in the mysql op\nrun:\n  parallel:\n    - op:\n        ref: ../mysql # we reference the mysql op using its relative path to this op\n        inputs:\n          # we pass the relevant inputs through to the mysql op\n          MYSQL_USER:\n          MYSQL_PASSWORD:\n          MYSQL_HOST:\n          MYSQL_DATABASE:\n    - container:\n        image:\n          ref: 'golang:1.15'\n        name: go-svc\n        dirs:\n          # mount the source code of our app to the container\n          /src: $(../..) \n        envVars:\n          # let our code know how to connect to the DB\n          MYSQL_USER:\n          MYSQL_PASSWORD:\n          MYSQL_HOST:\n          MYSQL_DATABASE:\n        workDir: /src\n        ports:\n          8080: 8080\n        cmd:\n          - sh\n          - -ce\n          - |\n            go get -u github.com/cespare/reflex # get reflex to watch and hot-reload our main.go\n            sleep 30 # we'll sleep while the MySQL DB starts\n            /go/bin/reflex -g 'main.go' -s -- sh -c 'go run main.go'\n")),Object(o.b)("p",null,"This op will do the following, in parallel:\n1. Run the mysql op, passing in the inputs needed to create it\n2. run our service in the ",Object(o.b)("inlineCode",{parentName:"p"},"golang")," container. We'll use ",Object(o.b)("a",Object(r.a)({parentName:"p"},{href:"https://github.com/cespare/reflex"}),"reflex")," to make development easier and avoid having to restart the whole operation after every change to main.go"),Object(o.b)("p",null,"Now to run the service locally, we only run ",Object(o.b)("inlineCode",{parentName:"p"},"opctl run dev")),Object(o.b)("p",null,"Notice how we can run ops in a ",Object(o.b)("inlineCode",{parentName:"p"},"parallel")," block. Our app runs parallel to the MySQL DB container (via the ",Object(o.b)("inlineCode",{parentName:"p"},"mysql")," op), because we need mysql running while our app runs."))}u.isMDXComponent=!0},176:function(e,t,n){"use strict";n.d(t,"a",(function(){return u})),n.d(t,"b",(function(){return h}));var r=n(0),a=n.n(r);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function s(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function c(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var l=a.a.createContext({}),p=function(e){var t=a.a.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):s({},t,{},e)),n},u=function(e){var t=p(e.components);return(a.a.createElement(l.Provider,{value:t},e.children))},d="mdxType",m={inlineCode:"code",wrapper:function(e){var t=e.children;return a.a.createElement(a.a.Fragment,{},t)}},b=Object(r.forwardRef)((function(e,t){var n=e.components,r=e.mdxType,o=e.originalType,i=e.parentName,l=c(e,["components","mdxType","originalType","parentName"]),u=p(n),d=r,b=u["".concat(i,".").concat(d)]||u[d]||m[d]||o;return n?a.a.createElement(b,s({ref:t},l,{components:n})):a.a.createElement(b,s({ref:t},l))}));function h(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var o=n.length,i=new Array(o);i[0]=b;var s={};for(var c in t)hasOwnProperty.call(t,c)&&(s[c]=t[c]);s.originalType=e,s[d]="string"==typeof e?e:r,i[1]=s;for(var l=2;l<o;l++)i[l]=n[l];return a.a.createElement.apply(null,i)}return a.a.createElement.apply(null,n)}b.displayName="MDXCreateElement"}}]);
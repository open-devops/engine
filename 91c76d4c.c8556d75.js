(window.webpackJsonp=window.webpackJsonp||[]).push([[30],{170:function(e,t,n){"use strict";n.r(t),n.d(t,"frontMatter",(function(){return c})),n.d(t,"metadata",(function(){return i})),n.d(t,"rightToc",(function(){return r})),n.d(t,"default",(function(){return s}));var a=n(1),l=n(9),o=(n(0),n(197)),c={title:"Commands",sidebar_label:"CLI"},i={id:"reference/cli",title:"Commands",description:"## `opctl`",source:"@site/docs/reference/cli.md",permalink:"/docs/reference/cli",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/reference/cli.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1588472011,sidebar_label:"CLI",sidebar:"docs",previous:{title:"UI",permalink:"/docs/reference/ui"}},r=[{value:"<code>opctl</code>",id:"opctl",children:[{value:"Options",id:"options",children:[]}]},{value:"<code>opctl events</code>",id:"opctl-events",children:[{value:"Examples",id:"examples-2",children:[]}]},{value:"<code>opctl ls [DIR_REF]</code>",id:"opctl-ls-dir_ref",children:[{value:"Arguments",id:"arguments",children:[]},{value:"Examples",id:"examples-3",children:[]}]},{value:"<code>opctl run [OPTIONS] OP_REF</code>",id:"opctl-run-options-op_ref",children:[{value:"Arguments",id:"arguments-1",children:[]},{value:"Options",id:"options-1",children:[]},{value:"Examples",id:"examples-4",children:[]},{value:"Notes",id:"notes",children:[]}]},{value:"<code>opctl node create</code>",id:"opctl-node-create",children:[{value:"Options",id:"options-2",children:[]},{value:"Notes",id:"notes-1",children:[]}]},{value:"<code>opctl node kill</code>",id:"opctl-node-kill",children:[]},{value:"<code>opctl op create [OPTIONS] NAME</code>",id:"opctl-op-create-options-name",children:[{value:"Arguments",id:"arguments-2",children:[]},{value:"Options",id:"options-3",children:[]},{value:"Examples",id:"examples-5",children:[]}]},{value:"<code>op install [OPTIONS] OP_REF</code>",id:"op-install-options-op_ref",children:[{value:"Arguments",id:"arguments-3",children:[]},{value:"Options",id:"options-4",children:[]},{value:"Examples",id:"examples-6",children:[]},{value:"Notes",id:"notes-2",children:[]}]},{value:"<code>op kill OP_ID</code>",id:"op-kill-op_id",children:[{value:"Arguments",id:"arguments-4",children:[]}]},{value:"<code>op validate OP_REF</code>",id:"op-validate-op_ref",children:[{value:"Arguments",id:"arguments-5",children:[]},{value:"Examples",id:"examples-7",children:[]},{value:"Notes",id:"notes-3",children:[]}]},{value:"<code>opctl self-update [OPTIONS]</code>",id:"opctl-self-update-options",children:[{value:"Options",id:"options-5",children:[]},{value:"Examples",id:"examples-8",children:[]}]},{value:"<code>ui [MOUNT_REF]</code>",id:"ui-mount_ref",children:[{value:"Arguments",id:"arguments-6",children:[]},{value:"Examples",id:"examples-9",children:[]}]}],p={rightToc:r},b="wrapper";function s(e){var t=e.components,n=Object(l.a)(e,["components"]);return Object(o.b)(b,Object(a.a)({},p,n,{components:t,mdxType:"MDXLayout"}),Object(o.b)("h2",{id:"opctl"},Object(o.b)("inlineCode",{parentName:"h2"},"opctl")),Object(o.b)("h3",{id:"options"},"Options"),Object(o.b)("h4",{id:"--no-color"},Object(o.b)("inlineCode",{parentName:"h4"},"--no-color")),Object(o.b)("p",null,"To disable color, include a ",Object(o.b)("inlineCode",{parentName:"p"},"--no-color")," flag w/ your\ncommand."),Object(o.b)("blockquote",null,Object(o.b)("p",{parentName:"blockquote"},"this may increase readability in environments not supporting\ncolor escape codes or piping output to another program")),Object(o.b)("h5",{id:"examples"},"Examples"),Object(o.b)("pre",null,Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"opctl --no-color events\n")),Object(o.b)("h4",{id:"-h-or---help"},Object(o.b)("inlineCode",{parentName:"h4"},"-h")," or ",Object(o.b)("inlineCode",{parentName:"h4"},"--help")),Object(o.b)("p",null,"For context specific help, include a ",Object(o.b)("inlineCode",{parentName:"p"},"-h")," (or ",Object(o.b)("inlineCode",{parentName:"p"},"--help"),") flag w/ your\ncommand."),Object(o.b)("h5",{id:"examples-1"},"Examples"),Object(o.b)("pre",null,Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"opctl node create -h\n\nUsage: opctl node create\n\nCreates a node\n")),Object(o.b)("h2",{id:"opctl-events"},Object(o.b)("inlineCode",{parentName:"h2"},"opctl events")),Object(o.b)("p",null,"listen to node events."),Object(o.b)("blockquote",null,Object(o.b)("p",{parentName:"blockquote"},"if a node isn't running, one will be automatically created.")),Object(o.b)("h3",{id:"examples-2"},"Examples"),Object(o.b)("h4",{id:"event-replay"},"Event Replay"),Object(o.b)("p",null,"Events are persisted to disk and can be replayed."),Object(o.b)("blockquote",null,Object(o.b)("p",{parentName:"blockquote"},"events are not held across node restarts; any time a node starts it\nclears its event db.")),Object(o.b)("ol",null,Object(o.b)("li",{parentName:"ol"},Object(o.b)("p",{parentName:"li"},"open terminal & run an op so we have some events"),Object(o.b)("pre",{parentName:"li"},Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"opctl run github.com/opspec-pkgs/uuid.v4.generate#1.1.0\n"))),Object(o.b)("li",{parentName:"ol"},Object(o.b)("p",{parentName:"li"},"exit terminal"),Object(o.b)("pre",{parentName:"li"},Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"exit\n"))),Object(o.b)("li",{parentName:"ol"},Object(o.b)("p",{parentName:"li"},"re open terminal & replay events"),Object(o.b)("pre",{parentName:"li"},Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"opctl run events\n")))),Object(o.b)("h4",{id:"event-streaming"},"Event Streaming"),Object(o.b)("p",null,"Events are streamed in realtime as they occur. They can be streamed in parallel to any number of terminals."),Object(o.b)("blockquote",null,Object(o.b)("p",{parentName:"blockquote"},"behind the scenes, events are delivered over websockets")),Object(o.b)("ol",null,Object(o.b)("li",{parentName:"ol"},Object(o.b)("p",{parentName:"li"},"open multiple terminals & open event stream on each"),Object(o.b)("pre",{parentName:"li"},Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"opctl events\n"))),Object(o.b)("li",{parentName:"ol"},Object(o.b)("p",{parentName:"li"},"open another terminal & run an op; watch events show up on all terminals simultaneously in real-time"),Object(o.b)("pre",{parentName:"li"},Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"opctl run some-op\n")))),Object(o.b)("h2",{id:"opctl-ls-dir_ref"},Object(o.b)("inlineCode",{parentName:"h2"},"opctl ls [DIR_REF]")),Object(o.b)("p",null,"List ops in a directory."),Object(o.b)("h3",{id:"arguments"},"Arguments"),Object(o.b)("h4",{id:"dir_ref-default-opspec"},Object(o.b)("inlineCode",{parentName:"h4"},"DIR_REF")," ",Object(o.b)("em",{parentName:"h4"},"default: ",Object(o.b)("inlineCode",{parentName:"em"},".opspec"))),Object(o.b)("p",null,"Reference to dir ops will be listed from (either ",Object(o.b)("inlineCode",{parentName:"p"},"relative/path"),", ",Object(o.b)("inlineCode",{parentName:"p"},"/absolute/path"),", ",Object(o.b)("inlineCode",{parentName:"p"},"host/path/repo#tag"),", or ",Object(o.b)("inlineCode",{parentName:"p"},"host/path/repo#tag/path"),")"),Object(o.b)("h3",{id:"examples-3"},"Examples"),Object(o.b)("h4",{id:"opspec-dir"},Object(o.b)("inlineCode",{parentName:"h4"},".opspec")," dir"),Object(o.b)("p",null,"lists ops from ",Object(o.b)("inlineCode",{parentName:"p"},"./.opspec")),Object(o.b)("pre",null,Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"opctl ls\n")),Object(o.b)("h4",{id:"remote-dir"},"remote dir"),Object(o.b)("p",null,"lists ops from ",Object(o.b)("a",Object(a.a)({parentName:"p"},{href:"https://github.com/opctl/opctl/tree/0.1.24"}),"github.com/opctl/opctl#0.1.24")),Object(o.b)("pre",null,Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"opctl ls github.com/opctl/opctl#0.1.24/\n")),Object(o.b)("h2",{id:"opctl-run-options-op_ref"},Object(o.b)("inlineCode",{parentName:"h2"},"opctl run [OPTIONS] OP_REF")),Object(o.b)("p",null,"Start and wait on an op"),Object(o.b)("blockquote",null,Object(o.b)("p",{parentName:"blockquote"},"if a node isn't running, one will be automatically created")),Object(o.b)("h3",{id:"arguments-1"},"Arguments"),Object(o.b)("h4",{id:"op_ref"},Object(o.b)("inlineCode",{parentName:"h4"},"OP_REF")),Object(o.b)("p",null,"Op reference (either ",Object(o.b)("inlineCode",{parentName:"p"},"relative/path"),", ",Object(o.b)("inlineCode",{parentName:"p"},"/absolute/path"),", ",Object(o.b)("inlineCode",{parentName:"p"},"host/path/repo#tag"),", or ",Object(o.b)("inlineCode",{parentName:"p"},"host/path/repo#tag/path"),")"),Object(o.b)("h3",{id:"options-1"},"Options"),Object(o.b)("h4",{id:"-a"},Object(o.b)("inlineCode",{parentName:"h4"},"-a")),Object(o.b)("p",null,"Explicitly pass args to op in format ",Object(o.b)("inlineCode",{parentName:"p"},"-a NAME1=VALUE1 -a NAME2=VALUE2")),Object(o.b)("h4",{id:"--arg-file-default-opspecargsyml"},Object(o.b)("inlineCode",{parentName:"h4"},"--arg-file")," ",Object(o.b)("em",{parentName:"h4"},"default: ",Object(o.b)("inlineCode",{parentName:"em"},".opspec/args.yml"))),Object(o.b)("p",null,"Read in a file of args in yml format"),Object(o.b)("h3",{id:"examples-4"},"Examples"),Object(o.b)("h4",{id:"local-op-ref-wout-args"},"local op ref w/out args"),Object(o.b)("pre",null,Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"opctl run myop\n")),Object(o.b)("h4",{id:"remote-op-ref-w-args"},"remote op ref w/ args"),Object(o.b)("pre",null,Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),'opctl run -a apiToken="my-token" -a channelName="my-channel" -a msg="hello!" github.com/opspec-pkgs/slack.chat.post-message#0.1.1\n')),Object(o.b)("h3",{id:"notes"},"Notes"),Object(o.b)("h4",{id:"op-source-usernamepassword-prompt"},"op source username/password prompt"),Object(o.b)("p",null,"If auth w/ the op source fails the cli will (re)prompt for username &\npassword."),Object(o.b)("blockquote",null,Object(o.b)("p",{parentName:"blockquote"},"in non-interactive terminals, the cli will note that it can't prompt\ndue to being in a non-interactive terminal and exit with a non zero\nexit code.")),Object(o.b)("h4",{id:"input-sources"},"input sources"),Object(o.b)("p",null,"Input sources are checked according to the following precedence:"),Object(o.b)("ul",null,Object(o.b)("li",{parentName:"ul"},"arg provided via ",Object(o.b)("inlineCode",{parentName:"li"},"-a")," option"),Object(o.b)("li",{parentName:"ul"},"arg file"),Object(o.b)("li",{parentName:"ul"},"env var"),Object(o.b)("li",{parentName:"ul"},"default"),Object(o.b)("li",{parentName:"ul"},"prompt")),Object(o.b)("h4",{id:"input-prompts"},"input prompts"),Object(o.b)("p",null,"Inputs which are invalid or missing will result in the cli prompting for\nthem."),Object(o.b)("blockquote",null,Object(o.b)("p",{parentName:"blockquote"},"in non-interactive terminals, the cli will provide details about the\ninvalid or missing input, note that it's giving up due to being in a\nnon-interactive terminal and exit with a non zero exit code.")),Object(o.b)("p",null,"example:"),Object(o.b)("pre",null,Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"\n-\n  Please provide value for parameter.\n  Name: version\n  Description: version of app being compiled\n-\n")),Object(o.b)("h5",{id:"validation"},"validation"),Object(o.b)("p",null,"When inputs don't meet constraints, the cli will (re)prompt for the\ninput until a satisfactory value is obtained."),Object(o.b)("h4",{id:"containers"},"containers"),Object(o.b)("h5",{id:"image"},"image"),Object(o.b)("h6",{id:"image-layer-caching"},"image layer caching"),Object(o.b)("p",null,"All pulled image layers will be cached"),Object(o.b)("h6",{id:"image-updates"},"image updates"),Object(o.b)("p",null,"Prior to container creation, updates to the referenced image will be\npulled and applied."),Object(o.b)("p",null,"If checking for or applying updated image layers fails, graceful\nfallback to cached image layers will occur"),Object(o.b)("h5",{id:"networking"},"networking"),Object(o.b)("p",null,"All containers created by opctl will be attached to a single managed\nnetwork."),Object(o.b)("blockquote",null,Object(o.b)("p",{parentName:"blockquote"},"the network is visible from ",Object(o.b)("inlineCode",{parentName:"p"},"docker network ls")," as ",Object(o.b)("inlineCode",{parentName:"p"},"opctl"),".")),Object(o.b)("h5",{id:"cleanup"},"cleanup"),Object(o.b)("p",null,"Containers will be removed as they exit."),Object(o.b)("h2",{id:"opctl-node-create"},Object(o.b)("inlineCode",{parentName:"h2"},"opctl node create")),Object(o.b)("p",null,"Create an in-process node which inherits current\nstderr/stdout/stdin/PGId (process group id) and blocks until killed."),Object(o.b)("blockquote",null,Object(o.b)("p",{parentName:"blockquote"},"There can be only one node running at a time on a given machine.")),Object(o.b)("h3",{id:"options-2"},"Options"),Object(o.b)("h4",{id:"--data-dir"},Object(o.b)("inlineCode",{parentName:"h4"},"--data-dir")),Object(o.b)("p",null,"Path of dir used to store node data"),Object(o.b)("h3",{id:"notes-1"},"Notes"),Object(o.b)("h4",{id:"lockfile"},"lockfile"),Object(o.b)("p",null,"Upon creation, nodes populate a lockfile at ",Object(o.b)("inlineCode",{parentName:"p"},"DATA_DIR/lockfile.pid"),"\ncontaining their PId (process id)."),Object(o.b)("h4",{id:"concurrency"},"concurrency"),Object(o.b)("p",null,"Prior to node creation, if a lockfile exists, the existing lock holder\nwill be liveness tested."),Object(o.b)("p",null,"Only in the event the existing lock holder is dead will creation of a\nnew node occur."),Object(o.b)("h4",{id:"debugging"},"debugging"),Object(o.b)("p",null,"Debugging can be accomplished by running ",Object(o.b)("inlineCode",{parentName:"p"},"node create")," from a terminal\nwhere it's output is easily monitored."),Object(o.b)("h4",{id:"cleanup-1"},"cleanup"),Object(o.b)("p",null,"During creation, ",Object(o.b)("inlineCode",{parentName:"p"},"DATA_DIR")," will be\ncleaned of any existing events, ops, and temp files/dirs to ensure\nthe created node starts from a clean slate."),Object(o.b)("h2",{id:"opctl-node-kill"},Object(o.b)("inlineCode",{parentName:"h2"},"opctl node kill")),Object(o.b)("p",null,"Kill the running node."),Object(o.b)("h2",{id:"opctl-op-create-options-name"},Object(o.b)("inlineCode",{parentName:"h2"},"opctl op create [OPTIONS] NAME")),Object(o.b)("p",null,"Creates an op"),Object(o.b)("h3",{id:"arguments-2"},"Arguments"),Object(o.b)("h4",{id:"name"},Object(o.b)("inlineCode",{parentName:"h4"},"NAME")),Object(o.b)("p",null,"Name of the op"),Object(o.b)("h3",{id:"options-3"},"Options"),Object(o.b)("h4",{id:"-d-or---description"},Object(o.b)("inlineCode",{parentName:"h4"},"-d")," or ",Object(o.b)("inlineCode",{parentName:"h4"},"--description")),Object(o.b)("p",null,"Description of the op"),Object(o.b)("h4",{id:"--path-default-opspec"},Object(o.b)("inlineCode",{parentName:"h4"},"--path")," ",Object(o.b)("em",{parentName:"h4"},"default: ",Object(o.b)("inlineCode",{parentName:"em"},".opspec"))),Object(o.b)("p",null,"Path to create the op at"),Object(o.b)("h3",{id:"examples-5"},"Examples"),Object(o.b)("pre",null,Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),'opctl op create -d "my awesome op description" --path some/path my-awesome-op-name\n')),Object(o.b)("h2",{id:"op-install-options-op_ref"},Object(o.b)("inlineCode",{parentName:"h2"},"op install [OPTIONS] OP_REF")),Object(o.b)("p",null,"Installs an op"),Object(o.b)("h3",{id:"arguments-3"},"Arguments"),Object(o.b)("h4",{id:"op_ref-1"},Object(o.b)("inlineCode",{parentName:"h4"},"OP_REF")),Object(o.b)("p",null,"Op reference (",Object(o.b)("inlineCode",{parentName:"p"},"host/path/repo#tag"),", or ",Object(o.b)("inlineCode",{parentName:"p"},"host/path/repo#tag/path"),")"),Object(o.b)("h3",{id:"options-4"},"Options"),Object(o.b)("h4",{id:"--path-default-opspecop_ref"},Object(o.b)("inlineCode",{parentName:"h4"},"--path")," ",Object(o.b)("em",{parentName:"h4"},"default: ",Object(o.b)("inlineCode",{parentName:"em"},".opspec/OP_REF"))),Object(o.b)("p",null,"Path to install the op at"),Object(o.b)("h4",{id:"-u-or---username"},Object(o.b)("inlineCode",{parentName:"h4"},"-u")," or ",Object(o.b)("inlineCode",{parentName:"h4"},"--username")),Object(o.b)("p",null,"Username used to auth w/ the op source"),Object(o.b)("h4",{id:"-p-or---password"},Object(o.b)("inlineCode",{parentName:"h4"},"-p")," or ",Object(o.b)("inlineCode",{parentName:"h4"},"--password")),Object(o.b)("p",null,"Password used to auth w/ the op source"),Object(o.b)("h3",{id:"examples-6"},"Examples"),Object(o.b)("pre",null,Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"opctl op install -u someUser -p somePass host/path/repo#tag\n")),Object(o.b)("h3",{id:"notes-2"},"Notes"),Object(o.b)("h4",{id:"op-source-usernamepassword-prompt-1"},"op source username/password prompt"),Object(o.b)("p",null,"If auth w/ the op source fails the cli will (re)prompt for username &\npassword."),Object(o.b)("blockquote",null,Object(o.b)("p",{parentName:"blockquote"},"in non-interactive terminals, the cli will note that it can't prompt\nand exit with a non zero exit code.")),Object(o.b)("h2",{id:"op-kill-op_id"},Object(o.b)("inlineCode",{parentName:"h2"},"op kill OP_ID")),Object(o.b)("p",null,"Kill an op. "),Object(o.b)("h3",{id:"arguments-4"},"Arguments"),Object(o.b)("h4",{id:"op_id"},Object(o.b)("inlineCode",{parentName:"h4"},"OP_ID")),Object(o.b)("p",null,"Id of the op to kill"),Object(o.b)("h2",{id:"op-validate-op_ref"},Object(o.b)("inlineCode",{parentName:"h2"},"op validate OP_REF")),Object(o.b)("p",null,"Validates an op according to:"),Object(o.b)("ul",null,Object(o.b)("li",{parentName:"ul"},"existence of ",Object(o.b)("inlineCode",{parentName:"li"},"op.yml")),Object(o.b)("li",{parentName:"ul"},"validity of ",Object(o.b)("inlineCode",{parentName:"li"},"op.yml")," (per\n",Object(o.b)("a",Object(a.a)({parentName:"li"},{href:"https://opctl.io/0.1.6/op.yml.schema.json"}),"schema"),")")),Object(o.b)("h3",{id:"arguments-5"},"Arguments"),Object(o.b)("h4",{id:"op_ref-2"},Object(o.b)("inlineCode",{parentName:"h4"},"OP_REF")),Object(o.b)("p",null,"Op reference (either ",Object(o.b)("inlineCode",{parentName:"p"},"relative/path"),", ",Object(o.b)("inlineCode",{parentName:"p"},"/absolute/path"),", ",Object(o.b)("inlineCode",{parentName:"p"},"host/path/repo#tag"),", or ",Object(o.b)("inlineCode",{parentName:"p"},"host/path/repo#tag/path"),")."),Object(o.b)("h3",{id:"examples-7"},"Examples"),Object(o.b)("pre",null,Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"opctl op validate myop\n")),Object(o.b)("h3",{id:"notes-3"},"Notes"),Object(o.b)("h4",{id:"op-source-usernamepassword-prompt-2"},"op source username/password prompt"),Object(o.b)("p",null,"If auth w/ the op source fails the cli will (re)prompt for username & password."),Object(o.b)("blockquote",null,Object(o.b)("p",{parentName:"blockquote"},"in non-interactive terminals, the cli will note that it can't prompt and exit with a non zero exit code.")),Object(o.b)("h2",{id:"opctl-self-update-options"},Object(o.b)("inlineCode",{parentName:"h2"},"opctl self-update [OPTIONS]")),Object(o.b)("p",null,"Updates the current version of opctl."),Object(o.b)("blockquote",null,Object(o.b)("p",{parentName:"blockquote"},"if a node is running, it will be automatically killed")),Object(o.b)("h3",{id:"options-5"},"Options"),Object(o.b)("h5",{id:"-c-or---channel-default-stable"},Object(o.b)("inlineCode",{parentName:"h5"},"-c")," or ",Object(o.b)("inlineCode",{parentName:"h5"},"--channel")," ",Object(o.b)("em",{parentName:"h5"},"default: ",Object(o.b)("inlineCode",{parentName:"em"},"stable"))),Object(o.b)("p",null,"The release channel to update from"),Object(o.b)("ul",null,Object(o.b)("li",{parentName:"ul"},Object(o.b)("inlineCode",{parentName:"li"},"stable")),Object(o.b)("li",{parentName:"ul"},Object(o.b)("inlineCode",{parentName:"li"},"beta")," (smoke tested alpha channel)"),Object(o.b)("li",{parentName:"ul"},Object(o.b)("inlineCode",{parentName:"li"},"alpha")," (all bets are off)")),Object(o.b)("h3",{id:"examples-8"},"Examples"),Object(o.b)("p",null,"get latest stable release"),Object(o.b)("pre",null,Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"opctl self-update\n# output: Updated to new version: 0.1.24!\n")),Object(o.b)("p",null,"play around w/ latest beta release"),Object(o.b)("pre",null,Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"opctl self-update -c beta\n# output: Updated to new version: 0.1.24-beta.1!\n")),Object(o.b)("p",null,"play times over; switch back to latest stable release"),Object(o.b)("pre",null,Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"opctl self-update\n# output: Updated to new version: 0.1.24!\n")),Object(o.b)("h2",{id:"ui-mount_ref"},Object(o.b)("inlineCode",{parentName:"h2"},"ui [MOUNT_REF]")),Object(o.b)("p",null,"Opens the opctl web UI to the current working directory."),Object(o.b)("h3",{id:"arguments-6"},"Arguments"),Object(o.b)("h4",{id:"mount_ref-default-"},Object(o.b)("inlineCode",{parentName:"h4"},"MOUNT_REF")," ",Object(o.b)("em",{parentName:"h4"},"default: ",Object(o.b)("inlineCode",{parentName:"em"},"."))),Object(o.b)("p",null,"Reference to mount (either ",Object(o.b)("inlineCode",{parentName:"p"},"relative/path"),", ",Object(o.b)("inlineCode",{parentName:"p"},"/absolute/path"),", ",Object(o.b)("inlineCode",{parentName:"p"},"host/path/repo#tag"),", or ",Object(o.b)("inlineCode",{parentName:"p"},"host/path/repo#tag/path"),")."),Object(o.b)("h3",{id:"examples-9"},"Examples"),Object(o.b)("p",null,"Open web UI to current working directory"),Object(o.b)("pre",null,Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"opctl ui\n")),Object(o.b)("p",null,"Open web UI to remote op (github.com/opspec-pkgs/_.op.create#3.3.1)"),Object(o.b)("pre",null,Object(o.b)("code",Object(a.a)({parentName:"pre"},{className:"language-sh"}),"opctl ui github.com/opspec-pkgs/_.op.create#3.3.1\n")))}s.isMDXComponent=!0},197:function(e,t,n){"use strict";n.d(t,"a",(function(){return s})),n.d(t,"b",(function(){return O}));var a=n(0),l=n.n(a);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function c(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?c(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):c(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function r(e,t){if(null==e)return{};var n,a,l=function(e,t){if(null==e)return{};var n,a,l={},o=Object.keys(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||(l[n]=e[n]);return l}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(l[n]=e[n])}return l}var p=l.a.createContext({}),b=function(e){var t=l.a.useContext(p),n=t;return e&&(n="function"==typeof e?e(t):i({},t,{},e)),n},s=function(e){var t=b(e.components);return(l.a.createElement(p.Provider,{value:t},e.children))},d="mdxType",u={inlineCode:"code",wrapper:function(e){var t=e.children;return l.a.createElement(l.a.Fragment,{},t)}},m=Object(a.forwardRef)((function(e,t){var n=e.components,a=e.mdxType,o=e.originalType,c=e.parentName,p=r(e,["components","mdxType","originalType","parentName"]),s=b(n),d=a,m=s["".concat(c,".").concat(d)]||s[d]||u[d]||o;return n?l.a.createElement(m,i({ref:t},p,{components:n})):l.a.createElement(m,i({ref:t},p))}));function O(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=n.length,c=new Array(o);c[0]=m;var i={};for(var r in t)hasOwnProperty.call(t,r)&&(i[r]=t[r]);i.originalType=e,i[d]="string"==typeof e?e:a,c[1]=i;for(var p=2;p<o;p++)c[p]=n[p];return l.a.createElement.apply(null,c)}return l.a.createElement.apply(null,n)}m.displayName="MDXCreateElement"}}]);
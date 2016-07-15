package htmllog

var DefaultStyle = `
body {
	font-family: menlo, monospace;
	font-size: 10pt;
}
.time {
	display:inline;
	color: #DAD5D5;
	margin-right:3px;
}
.info {
	display:inline;
	color: #879BB5;
	margin-right:3px;
}
.infomsg {	
	display:inline;
}
.warn {
	display:inline;
	color: #A127A0;
	margin-right:2px;
}
.warnmsg {
	display:inline;
	color: #A127A0;
}
.debug {
	display:inline;
	color: #7FB582;
	margin-right:2px;
}
.debugmsg {
	display:inline;
}
.error {
	display:inline;
	color: #FF0404;
	margin-right:2px;
}
.fatal {
	display:inline;
	color: #FF0000;
	margin-right:2px;
}
.errormsg {
	display:inline;
	color: #FF0404;
}
`

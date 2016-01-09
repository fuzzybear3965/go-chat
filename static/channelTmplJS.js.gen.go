// AUTOMATICALLY GENERATED FILE. DO NOT EDIT.

package static

var ChannelTmplJS = js(asset{Name: "ChannelTmplJS.js", Content: "" +
	"var textarea = document.getElementById(\"messageArea\");\r\nvar chatForm = document.getElementById(\"chatForm\");\r\n\r\ntry {\r\n    textarea.addEventListener(\"keydown\", keyPress, false);\r\n} catch(e) {\r\n    textarea.attachEvent(\"onkeydown\", keypress);\r\n}\r\n\r\nfunction keyPress(e) {\r\n    if (e.keyCode === 13) {\r\n        chatForm.submit()\r\n    } else {\r\n        return;\r\n    }\r\n}\r\n\r\nfunction setFocusToTextBox(){\r\n    document.getElementById(\"messageArea\").focus();\r\n}\r\n" +
	"", etag: `"E1zfq1UxQDo="`})

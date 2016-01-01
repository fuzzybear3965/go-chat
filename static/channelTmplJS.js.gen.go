// AUTOMATICALLY GENERATED FILE. DO NOT EDIT.

package static

var ChannelTmplJS = js(asset{Name: "ChannelTmplJS.js", Content: "" +
	"var textarea = document.getElementById(\"messageArea\");\nvar chatForm = document.getElementById(\"chatForm\");\n\ntry {\n    textarea.addEventListener(\"keydown\", keyPress, false);\n} catch(e) {\n    textarea.attachEvent(\"onkeydown\", keypress);\n}\n\nfunction keyPress(e) {\n    if (e.keyCode === 13) {\n        chatForm.submit()\n    } else {\n        return;\n    }\n}\n" +
	"", etag: `"xCZegz8/mXo="`})

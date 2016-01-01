// AUTOMATICALLY GENERATED FILE. DO NOT EDIT.

package static

var ChannelTemplate = tmpl(asset{Name: "ChannelTemplate.html.tmpl", Content: "" +
	"<html>\n    <head>\n        <style>\n            {{.Template.ChannelTmplCSS}}\n        </style>\n    </head>\n    <body onload=\"setFocusToTextBox()\">\n        <h1>Welcome to the {{toLower .Channel.ChannelName | printf \"%s\"}} channel!</h1>\n        <p> Go ahead and start chatting below. Just hit enter to submit your\n            text. \n        </p>\n\n        <form action=\"/c/{{.Channel.ChannelName}}\" method=\"POST\" id=\"chatForm\">\n            <textarea name=\"message\" id=\"messageArea\" rows=\"5\" cols=\"40\"></textarea>\n        </form>\n\n        <pre> {{printf \"%s\" .Channel.Body}} </pre>\n    </body>\n\n    <script>\n        {{.Template.ChannelTmplJS}}\n    </script>\n</html>\n" +
	"", etag: `"N2Ad1YD48x0="`})

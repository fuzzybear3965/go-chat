// AUTOMATICALLY GENERATED FILE. DO NOT EDIT.

package static

var ChannelTemplate = html(asset{Name: "ChannelTemplate.html", Content: "" +
	"<html>\r\n    <head>\r\n        <style>\r\n            {{.Template.ChannelTmplCSS}}\r\n        </style>\r\n    </head>\r\n    <body onload=\"setFocusToTextBox()\">\r\n        <h1>Welcome to the {{toLower .ChannelName | printf \"%s\"}} channel!</h1>\r\n        <p> Go ahead and start chatting below. Just hit enter to submit your\r\n            text. \r\n        </p>\r\n\r\n        <form action=\"/c/{{.ChannelName}}\" method=\"POST\" id=\"chatForm\">\r\n            <textarea name=\"message\" id=\"messageArea\" rows=\"5\" cols=\"40\"></textarea>\r\n        </form>\r\n\r\n        <pre> {{.ChannelLog}} </pre>\r\n    </body>\r\n\r\n    <script>\r\n        {{.Template.ChannelTmplJS}}\r\n    </script>\r\n</html>\r\n" +
	"", etag: `"Jq9a/r0GpYU="`})

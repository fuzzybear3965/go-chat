// AUTOMATICALLY GENERATED FILE. DO NOT EDIT.

package static

var ChannelTemplate = tmpl(asset{Name: "ChannelTemplate.html.tmpl", Content: "" +
	"<html>\n    <head>\n        <style>\n            {{.Template.ChannelTmplCSS}}\n        </style>\n    </head>\n    <body>\n        <h1>Welcome to the {{printf \"%s\" .Channel.ChannelName}} channel!</h1>\n        <p> Go ahead and start chatting below. </p>\n        <p>CSS Data : {{.Template.ChannelTmplCSS}} </p>\n\n        <form action=\"/c/{{.Channel.ChannelName}}\" method=\"POST\" id=\"chatForm\">\n            <textarea name=\"message\" id=\"messageArea\" rows=\"5\" cols=\"40\"></textarea>\n            <input type=\"submit\" value=\"Submit\">\n        </form>\n\n        <pre> {{printf \"%s\" .Channel.Body}} </pre>\n    </body>\n\n    <script>\n        {{.Template.ChannelTmplJS}}\n    </script>\n</html>\n" +
	"", etag: `"gSCKUkdCT8U="`})

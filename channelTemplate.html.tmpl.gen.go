// AUTOMATICALLY GENERATED FILE. DO NOT EDIT.

package main

var channelTemplate = tmpl(asset{Name: "channelTemplate.html.tmpl", Content: "" +
	"<html>\r\n    <head>\r\n        <style>\r\n            {{$static.channelTmplCSS}}\r\n        </style>\r\n    </head>\r\n    <body>\r\n        <h1>Welcome to the {{.ChannelName}} channel!</h1>\r\n        <p> Go ahead and start chatting below. </p>\r\n\r\n        <form action=\"/c/{{.ChannelName}}\" method=\"POST\" id=\"chatForm\">\r\n            <textarea name=\"message\" id=\"messageArea\" rows=\"5\" cols=\"40\"></textarea>\r\n            <input type=\"submit\" value=\"Submit\">\r\n        </form>\r\n\r\n        <pre> {{printf \"%s\" .Body}} </pre>\r\n    </body>\r\n\r\n    <script>\r\n        {{$static.channelTmplJS}}\r\n    </script>\r\n</html>\r\n" +
	"", etag: `"IZ040r/mynY="`})

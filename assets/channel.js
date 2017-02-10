var textarea = document.getElementById("messageArea");
var chatForm = document.getElementById("chatForm");

try {
    textarea.addEventListener("keydown", keyPress, false);
} catch(e) {
    textarea.attachEvent("onkeydown", keypress);
}

var ws = new WebSocket("ws" + document.URL.slice(4,document.URL.length))

ws.addEventListener("message", wsHandler(event))

// Use function hoisting
function keyPress(e) {
    if (e.keyCode === 13) {
        chatForm.submit()
    } else {
        return;
    }
}

function setFocusToTextBox(){
    document.getElementById("messageArea").focus();
}

function wsHandler(evt) {
   console.log(evt)
}

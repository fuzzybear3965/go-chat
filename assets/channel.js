var textarea = document.getElementById("messageArea");
var chatForm = document.getElementById("chatForm");

try {
    textarea.addEventListener("keydown", keyPress, false);
} catch(e) {
    textarea.attachEvent("onkeydown", keypress);
}

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

var ws = new WebSocket("ws://127.0.0.1/c/{{.ChannelName}}")

window.onbeforeunload = function() {
    websocket.onclose = function () {}; // disable onclose handler first
    websocket.close()
};

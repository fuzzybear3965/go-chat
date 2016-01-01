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

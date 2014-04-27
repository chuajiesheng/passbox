function getHTML(id) {
    return document.getElementById(id).innerHTML;
}

function setHTML(id, html) {
    document.getElementById(id).innerHTML = html;
}

function getValue(id) {
    return document.getElementById(id).value;
}

function setValue(id, value) {
    document.getElementById(id).value = value;
}

function notification(caption, content) {
    var not = $.Notify({
        style: {background: "red", color: "white"},
        caption: caption,
        content: content,
        timeout: 10000 // 10 seconds,
    });
}

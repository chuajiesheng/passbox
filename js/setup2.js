function setError(id) {
    document.getElementById(id).className += " error-state";
}

function removeError(id) {
    document.getElementById(id).className = "input-control size4 password";
}

function complexityPassed(id) {
    regex = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[^a-zA-Z0-9])(?!.*\s).{8,15}$/;
    return getValue(id).match(regex);
}

function updateUi(element) {
    row = parseInt(element.id.substring(element.id.length - 1, element.id.length));
    pw = getValue('pw' + row);
    if (pw.length < 8) {
        setError('control' + row);

        var not = $.Notify({
            style: {background: "red", color: "white"},
            caption: "Password length is insufficient!",
            content: "Please provider a stronger password.",
            timeout: 10000 // 10 seconds,
        });

    } else if (!complexityPassed('pw' + row)) {
        setError('control' + row);

        var not = $.Notify({
            style: {background: "red", color: "white"},
            caption: "Password complexity insufficient!",
            content: "Please use upper case letters, lower case letters, digits and symbols.",
            timeout: 10000 // 10 seconds,
        });

    } else {
        removeError('control' + row);
    }
}

function validation() {
    var pw1 = getValue("pw1");
    var pw2 = getValue("pw2");

    if (pw1 != pw2) {
        notification("Differences between password fields detected.", "Please verify entry.");
        return false;
    } else {
        return true;
    }
}

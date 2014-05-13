function setError(id) {
    document.getElementById(id).className += " error-state";
}

function removeError(id) {
    document.getElementById(id).className = "input-control size4 password";
}

function complexityPassed(pw) {
    regex = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[^a-zA-Z0-9])(?!.*\s).{8,15}$/;
    return pw.match(regex);
}

function updateUi(element) {
    row = parseInt(element.id.substring(element.id.length - 1, element.id.length));
    pw = getValue('pw' + row);
    if (pw.length < 8) {
        setError('control' + row);
        notification("Password length is insufficient!",
                     "Please provider a stronger password.");
    } else if (!complexityPassed(getValue('pw' + row))) {
        setError('control' + row);
        notification("Password complexity insufficient!",
                     "Please use upper case letters, lower case letters, digits and symbols.")
    } else {
        removeError('control' + row);
    }
}

function validation() {
    var pw1 = getValue("pw1");
    var pw2 = getValue("pw2");

    if (pw1.length < 1 || pw2.length < 1) {
        notification("Either of the fields are empty.",
                     "Please fill in both entry.");
        return false;
    } else if (pw1 != pw2) {
        notification("Differences between password fields detected.",
                     "Please verify entry.");
        return false;
    } else if (pw1.length < 8 || pw2.length < 8) {
        notification("Password length is insufficient!",
                     "Please provider a stronger password.");
        return false;
    } else if (!complexityPassed(pw1)) {
        notification("Password complexity insufficient!",
                     "Please use upper case letters, lower case letters, digits and symbols.")
        return false;
    }

    return true;
}

function getContent(array) {
    content = "Please fill in: ";

    for (i = 0; i < array.length; i++) {
        content += array[i];
        if (i < (array.length - 1)) {
            content += ", ";
        }
    }

    return content;
}

function validation() {
    var emptyArray = [];

    var key1 = getValue("key1");
    var username = getValue("username");
    var pw1 = getValue("pw1");
    var pw2 = getValue("pw2");

    if (key1.length < 1) {
        emptyArray.push("Credential Key");
    }

    if (username.length < 1) {
        emptyArray.push("Username");
    }

    if (pw1.length < 1 || pw2.length < 1) {
        emptyArray.push("Password");
    }


    if (emptyArray.length > 0) {
        notification("Empty fields detected!", getContent(emptyArray));
    } else if (pw1 != pw2) {
        notification("Differences between password fields detected.", "Please verify entry.");
    } else {
        return true;
    }

    return false;
}

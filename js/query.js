function validation() {
    var key1 = getValue("key1");
    var key2 = getValue("key2");

    if (key1.length < 1) {
        notification("Empty field detected!", "Please fill in credential key.");
        return false;
    }

    return true;
}

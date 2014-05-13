function removeDisabled(no) {
    document.getElementById('qns' + no).disabled = false;
    document.getElementById('ans' + no).disabled = false;
}

function updateUi(element) {
    row = parseInt(element.id.substring(element.id.length - 1, element.id.length));
    qns = getValue('qns' + row);
    ans = getValue('ans' + row);
    if (qns.length > 0 && ans.length > 0) {
        removeDisabled(row + 1)
    }
}

function check(row) {
    qns = getValue('qns' + row);
    ans = getValue('ans' + row);
    return qns.length > 0 && ans.length > 0;
}

function validation() {
    var start = 1;
    var limit = 4;
    var count = 0;

    for (var i = start; i <= limit; i++) {
        if (check(i)) {
            count++;
        }
    }

    if (count < 1) {
        notification("Empty question/answer pair.", "Please fill in the details.")
    }

    return count > 0;
}

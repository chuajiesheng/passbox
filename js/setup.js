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

function moveToInputEnd(id) {
  let el = document.getElementById(id);
  let end = el.value.length;
  el.setSelectionRange(end, end);
  el.focus();
}

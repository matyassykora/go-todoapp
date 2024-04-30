'use strict';

const moveToInputEnd = (id) => {
  let el = document.getElementById(id);
  let end = el.value.length;
  el.setSelectionRange(end, end);
  el.focus();
};

htmx.onLoad(function (content) {
  var sortables = content.querySelectorAll('.sortable');
  for (var i = 0; i < sortables.length; i++) {
    var sortable = sortables[i];
    var sortableInstance = new Sortable(sortable, {
      animation: 150,
      swapThreshold: 0.7,
      handle: '.handle',
      ghostClass: 'blue-background-class',

      // Disable sorting on the `end` event
      onEnd: function (evt) {
        this.option('disabled', true);
      },
    });

    // Re-enable sorting on the `htmx:afterSwap` event
    sortable.addEventListener('htmx:afterSwap', function () {
      sortableInstance.option('disabled', false);
    });
  }
});

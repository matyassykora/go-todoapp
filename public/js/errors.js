'use strict';

const errorTarget = document.getElementById('htmx-alert');

const closeErrorMessage = () => {
  errorTarget.setAttribute('hidden', 'true');
  errorTarget.classList.remove('d-inline-flex');
  errorTarget.innerText = '';
};

const closeBtn = document.createElement('button');
closeBtn.innerText = 'OK';
closeBtn.classList.add('btn', 'btn-danger');
closeBtn.onclick = closeErrorMessage;

document.body.addEventListener('htmx:afterRequest', function (evt) {
  errorTarget.classList.add('d-inline-flex');

  if (evt.detail.successful) {
    // Successful request, clear out alert
    closeErrorMessage();
  } else if (evt.detail.failed && evt.detail.xhr) {
    // Server error with response contents, equivalent to htmx:responseError
    const xhr = evt.detail.xhr;
    errorTarget.innerText = `${xhr.response}`;
    errorTarget.removeAttribute('hidden');
  } else {
    // Unspecified failure, usually caused by network error
    console.error('Unexpected htmx error', evt.detail);
    errorTarget.innerText =
      'Unexpected error, check your connection and try to refresh the page.';
    errorTarget.removeAttribute('hidden');
  }

  errorTarget.append(closeBtn);
});

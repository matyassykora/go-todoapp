'use strict';

const getStoredTheme = () => localStorage.getItem('theme');
const setStoredTheme = (theme) => localStorage.setItem('theme', theme);

const loadTheme = () => {
  document.documentElement.setAttribute('data-bs-theme', getStoredTheme());
};

const toggleTheme = () => {
  if (getStoredTheme() == 'dark') {
    setStoredTheme('light');
  } else {
    setStoredTheme('dark');
  }
  loadTheme();
};

const themeCheck = document.getElementById('theme-toggle');
const label = themeCheck.nextElementSibling;

themeCheck.addEventListener('input', (e) => {
  const checked = e.target.checked;

  if (checked) {
    label.innerText = '☀️';
  } else {
    label.innerText = '🌙';
  }
});

/**
 * @param {Element} element 
 */
const setInner = (element) => {
  const theme = getStoredTheme();
  if (theme === 'dark') {
    element.innerText = '☀️';
    themeCheck.checked = true;
  } else {
    element.innerText = '🌙';
  }
};

loadTheme();
setInner(label);

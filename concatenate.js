const fs = require('fs-extra');
const concat = require('concat');

const concatenate = async () => {
  const files = [
    './dist/survey-widget/runtime.js',
    './dist/survey-widget/polyfills.js',
    './dist/survey-widget/es2015-polyfills.js',
    './dist/survey-widget/scripts.js',
    './dist/survey-widget/main.js',
  ];

  await fs.ensureDir('output');
  await concat(files, 'output/contact-form.js');
};

concatenate();

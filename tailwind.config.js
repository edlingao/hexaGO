/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./web/components/**/*.templ",
    "./web/views/**/*.templ",
  ],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ['halloween'],
  },
  plugins: [
    require('daisyui'),
  ],
}


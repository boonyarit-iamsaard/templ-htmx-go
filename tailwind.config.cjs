const defaultTheme = require('tailwindcss/defaultTheme');

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./views/**/*.templ', './**/*.templ'],
  theme: {
    container: {
      center: true,
    },
    extend: {
      fontFamily: {
        sans: ['Noto Sans Thai', ...defaultTheme.fontFamily.sans],
      },
    },
  },
  plugins: [],
};

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.templ", "./assets/css/tailwind.input.css"],
  theme: {
    extend: {
      fontFamily: {
        base: ["Inconsolata", "ui-serif"],
        kr: ["Noto Serif KR", "ui-serif"],
      },
    },
  },
  plugins: [],
};

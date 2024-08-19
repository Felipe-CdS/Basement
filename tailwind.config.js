/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/*.templ", "./assets/css/tailwind.input.css"],
  theme: {
    extend: {
      fontFamily: {
        kr: ["NotoSerifKR", "ui-serif"],
      },
    },
  },
  plugins: [],
};

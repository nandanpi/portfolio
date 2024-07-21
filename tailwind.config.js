/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
  theme: {
    extend: {
      colors: {
        primary: "#ffc86a",
      },
    },
  },
  plugins: [],
};

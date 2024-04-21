/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./ui/**/*.{html,js}"],
  theme: {
    extend: {
      fontFamily: {
          sans: ["Iosevka Aile Iaso", "sans-serif"],
          mono: ["Iosevka Curly Iaso", "monospace"],
          serif: ["Iosevka Etoile Iaso", "serif"],
      },
    },
  },
  plugins: [
    require('tailwindcss'),
    require('autoprefixer'),
  ],
}


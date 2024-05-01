/** @type {import('tailwindcss').Config} */
export const content = ["./ui/**/*.{html,js}"];
export const theme = {
  extend: {
    fontFamily: {
      sans: ["Iosevka Aile Iaso", "sans-serif"],
      mono: ["Iosevka Curly Iaso", "monospace"],
      serif: ["Iosevka Etoile Iaso", "serif"],
    },
  },
};
export const plugins = [
  require('tailwindcss'),
  require('autoprefixer'),
];


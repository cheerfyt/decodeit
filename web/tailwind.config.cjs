/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/*.{ts,svelte}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
};

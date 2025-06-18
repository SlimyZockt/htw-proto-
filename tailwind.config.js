
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./include_dir/*.{html,js}", "*.{go,templ}"],
  theme: {
    extend: {},
  },
  //plugins: [daisyui, addIconSelectors(["tabler"])],
  //daisyui: {
  //  themes: ["light", "dark", "business"],
  //  darkTheme: "business",
  //},
};

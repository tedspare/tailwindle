const colors = require("tailwindcss/colors");

module.exports = {
  content: ["./src/**/*.svelte", "./src/**/*.js", "./src/**/*.ts"],
  theme: {
    colors: {
      gray: colors.neutral,
      white: colors.white,
      black: colors.black,
      green: colors.green,
      yellow: colors.amber,
      blue: colors.blue,
      red: colors.red,
    },
    extend: {},
  },
  plugins: [],
};

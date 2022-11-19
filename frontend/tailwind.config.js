/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      colors: {
        background: "#46178F",
        foreground: "#3F1481",
        accent: "#FF9551",
        primary: "#EEE",
        secondary: "#2f2b2b",
      },
      boxShadow: {},
    },
  },
  plugins: [],
};

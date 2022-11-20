/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      colors: {
        background: "#46178F",
        foreground: "#3F1481",
        accent: "#FF9551",
        accent2: "#864BBF",
        primary: "#E9D2FF",
        secondary: "#250049",
      },
      boxShadow: {},
    },
  },
  plugins: [],
};

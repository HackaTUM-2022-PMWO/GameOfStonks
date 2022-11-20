/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      colors: {
        background: "#46178F",
        foreground: "#FF9551",
        accent: "#FF9551",
        accent2: "#864BBF",
        primary: "#EEE",
        secondary: "#2f2b2b",
      },
      boxShadow: {},
    },
  },
  plugins: [],
};

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      colors: {
        background: "#5F43D0",
        foreground: "#F4F4F4",
        accent: "#FF9551",
        primary: "#EEE",
        secondary: "#EFEFEF",
      },
      boxShadow: {},
    },
  },
  plugins: [],
};

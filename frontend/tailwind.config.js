/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors:{
        'mint': '#B9FFF8',
        'teal': '#6FEDD6',
        'orange': '#FF9551',
        'red': '#FF4A4A',
      },
    },
  },
  plugins: [],
}

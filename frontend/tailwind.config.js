/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors:{
        'background': '#B9FFF8',
        'foreground': '#6FEDD6',
        'accent': '#FF9551',
        'primary': '#26263b',
        'secondary': '#FFFFFF',
      },
    },
  },
  plugins: [],
}

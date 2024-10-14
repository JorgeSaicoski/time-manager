/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './src/**/*.{html,js,svelte}', 
  ],
  theme: {
    extend: {
      fontFamily: {
        nerd: ['Martian', 'sans-serif'],
        body: ['Roboto', 'sans-serif'],
      },
      colors: {
        primary: '#00A8A8',
        secondary: '#2C3E50',
        accent: '#ECF0F1',
        hover: '#16D9D9',
        error: '#E74C3C',     
      },
    },
  },
  plugins: [
  ],
}

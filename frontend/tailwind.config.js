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
        primary: '#008B8B',     
        secondary: '#1F2E40',  
        accent: '#FFC107',     
        hover: '#00CED1',  
        error: '#C0392B',   
      },
    },
  },
  plugins: [
  ],
}

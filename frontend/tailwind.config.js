/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './src/**/*.{html,js,svelte}', 
  ],
  theme: {
    extend: {
      fontFamily: {
        nerd: ['Martian', 'sans-serif'],
        body: ['Nunito', 'sans-serif'],
      },
    },
  },
  plugins: [
  ],
}

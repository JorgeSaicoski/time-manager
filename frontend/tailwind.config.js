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
        secondaryAccent: '#354F6F',
        accent: '#FFC107',     
        hover: '#20B2AA',  
        error: '#C0392B',   

        textPrimary: '#FFFFFF',   
        textSecondary: '#B0BEC5', 
        textAccent: '#FFA000', 
        textError: '#E74C3C', 
        textDark: '#000000',

        buttonPrimaryBg: '#007A74',
        buttonPrimaryText: '#F0F8FF',
        buttonAccentBg: '#FFC93C',
        buttonAccentText: '#1A2430',
        buttonHoverBg: '#00CED1',     
        buttonErrorBg: '#C0392B',     
        buttonErrorText: '#F5F5F5',
      },
    },
  },
  plugins: [
  ],
}

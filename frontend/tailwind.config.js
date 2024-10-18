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

        textPrimary: '#FFFFFF',   
        textSecondary: '#B0BEC5', 
        textAccent: '#FFC107', 
        textError: '#E74C3C', 

        buttonPrimaryBg: '#008B8B', 
        buttonPrimaryText: '#FFFFFF', 
        buttonAccentBg: '#FFC107',    
        buttonAccentText: '#1F2E40', 
        buttonHoverBg: '#00CED1',    
        buttonErrorBg: '#C0392B',  
        buttonErrorText: '#FFFFFF',
      },
    },
  },
  plugins: [
  ],
}

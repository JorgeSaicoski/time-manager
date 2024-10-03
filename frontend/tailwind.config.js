/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './src/**/*.{svelte,js,ts}',  // Include TypeScript files
    './node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}'
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('flowbite/plugin')
  ],
}

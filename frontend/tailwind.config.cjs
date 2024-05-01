/** @type {import('tailwindcss').Config}*/
const vitePreprocess = import('@sveltejs/vite-plugin-svelte').then(m => m.vitePreprocess())
const config = {
  content: [
    "./src/**/*.{html,js,svelte,ts}",
    "./node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}",
    "./node_modules/flowbite-svelte-icons/**/*.{html,js,svelte,ts}",
  ],
  

  theme: {
    extend: {},
  },

  plugins: [
    require('flowbite/plugin')
  ],
  darkMode: 'class',
};


module.exports = {
    preprocess: {
        script:async (options) => (await vitePreprocess).script(options),
        style:async (options) => (await vitePreprocess).style(options),
    }
}
module.exports = config;

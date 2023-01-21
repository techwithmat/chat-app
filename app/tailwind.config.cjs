/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./index.html', './src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			screens: {
				mb: '400px'
			}
		}
	},
	plugins: [require('tailwind-scrollbar')]
}

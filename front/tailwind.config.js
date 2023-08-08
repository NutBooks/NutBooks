/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{js,ts,jsx,tsx,mdx}'],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Pretendard'],
        PreB: ['Pretendard-Bold'],
        PreR: ['Pretendard-Regular'],
        PreEB: ['Pretendard-ExtraBold'],
      },
      colors: {
        grassgreen: '#75A86C',
        black: '#252525',
        lightgrey: '#FAFAFA',
      },
    },
  },
  plugins: [],
};

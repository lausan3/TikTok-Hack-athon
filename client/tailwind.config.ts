import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
      fontFamily: {
        sans: ['Inter', 'sans-serif'],
      },
      colors: {
        "brand-primary": '#fc3d3d',
        "brand-secondary": '#3d83fc',
        "bg-primary": '#F5F5F5',
        "bg-secondary": '#E5E5E5',
      },
      dropShadow: {
        'logo-red': '2px 2px 0px rgba(252, 61, 61, 1)',
        'logo-blue': '2px 2px 0px rgba(61, 131, 252, 1)',
      }
    },
    },
  plugins: [],
};
export default config;

import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: {
    template: "%s | TikTok Hack",
    default: "TikTok Hack",
  },
  description: "Anthony Lau and Marouane Hachemi's submission for the TikTok shop track for the second Headstarter Hackathon.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${inter.className}`}>
        <div className="text-white min-h-dvh flex flex-col bg-gradient-radial from-gray-800 to-zinc-900">
          {children}
        </div>
      </body>
    </html>
  );
}

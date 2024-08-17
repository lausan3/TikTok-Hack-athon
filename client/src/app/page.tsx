import Footer from "@/components/footer";
import Posts from "@/components/home/posts";
import Navbar from "@/components/navbar";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: 'Home | TikTok Hack'
}

export default function Home() {
  
  return (
    <div className="h-dvh w-dvw flex flex-col bg-bg-primary p-4 gap-y-4">
      <Navbar />

      <Posts />

      <Footer />
    </div>
  );
}

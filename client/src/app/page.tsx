import Footer from "@/ui/footer";
import Posts from "@/ui/home/posts";
import Navbar from "@/ui/navbar";
import Sidebar from "@/ui/sidebar";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: 'Home | TikTok Hack'
}

export default function Home() {
  
  return (
    <>
      <Navbar />

      <div className="flex flex-row flex-grow">
        <Sidebar />
        <Posts />
      </div>

      <Footer />
    </>
  );
}

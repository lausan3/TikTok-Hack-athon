import Footer from "@/ui/footer";
import Posts from "@/ui/home/posts";
import Navbar from "@/ui/navbar";
import Sidebar from "@/ui/sidebar/sidebar";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: 'Home | TikTok Hack'
}

export default function Home() {
  
  return (
    <>
      <Navbar />

      <div className="content-container">
        <Sidebar />

        <div className="inner-content-container">
          <Posts />
        </div>

      </div>

      <Footer />
    </>
  );
}

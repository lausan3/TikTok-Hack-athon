import Footer from "@/ui/footer";
import Posts from "@/ui/home/posts";
import Navbar from "@/ui/navbar";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: 'Home | TikTok Hack'
}

export default function Home() {
  
  return (
    <>
      <Navbar />

      <Posts />

      <Footer />
    </>
  );
}

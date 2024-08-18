"use client"

import Footer from "@/ui/footer";
import Navbar from "@/ui/navbar";
import CreatePostForm from "@/ui/posts/create-post/create-post-form";
import Sidebar from "@/ui/sidebar";
import { redirect } from "next/navigation";

export default function CreatePost() {
  const data = localStorage.getItem('tiktok_user_login');
  
  if (!data) {
    redirect('/auth/login');
  }

  return (
    <>
      <Navbar />

      <div className="content-container">
        <Sidebar />

        <div className="inner-content-container">
          <CreatePostForm />
        </div>

      </div>

      <Footer />
        
    </>
  );
}

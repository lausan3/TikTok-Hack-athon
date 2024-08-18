import Footer from "@/ui/footer";
import Navbar from "@/ui/navbar";
import CreatePostForm from "@/ui/posts/create-post/create-post-form";
import Sidebar from "@/ui/sidebar";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: 'Post'
}

export default function CreatePost() {
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

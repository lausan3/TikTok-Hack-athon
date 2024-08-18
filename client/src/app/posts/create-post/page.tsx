import Footer from "@/ui/footer";
import Navbar from "@/ui/navbar";
import CreatePostForm from "@/ui/posts/create-post/create-post-form";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: 'Post'
}

export default function CreatePost() {
  return (
    <>
      <Navbar />

        <CreatePostForm />
        
        <Footer />
    </>
  );
}

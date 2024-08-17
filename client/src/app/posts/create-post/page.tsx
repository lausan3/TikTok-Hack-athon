import Footer from "@/components/footer";
import Navbar from "@/components/navbar";
import CreatePostForm from "@/components/posts/create-post/create-post-form";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: 'Post'
}

export default function CreatePost() {
  return (
    <div className="h-dvh w-dvw flex flex-col bg-bg-primary p-4 gap-y-4">
      <Navbar />

      <CreatePostForm />

      <Footer />
    </div>
  );
}

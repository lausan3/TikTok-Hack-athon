"use client"

import { createPost } from "@/lib/create-post";
import { useRouter } from "next/navigation";
import { FormEvent, useState } from "react";

export default function CreatePostForm() {
  const [formData, setFormData] = useState({
    title: "",
    description: ""
  });
  const [error, setError] = useState<string | null>(null);

  const router = useRouter();

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    
    try {
      const response = await createPost(formData.title, formData.description);
     
      if (response.status === 403) {
        console.log("redirect");
        router.push('/auth/login');
        return;
      }

      const data = await response.json();

      if (response.ok) {
        router.push('/');
      } else {
        setError(data.error);
      }
    } catch (error) {
      console.error(error);
    }
  }

  return (
    <form 
      className="form-container w-1/3 sm:w-1/2"
      onSubmit={(e) => handleSubmit(e)}
    >
      <h1 className="w-1/2 text-xl font-bold text-center mb-4">Share your thoughts with the world, we won&apos;t judge.</h1>

      <p className={`${error ? 'text-red-500' : 'hidden'} text-sm`}>{error}</p>

      <div className="flex flex-col my-4 gap-y-4 w-1/2">

        <div className="flex flex-col space-y-1">
          <label htmlFor="title">Title</label>
          <input 
            name="title" 
            type="text" 
            required
            placeholder="What are you thinking about?"
            value={formData.title}
            onChange={(e) => setFormData({ ...formData, title: e.target.value })}
          />
        </div>

        <div className="flex flex-col space-y-1">

          <label htmlFor="description">Description</label>
          <textarea 
            name="description" 
            rows={10}
            required
            placeholder="Show the world something cool!"
            value={formData.description}
            onChange={(e) => setFormData({ ...formData, description: e.target.value })}
          />
        </div>
      </div>

      <button 
      className="bg-blue-800 hover:bg-blue-700 text-white font-bold py-2 px-4"
        type="submit"
      >
        Submit
      </button>
    </form>
  )
}
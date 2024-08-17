"use client"

import { FormEvent, useState } from "react";

export default function CreatePostForm() {
  const [formData, setFormData] = useState({
    title: "",
    description: ""
  });

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log(formData);
  }

  return (
    <form 
      id="create-post-form" 
      className="flex flex-col gap-4"
      onSubmit={(e) => handleSubmit(e)}
    >
      <input 
        id="title" 
        className="p-2 border border-black rounded-lg" 
        type="text" 
        placeholder="Title"
        value={formData.title}
        onChange={(e) => setFormData({ ...formData, title: e.target.value })}
      />
      <input 
        id="description" 
        className="p-2 border border-black rounded-lg" 
        type="text" 
        placeholder="Description"
        value={formData.description}
        onChange={(e) => setFormData({ ...formData, description: e.target.value })}
      />
      <button 
        type="submit"
        className="p-2 bg-green border border-black rounded-lg 
              hover:bg-red-400"
      >
        Submit
      </button>
    </form>
  )
}
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
      className="form-container  sm:1/2 lg:1/4"
      onSubmit={(e) => handleSubmit(e)}
    >
      <h1 className="w-1/2 text-xl font-bold text-center mb-4">Share your thoughts with the world, we won't judge.</h1>

      <div className="flex flex-col my-4 gap-y-4">

        <div className="flex flex-col space-y-1">
          <label htmlFor="title">Title</label>
          <input 
            name="title" 
            type="text" 
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
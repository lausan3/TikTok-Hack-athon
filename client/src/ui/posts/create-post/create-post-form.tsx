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
      className="flex flex-col flex-grow p-8"
      onSubmit={(e) => handleSubmit(e)}
    >
      <div className="flex flex-col my-4 gap-y-2">

        <label htmlFor="title">Title</label>
        <input 
          id="title" 
          type="text" 
          placeholder="What are you thinking about?"
          value={formData.title}
          onChange={(e) => setFormData({ ...formData, title: e.target.value })}
        />
        <label htmlFor="description">Description</label>
        <textarea 
          id="description" 
          rows={10}
          placeholder="Show the world something cool!"
          value={formData.description}
          onChange={(e) => setFormData({ ...formData, description: e.target.value })}
        />
      </div>

      <button 
        type="submit"
      >
        Submit
      </button>
    </form>
  )
}
"use client"

import { useState } from "react";

export default function Posts() {
  const [posts, setPosts] = useState([{
    user: "Alice",
    title: "Hello World",
    content: "Welcome to my blog!",
  },
  {
    user: "Bob",
    title: "Hello Again",
    content: "Another post here!",
  }]);
  
  return (
    <main id="posts" className="flex flex-col gap-4">
      {posts.map((post, index) => (
        <div key={index} className="flex flex-col gap-2 p-4 border border-black rounded-lg">
          <h2 className="text-xl font-bold">{post.title}</h2>
          <p className="text-md">{post.content}</p>
        </div>
      ))}
    </main>
  )
}
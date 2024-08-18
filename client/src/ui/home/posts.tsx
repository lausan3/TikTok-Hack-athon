"use client"

import { fetchAllPosts } from "@/lib/fetch-posts";
import { IPostData } from "@/types/types";
import { useEffect, useState } from "react";
import { PostSkeleton } from "../skeletons";

export default function Posts() {
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getPosts = async () => {
      try {
        const posts = await fetchAllPosts()
        .then(data => data?.posts)
        .catch(error => console.error(error))
          
        setLoading(false);
        
        if (!posts) {
          return console.error("Failed to fetch posts");
        }

        if (!(posts instanceof Array)) {
          return console.error("Posts is not an array");
        }
        
        setPosts(posts as IPostData[]);
      } catch (error) {
        console.error(error);
      }
    }
    
    setTimeout(() => getPosts(), 750)
  }, [])

  const [posts, setPosts] = useState<IPostData[]>([]);

  if (loading) {
    return <PostSkeleton />;
  }
  
  return (
    <main id="posts" className="flex flex-col p-8 gap-4 flex-grow">
      {
        posts.length > 0 ?
          posts.map((post, index) => {
            const date: string = `${new Date(post.created_at).toLocaleDateString()} at ${new Date(post.created_at).toLocaleTimeString()}`;

            return (
              <div key={index} className="flex flex-col h-fit gap-2 p-8 border border-zinc-700 bg-bg-primary rounded-lg">
                <h2 className="text-xl font-bold">{post.title}</h2>

                <div className="flex flex-row w-fit text-start text-sm text-label">
                  <p className="mr-2">Posted by {post.user_name || "Anonymous"} on {date}</p>
                </div>
                  
                <p className="text-label">{post.content}</p>
              </div>
            )})
        :
        <p className="text-lg text-center">No posts available... :(</p>
    }
    </main>
  )
}
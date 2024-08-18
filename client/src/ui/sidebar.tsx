"use client"

import Link from "next/link";
import { useState } from "react";

export default function Sidebar() {
  const [notifications, setNotifications] = useState([]);
  
  return (
    <section id="sidebar" className="flex flex-col w-1/3 h-full py-2 px-4 bg-zinc-900">
      <div className="sidebar-stack mb-6">
        <p className="subheading-text">Navigation</p>
        
        <Link href="/">
          <p className="nav-link">All Posts</p>
        </Link>

        <Link href="/posts/create-post">
          <p className="nav-link">Create Post</p>
        </Link>

      </div>

      <div className="sidebar-stack">
        <p className="subheading-text">Notifications</p>
        
        {
          notifications.length > 0 ? 
            notifications.map(notification => (
              <p className="notification">{notification}</p>
            )) 
            :
              <p className="notification">No notifications D:</p>
        }
      </div>
    </section>
  )
}
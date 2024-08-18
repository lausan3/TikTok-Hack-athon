"use client"

import { connectWebsocket } from "@/lib/ws";
import { IPostData } from "@/types/types";
import Link from "next/link";
import { useEffect, useState } from "react";
import ClearNotifications from "./clear-notifications";
import Notification from "./notification";

let socket: WebSocket | null = null;

export default function Sidebar() {
  const [notifications, setNotifications] = useState<IPostData[]>([]);
  const [token, setToken] = useState("");

  useEffect(() => {
    
    if (!token) {
      const localData = localStorage.getItem("tiktok_user_login");
      
      if (localData) {
        const parsedData = JSON.parse(localData);
        setToken(parsedData.token);
      }
    }

    // Check for notifications every second
    const intervalId = setInterval(() => { 
      // console.log("Checking for notifications...");

      const prevNotifications = localStorage.getItem("notifications");

      if (prevNotifications) {
        setNotifications(JSON.parse(prevNotifications));
      } else {
        setNotifications([]);
      }

      // Check if WebSocket is not open and reconnect if necessary
      if (token && (!socket || socket.readyState !== WebSocket.OPEN)) {
        console.log("WebSocket is not open, reconnecting...");
        socket = connectWebsocket(token);
      }
    }, 1000);

    return () => clearInterval(intervalId);
  }, [token]);
  
  return (
    <section 
      id="sidebar" 
      className="w-1/4 h-auto p-4 bg-zinc-900"
    >
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
        
        <ClearNotifications />
        
        {
          token ? (
            notifications.length > 0 ? 
              notifications.map((notification, index) => 
                <Notification key={index} notification={notification}/>
              ) 
            :
              <p className="notification">No notifications D:</p>
            )
          :
          <p className="notification">You must be logged in to get notifications!</p>
        }
      </div>
    </section>
  )
}
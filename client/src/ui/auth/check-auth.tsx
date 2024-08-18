"use client"

import { redirect } from "next/navigation";
import { useEffect } from "react";

export default function CheckAuth() {

  useEffect(() => {
    const loggedIn = localStorage.getItem('tiktok_user_login');

    if (!loggedIn) {
      redirect('/auth/login');
    }
  }, [])
  

  return (
    <></>
  )
}
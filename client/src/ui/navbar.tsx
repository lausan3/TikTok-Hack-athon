"use client"

import Logo from "./logo";
import { LoginButton, SignOutButton } from "./auth/auth-buttons";
import { useEffect, useState } from "react";

export default function Navbar() {
  const [loggedIn, setLoggedIn] = useState(false);

  useEffect(() => {
    const user = localStorage.getItem("tiktok_user_login");

    if (user) {
      setLoggedIn(true);
    } else {
      setLoggedIn(false);
    }
  }, []);

  return (
    <nav id="navbar" className="flex flex-row items-center bg-header-footer px-6 py-4 shadow-md shadow-zinc-900">

      <Logo />
      
      <div id="spacer" className="flex-grow"></div>

      {
        loggedIn ?
          <SignOutButton />
        :
          <LoginButton />
      }
    </nav>
  )
}
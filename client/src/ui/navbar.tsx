"use client"
import Logo from "./logo";
import { LoginButton, SignOutButton } from "./auth/auth-buttons";

export default function Navbar() {
  const user = localStorage.getItem('tiktok_user_login') ? true : false;

  return (
    <nav id="navbar" className="flex flex-row items-center bg-header-footer px-6 py-4 shadow-md shadow-zinc-900">

      <Logo />
      
      <div id="spacer" className="flex-grow"></div>

      {
        user ?
          <SignOutButton />
        :
          <LoginButton />
      }
    </nav>
  )
}
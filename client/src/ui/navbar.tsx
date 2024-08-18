import Link from "next/link";
import Logo from "./logo";

export default function Navbar() {
  const user = false;

  return (
    <nav id="navbar" className="flex flex-row items-center bg-header-footer px-6 py-4 shadow-md shadow-zinc-900">

      <Logo />
      
      <div id="spacer" className="flex-grow"></div>

      <Link
        href="/auth/login"
        className="nav-link"
      >
        {user ? "Sign Out" : "Sign In to Post"}
      </Link>
    </nav>
  )
}
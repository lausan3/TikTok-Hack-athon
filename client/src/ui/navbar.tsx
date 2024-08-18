import Link from "next/link";
import Logo from "./logo";

export default function Navbar() {
  return (
    <nav id="navbar" className="flex flex-row items-center bg-header-footer px-6 py-4 shadow-md shadow-zinc-800">

      <Logo />
      
      <div id="spacer" className="flex-grow"></div>

      <Link
        href="/posts/create-post"
        className="nav-link"
      >
        Create Post
      </Link>
    </nav>
  )
}
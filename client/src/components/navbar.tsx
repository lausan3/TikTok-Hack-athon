import Link from "next/link";
import Logo from "./logo";

export default function Navbar() {
  return (
    <nav id="navbar" className="flex flex-row items-center">

      <Logo />

      <Link
        href="/posts/create-post"
        className="flex justify-self-end items-center p-2 h-fit border border-black rounded-lg
        hover:bg-red-400"
      >
        Create Post
      </Link>
    </nav>
  )
}
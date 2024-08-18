import { signOutUser } from "@/lib/auth";
import Link from "next/link";

export function LoginButton() {
  return (
    <Link
      href="/auth/login"
      className="nav-link"
    >
      Sign in to post
    </Link>
  )
}

export function SignOutButton() {
  return (
    <Link
      href="/auth/sign-up"
      className="nav-link"
      onClick={signOutUser}
    >
      Sign Out
    </Link>
  )
}
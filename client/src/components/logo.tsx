import Link from "next/link";

export default function Logo() {
  return (
    <Link 
      id="logo" 
      className="flex flex-row p-4 text-center" 
      href="/"
    >
      <h1 className="text-4xl font-semibold drop-shadow-logo-red">Tik</h1>
      <h1 className="text-4xl font-semibold drop-shadow-logo-blue">Tok</h1>
    </Link>
  )
}
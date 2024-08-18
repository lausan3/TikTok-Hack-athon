import SignUpForm from "@/ui/auth/sign-up-form";
import Footer from "@/ui/footer";
import Navbar from "@/ui/navbar";

export default function LoginPage() {
  return (
    <>
      <Navbar />

      <div className="flex flex-col h-full items-center justify-center flex-grow">
        <SignUpForm />
      </div>

      <Footer />
    </>
  )
}
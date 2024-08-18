import LogInForm from "@/ui/auth/log-in-form";
import Footer from "@/ui/footer";
import Navbar from "@/ui/navbar";

export default function LoginPage() {
  return (
    <>
      <Navbar />

      <div className="flex flex-col h-full items-center justify-center flex-grow">
        <LogInForm />
      </div>

      <Footer />
    </>
  )
}
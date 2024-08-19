"use client"

import { signUpUser } from '@/lib/auth';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { FormEvent, useState } from 'react'

export default function LogInForm() {
  const [formData, setFormData] = useState<{ username: string, password: string, confirmPassword: string }>({ username: '', password: '', confirmPassword: '' });
  const [error, setError] = useState<string | null>(null);

  const router = useRouter();

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (formData.username === "" || formData.password === "") {
      setError('There are missing entries, please fill them out.');
      return;
    }

    if (formData.password !== formData.confirmPassword) {
      setError('Passwords do not match, please try again.');
      return;
    }

    try {

      const data = await signUpUser(formData.username, formData.password);

      if (data.token) {
        setError(null); 
        
        router.push('/');
      } else {
        setError(data.error); 
      }
    } catch (error) {
      console.error(error);
    }
  }

  return (
    <>
      <form 
        className='form-container w-1/2 lg:w-1/3'
        onSubmit={(e) => handleSubmit(e)}
      >
        <h1 className="w-1/2 text-xl font-bold text-center">Join us... Or else....</h1>
        
        <p className={`${error ? 'text-red-500' : 'hidden'}`}>{error}</p>

        <div className="flex flex-col space-y-4">

          <div 
            className="flex flex-col space-y-1"
          >

            <label
              htmlFor="username"
              >
              Username
            </label>
            <input
              name="username"
              type="text"
              required
              placeholder="super cool name..."
              value={formData.username}
              onChange={(e) => setFormData({ ...formData, username: e.target.value })}
            />

          </div>

          <div className="flex flex-col space-y-1">
            <label
              htmlFor="password"
              >
              Password
            </label>
            <input
              name="password"
              type="password"
              required
              placeholder="super strong password..."
              value={formData.password}
              onChange={(e) => setFormData({ ...formData, password: e.target.value })}
            />
          </div>

          <div className="flex flex-col space-y-1">
            <label
              htmlFor="confirm-password"
              >
              Confirm Password
            </label>
            <input
              name="confirm-password"
              type="password"
              required
              placeholder="make sure it's the same..."
              value={formData.confirmPassword}
              onChange={(e) => setFormData({ ...formData, confirmPassword: e.target.value })}
            />
          </div>

          <p className="text-sm text-label">
            Already have an account?
            <Link className="normal-link text-sm" href="/auth/login"> Log In!</Link>
          </p>
        </div>
        <button 
          className="bg-blue-800 hover:bg-blue-700 text-white font-bold py-2 px-4"
          type="submit"
        >Sign Up</button>
      </form>
    </>
  )
}
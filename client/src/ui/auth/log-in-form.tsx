"use client"

import { loginUser } from '@/lib/auth';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { FormEvent, useState } from 'react'

export default function LogInForm() {
  const [formData, setFormData] = useState<{ username: string, password: string}>({ username: '', password: '' });
  const [error, setError] = useState<string | null>(null);

  const router = useRouter();

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (formData.username === "" || formData.password === "") {
      setError('There are missing entries, please fill them out.');
      return;
    }
    try {

      const response = await loginUser(formData.username, formData.password);

      const data = await response.json()
      if (response.ok) {
        setError(null); 

        localStorage.setItem('tiktok_user_login', JSON.stringify({ username: formData.username, token: data.token, expires: data.expires }));

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
        <h1 className="w-1/2 text-xl font-bold text-center">Welcome back, friend!</h1>
        
        <p className={`${error ? 'text-red-500' : 'hidden'} text-sm`}>{error}</p>

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

          <p className="text-sm text-label">
            Don't have an account?
            <Link className="normal-link text-sm" href="/auth/sign-up"> Sign up!</Link>
          </p>
        </div>
        <button 
          className="bg-blue-800 hover:bg-blue-700 text-white font-bold py-2 px-4"
          type="submit"
        >Log In</button>
      </form>
    </>
  )
}
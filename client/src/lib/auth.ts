export const signUpUser = async (user_name: string, password: string) => {
  try {
    const URL = process.env.NEXT_PUBLIC_SERVER_URL;

    const response = await fetch(`${URL}/api/auth/signup`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ user_name, password }),
    });

    const data = await response.json();

    if (response.ok) {
      localStorage.setItem('tiktok_user_login', JSON.stringify({ username: user_name, token: data.token, expires: data.expires }));
    }

    return data;
  } catch (error) {
    console.error(error);
    return new Response(JSON.stringify({ error: error }), { status: 400 });
  }
}

export const loginUser = async (user_name: string, password: string) => {
  try {
    const URL = process.env.NEXT_PUBLIC_SERVER_URL;

    const response = await fetch(`${URL}/api/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ user_name, password }),
    });

    const data = await response.json();

    if (response.ok) {
      localStorage.setItem('tiktok_user_login', JSON.stringify({ username: user_name, token: data.token, expires: data.expires }));
    }

    return data;
  } catch (error) {
    console.error(error);
    return new Response(JSON.stringify({ error: error }), { status: 400 });
  }
}

export const signOutUser = () => {
  localStorage.removeItem('tiktok_user_login');

  window.location.reload();
}
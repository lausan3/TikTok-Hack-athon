export const signUpUser = async (username: string, password: string) => {
  try {
    const URL = process.env.NEXT_PUBLIC_SERVER_URL;

    const response = await fetch(`${URL}/api/auth/signup`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, password }),
    });

    if (response.ok) {
      const data = await response.json();
      return data;
    }
  } catch (error) {
    console.error(error);
  }
}

export const loginUser = async (username: string, password: string) => {
  try {
    const URL = process.env.NEXT_PUBLIC_SERVER_URL;

    const response = await fetch(`${URL}/api/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, password }),
    });

    if (response.ok) {
      const data = await response.json();
      return data;
    }
  } catch (error) {
    console.error(error);
  }
}
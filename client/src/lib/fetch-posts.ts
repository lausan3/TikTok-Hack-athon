export async function fetchPostsForUsername(username: string) {
  try {
    const URL = process.env.NEXT_PUBLIC_SERVER_URL;

    const response = await fetch(`${URL}/api/posts/${username}`, {
      method: "GET",
    });

    if (response.ok) {
      const data = await response.json();
      return data;
    }
  } catch (error) {
    console.error(error);
  }
}

export async function fetchAllPosts() {
  try {
    const URL = process.env.NEXT_PUBLIC_SERVER_URL;

    const response = await fetch(`${URL}/api/posts`, {
      method: "GET",
    });

    if (response.ok) {
      const data = await response.json();
      return data;
    }
  } catch (error) {
    console.error(error);
  }
}
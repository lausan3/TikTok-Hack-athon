export async function createPost(title: string, content: string) {
  const URL = process.env.NEXT_PUBLIC_SERVER_URL;
  
  const localData = localStorage.getItem('tiktok_user_login');
  let userData;

  if (localData) {
    userData = JSON.parse(localData);
  } else {
    return new Response("redirect", { status: 400, statusText: "No user data found" });
  }

  if (!userData.token || (!userData.expires && userData.expires < Date.now())) {
    return new Response(null, { status: 403 });
  }

  return fetch(`${URL}/api/posts`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${userData.token}`
    },
    body: JSON.stringify({ title: title, content: content })
  });
}
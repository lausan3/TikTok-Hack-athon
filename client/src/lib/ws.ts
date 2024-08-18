import { IPostData } from "@/types/types";

export function connectWebsocket(token: string) {  
  const HOST = process.env.NEXT_PUBLIC_HOST || "localhost:8080";
  const ws = new WebSocket(`ws://${HOST}/ws?token=${token}`);
  
  ws.onopen = () => {
    console.log("Connected to websocket");
  }

  ws.onmessage = (message) => {
    let data;
    try {
      data = JSON.parse(message.data);
    } catch (error) {
      console.error("Error parsing websocket message:", error);
      return;
    }

    console.log("New notification:", data);
    const post = data.content as IPostData;

    const storedNotifications = localStorage.getItem("notifications");
  
    if (storedNotifications) {
      const prevNotifications = JSON.parse(storedNotifications);
      const newNotifications = [...prevNotifications, post];
      localStorage.setItem("notifications", JSON.stringify(newNotifications));
    } else {
      localStorage.setItem("notifications", JSON.stringify([post]));
    }
  }

  ws.onerror = (error) => {
    console.error("Error connecting to websocket:", error);
  }

  ws.onclose = () => {
    console.log("Websocket connection closed");
    localStorage.removeItem("websocket-connected");
  }
  
  return ws
}
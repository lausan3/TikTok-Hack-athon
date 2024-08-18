import { IPostData } from "@/types/types";

export default function Notification({ notification, index }: {
  notification: IPostData,
  index: number,
}) {
  const time = new Date(notification.created_at + " UTC").toLocaleTimeString();

  return (
    <div key={index} className="text-autosize text-label flex flex-col p-4 bg-bg-primary rounded-xl">
      <h3 className="text-white">{notification.title}</h3>
      <p>{notification.content}</p>
      <small>{`From ${notification.user_name} at ${time}`}</small>
    </div>
  )
}
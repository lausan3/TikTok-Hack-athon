export default function ClearNotifications() {
  const handleClearNotifications = () => {
    localStorage.removeItem("notifications");
  }

  return (
    <button className="flex items-center justify-center cta-btn text-left">
      <p className="text-xs font-normal sm:text-lg">Clear Notifications</p>
    </button>
  )
}
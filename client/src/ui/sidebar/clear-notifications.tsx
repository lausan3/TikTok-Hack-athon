"use client"

export default function ClearNotifications() {
  const handleClearNotifications = () => {
    console.log("clearing notifications");

    localStorage.removeItem("notifications");
  }

  return (
    <button 
      className="flex items-center justify-center cta-btn text-left"
      onClick={handleClearNotifications}
    >
      <p className="text-xs font-normal sm:text-lg">Clear Notifications</p>
    </button>
  )
}
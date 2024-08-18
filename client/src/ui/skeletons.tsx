export const PostSkeleton = () => {
  return (
    <main className="flex flex-col p-8 gap-4">
      <div className="animate-pulse bg-bg-secondary w-150 h-8 mb-4"></div>
      <div className="flex flex-col h-fit gap-2 p-8 border border-zinc-700 bg-bg-primary rounded-lg">
        <div className="animate-pulse bg-bg-secondary w-40 h-8"></div>
        <div className="animate-pulse bg-bg-secondary w-96 h-5"></div>
        <div className="animate-pulse bg-bg-secondary w-24 h-5"></div>
      </div>
    </main>
  )
}
export const PostSkeleton = () => {
  return (
    <main className="flex flex-col flex-grow p-8">
      <div className={`flex flex-col gap-2 p-8 border border-zinc-700 bg-bg-primary rounded-lg`}>
        <div className="animate-pulse bg-bg-secondary w-40 h-8"></div>
        <div className="animate-pulse bg-bg-secondary w-96 h-5"></div>
        <div className="animate-pulse bg-bg-secondary w-24 h-5"></div>
      </div>
    </main>
  )
}
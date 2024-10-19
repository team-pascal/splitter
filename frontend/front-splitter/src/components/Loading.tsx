export function Loading() {
  return (
    <div className='absolute left-1/2 top-1/2 flex -translate-x-1/2 -translate-y-1/2 flex-col items-center p-40'>
      <div className='size-14 animate-spin rounded-full border-4 border-zinc-400 border-t-transparent'></div>
      <h1 className='pt-6 text-lg font-semibold'>Loading...</h1>
    </div>
  );
}

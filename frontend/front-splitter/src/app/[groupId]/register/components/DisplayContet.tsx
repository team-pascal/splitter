type Props = {
  name: string;
  children: React.ReactNode;
};

export function DisplayContent({ name, children }: Props) {
  return (
    <div className='flex flex-row items-center border-b-2 border-gray-300'>
      <p className=' grow text-lg text-gray-600'>{name}</p>
      {children}
    </div>
  );
}

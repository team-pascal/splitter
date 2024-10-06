type Props = {
  member: string;
  children: React.ReactNode;
};

export function Member({ member, children }: Props) {
  return (
    <div className='mb-6 flex flex-row border-b-2 border-gray-300'>
      <p className='flex grow items-center text-lg'>{member}</p>
      {children}
    </div>
  );
}

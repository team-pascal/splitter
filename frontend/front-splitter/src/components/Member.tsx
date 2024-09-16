type Props = {
  member: string;
  children: React.ReactNode;
};

export function Member({ member, children }: Props) {
  return (
    <div className="flex flex-row border-b-2 border-gray-300 mb-6">
      <p className="grow text-lg flex items-center">{member}</p>
      {children}
    </div>
  );
}

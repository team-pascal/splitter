type Props = {
  name: string;
  children: React.ReactNode;
};

export function DisplayContent({ name, children }: Props) {
  return (
    <div className="flex flex-row border-b-2 border-gray-300 items-center">
      <p className=" text-gray-600 text-lg grow">{name}</p>
      {children}
    </div>
  );
}

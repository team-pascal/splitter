import Link from 'next/link';

type Props = {
  contentText: string;
  register: () => void;
  path: string;
};

export function RegistrationBtn({ contentText, register, path }: Props) {
  return (
    <Link
      href={path}
      className="bg-transparent hover:bg-gray-300 rounded-3xl shadow-xl py-2 px-8 font-bold border-4 border-gray-400"
    >
      {contentText}
    </Link>
  );
}

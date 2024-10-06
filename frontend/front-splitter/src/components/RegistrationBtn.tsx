import Link from 'next/link';

type Props = {
  contentText: string;
  register: () => void;
  path: string;
};

export function RegistrationBtn({ contentText, path }: Props) {
  return (
    <Link
      href={path}
      className='rounded-3xl border-4 border-gray-400 bg-transparent px-8 py-2 font-bold shadow-xl hover:bg-gray-300'
    >
      {contentText}
    </Link>
  );
}

'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import HomeIcon from '@mui/icons-material/Home';
import HistoryIcon from '@mui/icons-material/History';
import AddCardIcon from '@mui/icons-material/AddCard';

const styles = {
  div: 'basis-1/3 text-center bg-custom-btn hover:bg-custom-btnhover border border-custom-btnborder p-2 mx-2 md:mx-10 rounded-full shadow-sm',
  p: 'text-xs md:text-sm pt-0.5',
};

export function Menu() {
  const currentPath = usePathname();
  const homePattern = new RegExp('^/[^/]*$');
  const registPattern = new RegExp('^/[^/]+/register$');
  const personalPattern = new RegExp('^/[^/]+/personal-history$');

  return (
    <ul className="inline-flex justify-center bg-custom-bg rounded-full border border-custom-border p-2 md:p-4 shadow-lg">
      <li>
        <Link
          href="/test"
          className={`font-medium text-xl ${
            homePattern.test(currentPath)
              ? 'text-white font-black'
              : 'text-gray-400'
          }`}
        >
          <div className={styles.div}>
            <HomeIcon
              sx={{
                fontSize: { md: 35 },
                color: `${homePattern.test(currentPath) ? 'white' : 'gray'}`,
              }}
            />
          </div>
        </Link>
      </li>
      <li>
        <Link
          href="/test/register"
          className={`font-medium text-xl ${
            registPattern.test(currentPath)
              ? 'text-white font-black'
              : 'text-gray-400'
          }`}
        >
          <div className={styles.div}>
            <AddCardIcon
              sx={{
                fontSize: { md: 35 },
                color: `${registPattern.test(currentPath) ? 'white' : 'gray'}`,
              }}
            />
          </div>
        </Link>
      </li>
      <li>
        <Link
          href="/test"
          className={`font-medium text-xl cursor-pointer ${
            personalPattern.test(currentPath)
              ? 'text-white font-black'
              : 'text-gray-400'
          }`}
        >
          <div className={styles.div}>
            <HistoryIcon
              sx={{
                fontSize: { md: 35 },
                color: `${
                  personalPattern.test(currentPath) ? 'white' : 'gray'
                }`,
              }}
            />
          </div>
        </Link>
      </li>
    </ul>
  );
}

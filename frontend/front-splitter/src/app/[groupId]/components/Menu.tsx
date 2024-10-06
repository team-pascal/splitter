'use client';

import AddCardIcon from '@mui/icons-material/AddCard';
import HistoryIcon from '@mui/icons-material/History';
import HomeIcon from '@mui/icons-material/Home';
import Link from 'next/link';
import { usePathname } from 'next/navigation';

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
    <ul className='inline-flex justify-center rounded-full border border-custom-border bg-custom-bg p-2 shadow-lg md:p-4'>
      <li>
        <Link
          href='/test'
          className={`text-xl font-medium ${
            homePattern.test(currentPath)
              ? 'font-black text-white'
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
          href='/test/register'
          className={`text-xl font-medium ${
            registPattern.test(currentPath)
              ? 'font-black text-white'
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
          href='/test'
          className={`cursor-pointer text-xl font-medium ${
            personalPattern.test(currentPath)
              ? 'font-black text-white'
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

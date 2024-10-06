import { ReactNode } from 'react';

import { Menu } from './components/Menu';

export default function Layout({ children }: { children: ReactNode }) {
  return (
    <div className='pb-5'>
      <div className='flex justify-center py-5'>
        <Menu />
      </div>
      {children}
    </div>
  );
}

import { Inter } from 'next/font/google';
import '../styles/globals.css';
import { ToastContainer } from 'react-toastify';

import { Header } from '@/components/Header';

import type { Metadata } from 'next';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'Splitter',
  description: 'Created by team',
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang='ja'>
      <body className={inter.className}>
        <header>
          <Header />
        </header>
        <ToastContainer
          position='top-right'
          autoClose={4000}
          hideProgressBar={true}
          pauseOnHover
          draggable
          pauseOnFocusLoss
        />
        <main className='container mx-auto px-5 md:px-14'>{children}</main>
      </body>
    </html>
  );
}

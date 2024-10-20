import Link from 'next/link';

import { HistoryItem } from './HistoryItem';
import { getPayment } from '@/api/payment';
import { HistoryResponse } from '@/types/history.type';

interface Props {
  groupId: string;
}

export async function HistoryList({ groupId }: Props) {
  const historyItems: HistoryResponse = await getPayment(groupId);

  return (
    <div className='min-w-full'>
      <div className='mb-2 flex justify-between text-gray-300'>
        <p className='text-sm md:text-base'>履歴</p>
        <a className='text-sm md:text-base'>もっと見る &gt;</a>
      </div>
      {historyItems.length === 0 ? (
        <p className='mt-40 text-center text-white'>履歴はありません</p>
      ) : (
        <>
          {historyItems.map((data, index) => (
            <Link key={index} href={`${groupId}/${data.genre}/${data.id}`}>
              <HistoryItem
                key={index}
                status={data.done}
                title={data.title}
                amount={data.amount}
                createdAt={data.created_at}
              />
            </Link>
          ))}
        </>
      )}
    </div>
  );
}

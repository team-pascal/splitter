import axios from 'axios';

import { History } from '@/types/history';

import { HistoryItem } from './HistoryItem';

interface Props {
  groupId: string;
}

async function fetchHistoryItems(groupId: string): Promise<History[]> {
  try {
    const response = await axios.get<History[]>(
      'http://localhost:3000/api/payments',
      {
        params: { groupId },
      },
    );
    return response.data;
  } catch (e) {
    console.error('Failed to fetch history data: ', e);
    return [];
  }
}

export async function HistoryList({ groupId }: Props) {
  const historyItems: History[] = await fetchHistoryItems(groupId);

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
            <HistoryItem
              key={index}
              status={data.done}
              title={data.title}
              amount={data.amount}
              createdAt={data.created_at}
            />
          ))}
        </>
      )}
    </div>
  );
}

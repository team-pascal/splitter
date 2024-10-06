import { HistoryItem } from './HistoryItem';

const sampleData = [
  {
    id: 1,
    status: false,
    title: 'ガソリン代',
    fare: '4000',
    createdAt: '2024/5/2 13:32',
  },
  {
    id: 2,
    status: false,
    title: 'おやつ代',
    fare: '4000',
    createdAt: '2024/5/2 13:32',
  },
  {
    id: 3,
    status: true,
    title: 'お昼代',
    fare: '4000',
    createdAt: '2024/5/2 13:32',
  },
  {
    id: 4,
    status: false,
    title: 'ガソリン代',
    fare: '4000',
    createdAt: '2024/5/2 13:32',
  },
  {
    id: 5,
    status: false,
    title: 'おやつ代',
    fare: '4000',
    createdAt: '2024/5/2 13:32',
  },
  {
    id: 6,
    status: true,
    title: 'お昼代',
    fare: '4000',
    createdAt: '2024/5/2 13:32',
  },
  {
    id: 7,
    status: false,
    title: 'ガソリン代',
    fare: '4000',
    createdAt: '2024/5/2 13:32',
  },
  {
    id: 8,
    status: false,
    title: 'おやつ代',
    fare: '4000',
    createdAt: '2024/5/2 13:32',
  },
  {
    id: 9,
    status: true,
    title: 'お昼代',
    fare: '4000',
    createdAt: '2024/5/2 13:32',
  },
];

export function HistoryList() {
  const styles = {
    p: 'text-sm md:text-base',
  };

  return (
    <div className='min-w-full'>
      <div className='mb-2 flex justify-between text-gray-300'>
        <p className={styles.p}>履歴</p>
        <a className={styles.p}>もっと見る &gt;</a>
      </div>
      {sampleData.map((data, index) => (
        <HistoryItem
          key={index}
          id={data.id}
          status={data.status}
          title={data.title}
          fare={data.fare}
          createdAt={data.createdAt}
        />
      ))}
    </div>
  );
}

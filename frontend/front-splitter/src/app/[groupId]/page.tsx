import { HistoryList } from '@/components/groupHome/HistoryList';

interface Params {
  params: { groupId: string };
}

export default function Home({ params }: Params) {
  const { groupId } = params;
  console.log(groupId);
  return (
    <>
      <HistoryList groupId={groupId} />
      <h1>test</h1>
    </>
  );
}

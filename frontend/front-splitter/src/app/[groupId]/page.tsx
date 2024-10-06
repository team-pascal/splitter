import { HistoryList } from './components/history/HistoryList';

interface Params {
  params: { groupId: string };
}

export default function Home({ params }: Params) {
  const { groupId } = params;
  return (
    <>
      <HistoryList groupId={groupId} />
    </>
  );
}

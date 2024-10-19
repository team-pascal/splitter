import { HistoryList } from '@/components/groupHome/HistoryList';

interface Params {
  params: { groupId: string };
}

export default function Home({ params }: Params) {
  const { groupId } = params;
  return <HistoryList groupId={groupId} />;
}

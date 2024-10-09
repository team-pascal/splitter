type Props = {
  id: number;
  status: boolean;
  title: string;
  fare: string;
  createdAt: string;
};
export function HistoryItem({ status, title, fare, createdAt }: Props) {
  const styles = {
    status: 'py-1 px-3 md:px-4 rounded-2xl text-xs md:text-sm text-white',
  };
  const statusStyle = status
    ? styles.status + ' bg-green-600'
    : styles.status + ' bg-red-500';

  return (
    <div className='mb-4 flex cursor-pointer rounded-2xl border-2 border-custom-btnborder bg-custom-btn px-3 py-2 hover:bg-custom-btnhover md:px-6 md:py-4'>
      <div className='mr-3 content-center md:mr-10'>
        <p className={statusStyle}>{status ? '対応済' : '未対応'}</p>
      </div>
      <div className='grow '>
        <p className='mb-1 text-sm md:text-xl'>{title}</p>
        <p className='text-xs text-gray-300 md:text-base'>{createdAt}</p>
      </div>
      <div className='ml-5 content-center'>
        <p className='text-base font-semibold text-white md:text-2xl'>
          {fare}円
        </p>
      </div>
    </div>
  );
}

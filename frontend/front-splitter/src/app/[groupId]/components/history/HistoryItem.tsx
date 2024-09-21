type Props = {
  id: number;
  status: boolean;
  title: string;
  fare: string;
  createdAt: string;
};
export function HistoryItem({ id, status, title, fare, createdAt }: Props) {
  const styles = {
    status: 'py-1 px-3 md:px-4 rounded-2xl text-xs md:text-sm text-white',
  };
  const statusStyle = status
    ? styles.status + ' bg-green-600'
    : styles.status + ' bg-red-500';

  return (
    <div className="flex bg-custom-btn hover:bg-custom-btnhover border-2 border-custom-btnborder rounded-2xl cursor-pointer py-2 px-3 md:py-4 md:px-6 mb-4">
      <div className="content-center mr-3 md:mr-10">
        <p className={statusStyle}>{status ? '対応済' : '未対応'}</p>
      </div>
      <div className="grow ">
        <p className="text-sm md:text-xl mb-1">{title}</p>
        <p className="text-xs md:text-base text-gray-300">{createdAt}</p>
      </div>
      <div className="content-center ml-5">
        <p className="text-base md:text-2xl font-semibold text-white">
          {fare}円
        </p>
      </div>
    </div>
  );
}

import { ErrorSVG } from './ErrorSVG';

export function ErrorComponent() {
  return (
    <div className='absolute left-1/2 top-1/2 flex w-full -translate-x-1/2 -translate-y-1/2 flex-col items-center p-10'>
      <ErrorSVG style={{ width: '100px', height: '100px' }} color='#d1d5db' />
      <h1 className='pt-10 text-lg font-semibold'>エラーが発生しました</h1>
    </div>
  );
}

import { InAdvanceSection } from '@/app/[groupId]/register/components/sections/InAdvanceSection';
import { PaymentSection } from '@/app/[groupId]/register/components/sections/PaymentSection';
import { SplitSection } from '@/app/[groupId]/register/components/sections/SplitSection';
import { Method } from '@/types/method.type';

type Props = {
  paymentMethod: Method;
};

export function PaymentMode({ paymentMethod }: Props) {
  const paymentModeTitle = () => {
    switch (paymentMethod.method) {
      case 'split':
        return '割り勘';
      case 'replacement':
        return '建て替え';
    }
  };

  const paymentMode = () => {
    switch (paymentMethod.method) {
      case 'split':
        return <SplitSection />;
      case 'replacement':
        return <InAdvanceSection />;
    }
  };

  return (
    <div>
      <div className='flex justify-center gap-x-10 py-16'>
        <p className={'cursor-pointer text-lg'}>{paymentModeTitle()}</p>
      </div>
      <PaymentSection />
      {paymentMode()}
    </div>
  );
}

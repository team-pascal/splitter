import { Dispatch, SetStateAction } from 'react';
import { InAdvanceSection } from './InAdvanceSection';
import { PaymentSection } from './PaymentSection';
import { SplitSection } from './SplitSection';
import { Mode } from '@/types';

type Props = {
  paymentMode: Mode;
  setPaymentMode: Dispatch<SetStateAction<Mode>>;
};

export function ModeSection({ paymentMode, setPaymentMode }: Props) {
  return (
    <div>
      <div className="flex justify-center gap-x-10 py-16">
        <p
          className={`text-lg cursor-pointer ${
            paymentMode === 'inAdvance' ? 'opacity-25' : null
          }`}
          onClick={() => setPaymentMode('split')}
        >
          割り勘
        </p>
        <p
          className={`text-lg cursor-pointer ${
            paymentMode === 'split' ? 'opacity-25' : null
          }`}
          onClick={() => setPaymentMode('inAdvance')}
        >
          建て替え
        </p>
      </div>
      <PaymentSection />
      {paymentMode === 'split' ? <SplitSection /> : <InAdvanceSection />}
    </div>
  );
}

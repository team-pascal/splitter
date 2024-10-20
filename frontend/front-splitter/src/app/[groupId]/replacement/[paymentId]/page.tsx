'use client';

import { usePathname } from 'next/navigation';

import { PaymentDetail } from '@/components/paymentDetailes/PaymentDetail';
import { getPaymentId, getPaymentMethod } from '@/utils/getPaymentMethod';

export default function Register() {
  const pathname = usePathname();
  const method = getPaymentMethod(pathname);
  const paymentId = getPaymentId(pathname);

  return <PaymentDetail method={method} paymentId={paymentId} />;
}

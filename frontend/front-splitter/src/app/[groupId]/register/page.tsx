import { sampleMemberList } from '@/app/data/memberList';
import { PaymentInfo } from '@/app/types/type';
import { Index } from './Index';

export default function Register() {
  // fetch member
  const member: Array<PaymentInfo> = sampleMemberList;

  return <Index memberList={member} />;
}

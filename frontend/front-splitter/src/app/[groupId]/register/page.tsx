import { sampleMemberList } from '@/data/memberList';
import { PaymentInfo } from '@/types';

import RegisterForm from '../components/RegisterForm';

export default function Register() {
  // fetch member
  const member: Array<PaymentInfo> = sampleMemberList;

  return <RegisterForm memberList={member} />;
}

import { sampleMemberList } from '@/data/memberList';
import RegisterForm from '../components/RegisterForm';
import { PaymentInfo } from '@/types';

export default function Register() {
  // fetch member
  const member: Array<PaymentInfo> = sampleMemberList;

  return <RegisterForm memberList={member} />;
}

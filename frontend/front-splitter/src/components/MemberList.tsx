import CloseIcon from '@mui/icons-material/Close';

import { Member } from '@/components/Member';

type Props = {
  memberList: string[];
  deleteMember: (deleteMember: string) => void;
};

export function MemberList({ memberList, deleteMember }: Props) {
  const handleDeleteMember = (member: string) => {
    deleteMember(member);
  };
  return (
    <div className='mb-16 mt-10 min-w-full'>
      <p className='mb-5 text-center text-xl'>{memberList.length}人</p>
      {memberList.map((member, index) => (
        <Member key={index} member={member}>
          <CloseIcon
            fontSize='large'
            onClick={() => handleDeleteMember(member)}
            style={{ color: '#929292' }}
          />
        </Member>
      ))}
    </div>
  );
}

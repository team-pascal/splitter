import { Member } from '../../ui/common/Member';
import CloseIcon from '@mui/icons-material/Close';

type Props = {
  memberList: string[];
  deleteMember: (deleteMember: string) => void;
};

export function MemberList({ memberList, deleteMember }: Props) {
  const handleDeleteMember = (member: string) => {
    deleteMember(member);
  };
  return (
    <div className="mt-10 mb-16 min-w-full">
      <p className="mb-5 text-center text-xl">{memberList.length}人</p>
      {memberList.map((member, index) => (
        <Member key={index} member={member}>
          <CloseIcon
            fontSize="large"
            onClick={() => handleDeleteMember(member)}
            style={{ color: '#929292' }}
          />
        </Member>
      ))}
    </div>
  );
}

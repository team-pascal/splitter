import { MemberContext } from '@/app/context/home';
import { PaymentInfo, User } from '@/app/types/type';
import { Dispatch, SetStateAction, useContext } from 'react';

type Props = {
  isOpen: boolean;
  setIsOpen: Dispatch<SetStateAction<boolean>>;
  paymentMember?: Array<PaymentInfo>;
  setPaymentMember?: Dispatch<SetStateAction<PaymentInfo[]>>;
  splitMember?: Array<PaymentInfo>;
  setSplitMember?: Dispatch<SetStateAction<PaymentInfo[]>>;
  inAdvanceMember?: Array<PaymentInfo>;
  setInAdvanceMember?: Dispatch<SetStateAction<PaymentInfo[]>>;
  mode: string;
};

export function Modal({
  isOpen,
  setIsOpen,
  paymentMember,
  setPaymentMember,
  splitMember,
  setSplitMember,
  inAdvanceMember,
  setInAdvanceMember,
  mode,
}: Props) {
  const memberList = useContext(MemberContext);

  function handleClickAddMember(member: User) {
    switch (mode) {
      case 'payment':
        if (paymentMember && setPaymentMember) {
          setPaymentMember([
            ...paymentMember,
            { id: member.id, name: member.name, fee: 0 },
          ]);
        }
        setIsOpen(false);
        return;
      case 'split':
        if (splitMember && setSplitMember) {
          setSplitMember([
            ...splitMember,
            { id: member.id, name: member.name, fee: 0 },
          ]);
        }
        setIsOpen(false);
        return;
      case 'inAdvance':
        if (inAdvanceMember && setInAdvanceMember) {
          setInAdvanceMember([
            ...inAdvanceMember,
            { id: member.id, name: member.name, fee: 0 },
          ]);
        }
        setIsOpen(false);
        return;
      default:
        setIsOpen(false);
        return;
    }
  }

  return (
    isOpen && (
      <div
        aria-hidden="true"
        className="bg-gray-100 top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-full h-full p-5 flex flex-col absolute z-20"
      >
        <div className="relative p-4 w-full h-full">
          <div className="relative  rounded-lg shadow bg-gray-700">
            <div className="flex items-center justify-between p-4 md:p-5 border-b rounded-t border-gray-600">
              <h3 className="text-xl font-semibold text-white">
                メンバーを選択する
              </h3>
              <button
                type="button"
                onClick={() => setIsOpen(false)}
                className="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white"
              >
                <svg
                  className="w-3 h-3"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 14 14"
                >
                  <path
                    stroke="currentColor"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"
                  />
                </svg>
              </button>
            </div>
            <div className="p-4 md:p-5 space-y-4">
              <ul className="text-base leading-relaxed text-gray-500 dark:text-gray-400">
                {memberList && memberList.length !== 0
                  ? memberList.map((member) => (
                      <li key={member.id} className="py-2 cursor-pointer">
                        <a
                          onClick={() => {
                            handleClickAddMember(member);
                          }}
                        >
                          <p className="text-xl">{member.name}</p>
                        </a>
                      </li>
                    ))
                  : null}
              </ul>
            </div>
          </div>
        </div>
      </div>
    )
  );
}

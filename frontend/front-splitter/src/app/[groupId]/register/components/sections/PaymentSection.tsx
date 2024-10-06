import CloseIcon from '@mui/icons-material/Close';
import { useFieldArray, useFormContext } from 'react-hook-form';

import { Select } from '@/app/[groupId]/register/components/Select';
import { Input } from '@/components/Input';
import { InputLabel } from '@/components/InputLabel';

import { FormInput } from '../../types';

export function PaymentSection() {
  const {
    register,
    formState: { errors },
    control,
  } = useFormContext<FormInput>();

  const {
    fields: memberFields,
    append: appendMember,
    remove: removeMember,
  } = useFieldArray({
    control,
    name: 'paymentMembers',
  });

  const {
    fields: costFields,
    append: appendCost,
    remove: removeCost,
  } = useFieldArray({
    control,
    name: 'paymentCost',
  });

  return (
    <>
      <div className='grid grid-cols-2 gap-16'>
        <div>
          <InputLabel label='支払い' />
          {memberFields.map((field, index) => (
            <div key={index} className='pb-2'>
              <Select
                name={`paymentMembers.${index}.name`}
                register={register}
                validation={{
                  required: '支払ったメンバーを選択してね！',
                }}
                error={errors?.paymentMembers?.[index]?.name}
              />
            </div>
          ))}
        </div>
        <div>
          <InputLabel label='金額' />
          {costFields.map((field, index) => (
            <div key={index} className='flex items-center'>
              <div key={index} className='flex-1 pb-2 pr-2'>
                <Input
                  type='number'
                  name={`paymentCost.${index}.cost`}
                  register={register}
                  validation={{ required: '支払った金額を入力してね！' }}
                  error={errors?.paymentCost?.[index]?.cost}
                >
                  <p className='text-base'>円</p>
                </Input>
              </div>
              <div className='mb-4'>
                <CloseIcon
                  fontSize='medium'
                  style={{ color: '#929292' }}
                  onClick={() => {
                    removeMember(index);
                    removeCost(index);
                  }}
                />
              </div>
            </div>
          ))}
        </div>
      </div>
      <button
        type='button'
        className='cursor-pointer text-gray-400'
        onClick={() => {
          appendMember({ name: '' });
          appendCost({ cost: '' });
        }}
      >
        +メンバーを追加する
      </button>
    </>
  );
}

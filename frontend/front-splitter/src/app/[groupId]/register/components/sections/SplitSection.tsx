import CloseIcon from '@mui/icons-material/Close';
import { useFieldArray, useFormContext } from 'react-hook-form';

import { Select } from '@/app/[groupId]/register/components/Select';
import { InputLabel } from '@/components/InputLabel';

import { FormInput } from '../../types';

export function SplitSection() {
  const {
    register,
    formState: { errors },
    control,
  } = useFormContext<FormInput>();

  const { fields, append, remove } = useFieldArray({
    control,
    name: 'splitMembers',
  });

  return (
    <div className='pt-10'>
      <InputLabel label='割り勘メンバー' />
      {fields.map((fields, index) => (
        <div key={index} className='flex items-center'>
          <div key={index} className='flex-1 pb-2 pr-2'>
            <Select
              key={index}
              name={`splitMembers.${index}.name`}
              register={register}
              validation={{ required: '割り勘するメンバーを選択してね！' }}
              error={errors?.splitMembers?.[index]?.name}
            />
          </div>
          <div className='mb-4'>
            <CloseIcon
              fontSize='medium'
              style={{ color: '#929292' }}
              onClick={() => {
                remove(index);
              }}
            />
          </div>
        </div>
      ))}
      <button
        type='button'
        className='cursor-pointer text-gray-400'
        onClick={() => append({ name: '' })}
      >
        +メンバーを追加する
      </button>
    </div>
  );
}

import { useFieldArray, useFormContext } from 'react-hook-form';
import { FormInput } from '@/app/[groupId]/register/Index';
import { Select } from '@/app/ui/register/Select';
import { InputLabel } from '@/app/ui/common/InputLabel';
import CloseIcon from '@mui/icons-material/Close';

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
    <div className="pt-10">
      <InputLabel label="割り勘メンバー" />
      {fields.map((fields, index) => (
        <div key={index} className="flex items-center">
          <div key={index} className="pb-2 flex-1 pr-2">
            <Select
              key={index}
              name={`splitMembers.${index}.name`}
              register={register}
              validation={{ required: '割り勘するメンバーを選択してね！' }}
              error={errors?.splitMembers?.[index]?.name}
            />
          </div>
          <div className="mb-4">
            <CloseIcon
              fontSize="medium"
              style={{ color: '#929292' }}
              onClick={() => {
                remove(index);
              }}
            />
          </div>
        </div>
      ))}
      <button
        type="button"
        className="text-gray-400 cursor-pointer"
        onClick={() => append({ name: '' })}
      >
        +メンバーを追加する
      </button>
    </div>
  );
}

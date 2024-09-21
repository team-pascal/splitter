import { useFieldArray, useFormContext } from 'react-hook-form';
import { Select } from '@/app/[groupId]/register/components/Select';
import CloseIcon from '@mui/icons-material/Close';
import { InputLabel } from '@/components/InputLabel';
import { Input } from '@/components/Input';
import { FormInput } from '../../types';

export function InAdvanceSection() {
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
    name: 'inAdvanceMembers',
  });

  const {
    fields: costFields,
    append: appendCost,
    remove: removeCost,
  } = useFieldArray({
    control,
    name: 'inAdvanceCost',
  });

  return (
    <div className="pt-10">
      <div className="grid grid-cols-2 gap-16">
        <div>
          <InputLabel label="建て替えてもらう" />
          {memberFields.map((fields, index) => (
            <div key={index} className="pb-2">
              <Select
                key={index}
                name={`inAdvanceMembers.${index}.name`}
                register={register}
                validation={{
                  required: '建て替えてもらったメンバーを選択してね！',
                }}
                error={errors?.inAdvanceMembers?.[index]?.name}
              />
            </div>
          ))}
        </div>
        <div>
          <InputLabel label="金額" />
          {costFields.map((field, index) => (
            <div key={index} className="flex items-center">
              <div key={index} className="pb-2 flex-1 pr-2">
                <Input
                  type="number"
                  name={`inAdvanceCost.${index}.cost`}
                  register={register}
                  validation={{ required: '建て替えた金額を入力してね！' }}
                  error={errors?.inAdvanceCost?.[index]?.cost}
                >
                  <p className="text-base leading-0">円</p>
                </Input>
              </div>
              <div className="mb-4">
                <CloseIcon
                  fontSize="medium"
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
        type="button"
        className="text-gray-400 cursor-pointer"
        onClick={() => {
          appendMember({ name: '' });
          appendCost({ cost: '' });
        }}
      >
        +メンバーを追加する
      </button>
    </div>
  );
}

import { InputLabel } from '@/app/ui/common/InputLabel';
import { Select } from '@/app/ui/register/Select';
import { useFieldArray, useFormContext } from 'react-hook-form';
import { FormInput } from '@/app/[groupId]/register/Index';
import { Input } from '@/app/ui/common/Input';
import CloseIcon from '@mui/icons-material/Close';

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
      <div className="grid grid-cols-2 gap-16">
        <div>
          <InputLabel label="支払い" />
          {memberFields.map((field, index) => (
            <div key={index} className="pb-2">
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
          <InputLabel label="金額" />
          {costFields.map((field, index) => (
            <div key={index} className="flex items-center">
              <div key={index} className="pb-2 flex-1 pr-2">
                <Input
                  type="number"
                  name={`paymentCost.${index}.cost`}
                  register={register}
                  validation={{ required: '支払った金額を入力してね！' }}
                  error={errors?.paymentCost?.[index]?.cost}
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
    </>
  );
}

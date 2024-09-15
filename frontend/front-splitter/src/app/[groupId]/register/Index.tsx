'use client';

import { Input } from '../../ui/common/Input';
import { ModeSection } from '../../components/home/register/ModeSection';
import { useState } from 'react';
import { Mode, PaymentInfo } from '@/app/types/type';
import { FormProvider, SubmitHandler, useForm } from 'react-hook-form';
import { InputLabel } from '@/app/ui/common/InputLabel';

export type FormInput = {
  title: string;
  cost: string;
  paymentMembers: { name: string }[];
  paymentCost: { cost: string }[];
  splitMembers: { name: string }[];
  inAdvanceMembers: { name: string }[];
  inAdvanceCost: { cost: string }[];
};

type Props = {
  memberList: Array<PaymentInfo>;
};

export function Index({ memberList }: Props) {
  const [paymentMode, setPaymentMode] = useState<Mode>('split');
  const methods = useForm<FormInput>({
    defaultValues: {
      title: '',
      cost: '',
      paymentMembers: [{ name: '' }],
      paymentCost: [{ cost: '' }],
      splitMembers: [{ name: '' }],
      inAdvanceMembers: [{ name: '' }],
      inAdvanceCost: [{ cost: '' }],
    },
  });

  const onSubmit: SubmitHandler<FormInput> = (data) => {
    console.log(data);
  };

  return (
    <main className="main">
      <FormProvider {...methods}>
        <form onSubmit={methods.handleSubmit(onSubmit)} className="min-w-full">
          <div className="pb-10">
            <InputLabel label="タイトル" />
            <Input
              name="title"
              type="text"
              register={methods.register}
              validation={{
                required: 'タイトルを入力してね！',
                maxLength: {
                  value: 20,
                  message: 'タイトルは20文字以下までだよ！',
                },
              }}
              error={methods.formState.errors.title}
            />
          </div>
          <div>
            <InputLabel label="金額" />
            <Input
              name="cost"
              type="number"
              register={methods.register}
              validation={{
                required: '金額を入力してね！',
              }}
              error={methods.formState.errors.cost}
            >
              <p>円</p>
            </Input>
          </div>
          <ModeSection
            paymentMode={paymentMode}
            setPaymentMode={setPaymentMode}
          />
          <div className="flex justify-end">
            <button className="px-4 py-2 border-2 rounded-md" type="submit">
              保存
            </button>
          </div>
        </form>
      </FormProvider>
    </main>
  );
}

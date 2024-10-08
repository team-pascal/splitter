/* eslint-disable @typescript-eslint/no-misused-promises */
'use client';

import { useState } from 'react';
import { FormProvider, SubmitHandler, useForm } from 'react-hook-form';

import { Input } from '@/components/Input';
import { InputLabel } from '@/components/InputLabel';
import { Mode } from '@/types';

import { ModeSection } from '../register/components/sections/ModeSection';
import { FormInput } from '../types';

export default function RegisterForm() {
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
    <FormProvider {...methods}>
      <form onSubmit={methods.handleSubmit(onSubmit)} className='min-w-full'>
        <div className='pb-10'>
          <InputLabel label='タイトル' />
          <Input
            name='title'
            type='text'
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
          <InputLabel label='金額' />
          <Input
            name='cost'
            type='number'
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
        <div className='flex justify-end'>
          <button className='rounded-md border-2 px-4 py-2' type='submit'>
            保存
          </button>
        </div>
      </form>
    </FormProvider>
  );
}

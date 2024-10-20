'use client';

import { FormProvider, SubmitHandler, useForm } from 'react-hook-form';

import { Input } from '@/components/Input';
import { InputLabel } from '@/components/InputLabel';
import { usePaymentDetail } from '@/hooks/usePaymentDetail';
import { Method } from '@/types/method.type';

import { PaymentMode } from './PaymentMode';
import { ErrorComponent } from '../ErrorComponent';
import { Loading } from '../Loading';
import { FormInput } from '@/types/paymentDetailes/formInput.type';

interface Props {
  method: Method | undefined;
  paymentId: string | undefined;
}

export function PaymentDetail({ method, paymentId }: Props) {
  const [loading, data, error] = usePaymentDetail(paymentId);
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

  if (method === undefined || paymentId === undefined) {
    return <div>ページが存在しません</div>;
  }
  if (error !== null) {
    return <ErrorComponent />;
  }
  if (loading) {
    return <Loading />;
  }

  console.log('error', error);
  console.log('data', data);

  console.log('aafdsfadsfdsjafhdjksafhjkdsahfkdj');

  return (
    <div className='main'>
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
          <PaymentMode paymentMethod={method} />
          <div className='flex justify-end'>
            <button className='rounded-md border-2 px-4 py-2' type='submit'>
              保存
            </button>
          </div>
        </form>
      </FormProvider>
    </div>
  );
}

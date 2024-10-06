'use client';

import CloseIcon from '@mui/icons-material/Close';
import { useState } from 'react';
import { useForm, SubmitHandler, useFieldArray } from 'react-hook-form';

import { ErrorText } from '@/components/ErrorText';
import { Input } from '@/components/Input';
import { InputLabel } from '@/components/InputLabel';

import { FormInput } from '../types';

export function RegisterForm() {
  const [isMemberNumError, setIsMemberNumError] = useState(false);

  const {
    register,
    handleSubmit,
    formState: { errors },
    control,
  } = useForm<FormInput>({
    defaultValues: {
      teamName: '',
      members: [{ name: '' }],
    },
  });

  const { fields, append, remove } = useFieldArray({
    control,
    name: 'members',
  });

  const onSubmit: SubmitHandler<FormInput> = (data) => {
    if (fields.length < 2) {
      setIsMemberNumError(true);
    } else {
      console.log('OK!!');
    }
    console.log(data);
  };

  return (
    <form
      // eslint-disable-next-line @typescript-eslint/no-misused-promises
      onSubmit={handleSubmit(onSubmit)}
      className='flex min-w-full flex-col'
    >
      <div className='pb-10'>
        <InputLabel label='チーム' />
        <Input
          name='teamName'
          type='text'
          register={register}
          validation={{
            required: 'チーム名を入力してね！',
            maxLength: {
              value: 20,
              message: 'チーム名は20文字以下までだよ！',
            },
          }}
          error={errors.teamName}
        />
      </div>
      <InputLabel label='メンバー' />
      <ul>
        {fields.map((field, index) => (
          <div key={field.id} className='pb-2'>
            <Input
              name={`members.${index}.name`}
              type='text'
              register={register}
              validation={{
                required: 'ニックネームを入力してね！',
                maxLength: {
                  value: 20,
                  message: 'ニックネームは20文字以下までだよ！',
                },
              }}
              error={errors?.members?.[index]?.name}
            >
              <CloseIcon
                fontSize='large'
                onClick={() => fields.length !== 1 && remove(index)}
                style={{ color: '#929292' }}
              />
            </Input>
          </div>
        ))}
      </ul>
      {isMemberNumError && (
        <div className='pb-2'>
          <ErrorText errorText='メンバーの数が足りないよ！？最低2人は必要だね！' />
        </div>
      )}
      <div>
        <button
          type='button'
          onClick={() => {
            append({ name: '' });
            setIsMemberNumError(false);
          }}
          className='text-gray-300 '
        >
          +メンバーを追加する
        </button>
      </div>
      <div className='w-full pt-20'>
        <button
          className=' w-full rounded-3xl border-2 border-custom-btnborder bg-custom-btn px-8 py-4 font-bold shadow-xl hover:bg-custom-btnhover'
          type='submit'
        >
          登録
        </button>
      </div>
    </form>
  );
}

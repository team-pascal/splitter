'use client';

import CloseIcon from '@mui/icons-material/Close';
import { useRouter } from 'next/navigation';
import { useState } from 'react';
import { useForm, SubmitHandler, useFieldArray } from 'react-hook-form';
import { toast } from 'react-toastify';

import { ErrorText } from '@/components/ErrorText';
import { Input } from '@/components/Input';
import { InputLabel } from '@/components/InputLabel';
import { GroupResponse } from '@/types/groups.type';

import { createRequestBody } from '../../logic/sing-up/createRequestBody';
import { FormInput } from '../../types/sing-up/top.type';

export function RegisterForm() {
  const [isMemberNumError, setIsMemberNumError] = useState(false);
  const router = useRouter();

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

  const onSubmit: SubmitHandler<FormInput> = async (data) => {
    if (fields.length < 2) {
      setIsMemberNumError(true);
      return;
    }

    const request = createRequestBody(data);

    try {
      const response = await fetch('api/groups', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(request),
      });

      const responseData = (await response.json()) as GroupResponse;

      if (response.ok) {
        router.replace(responseData.group.id);
      } else {
        toast.error('グループ作成に失敗しました。。');
      }
    } catch {
      toast.error('グループ作成に失敗しました。。');
    }
  };

  return (
    <form
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

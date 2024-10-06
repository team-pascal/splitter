import { useContext } from 'react';
import {
  FieldError,
  FieldValues,
  Path,
  RegisterOptions,
  UseFormRegister,
} from 'react-hook-form';

import { ErrorText } from '@/components/ErrorText';
import { MemberContext } from '@/context';

type Props<T extends FieldValues> = {
  name: Path<T>;
  register?: UseFormRegister<T>;
  validation?: RegisterOptions<T, Path<T>>;
  error?: FieldError;
};

export function Select<T extends FieldValues>(props: Props<T>) {
  const { name, register, validation, error } = props;
  const memberList = useContext(MemberContext);

  return (
    <>
      <select
        id={name}
        {...(register && register(name, validation))}
        className='md: mb-2 block w-full border-b-2 border-gray-300 bg-transparent py-2  text-lg text-white outline-none'
      >
        <option value=''></option>
        {memberList.map((member, index) => (
          <option key={index} value={member.name}>
            {member.name}
          </option>
        ))}
      </select>
      {error && <ErrorText errorText={error.message} />}
    </>
  );
}

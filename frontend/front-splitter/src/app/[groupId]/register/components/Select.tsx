import { ErrorText } from '@/components/ErrorText';
import { MemberContext } from '@/context';
import { useContext } from 'react';
import {
  FieldError,
  FieldValues,
  Path,
  RegisterOptions,
  UseFormRegister,
} from 'react-hook-form';

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
        className="bg-transparent outline-none border-b-2 border-gray-300 text-white block w-full py-2  mb-2 md: text-lg"
      >
        <option value=""></option>
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

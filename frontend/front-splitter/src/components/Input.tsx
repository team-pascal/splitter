import {
  FieldError,
  FieldValues,
  Path,
  RegisterOptions,
  UseFormRegister,
} from 'react-hook-form';

import { ErrorText } from '@/components/ErrorText';

type Props<T extends FieldValues> = {
  name: Path<T>;
  type: string;
  children?: React.ReactNode;
  register?: UseFormRegister<T>;
  validation?: RegisterOptions<T, Path<T>>;
  error?: FieldError;
  onChange?: (value: string) => void;
};

export function Input<T extends FieldValues>(props: Props<T>) {
  const { name, type, children, register, validation, error, onChange } = props;

  return (
    <>
      <div className='mb-2 flex flex-row items-center border-b-2 border-gray-300'>
        <input
          type={type}
          id={name}
          className='md: block grow bg-transparent py-2 text-lg leading-none text-white outline-none'
          {...(register && register(name, validation))}
          onChange={(e) => onChange && onChange(e.target.value)}
        />
        {children}
      </div>
      {error && <ErrorText errorText={error.message} />}
    </>
  );
}

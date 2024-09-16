import { ErrorText } from '@/components/ErrorText';
import {
  FieldError,
  FieldValues,
  Path,
  RegisterOptions,
  UseFormRegister,
} from 'react-hook-form';

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
      <div className="flex flex-row border-b-2 border-gray-300 mb-2 items-center">
        <input
          type={type}
          id={name}
          className="bg-transparent outline-none text-white block grow py-2 leading-none md: text-lg"
          {...(register && register(name, validation))}
          onChange={(e) => onChange && onChange(e.target.value)}
        />
        {children}
      </div>
      {error && <ErrorText errorText={error.message} />}
    </>
  );
}

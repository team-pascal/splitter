type Props = {
  errorText?: string;
};
export function ErrorText({ errorText }: Props) {
  return <p className="text-red-500">{errorText}</p>;
}

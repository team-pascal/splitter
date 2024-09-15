export type Mode = 'split' | 'inAdvance';

export type User = { id: number; name: string };

export type MemberList = Array<User>;

export type PaymentInfo = User & {
  fee: number;
};

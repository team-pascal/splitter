export type PaymentInfo = User & {
  fee: number;
};

export type User = { id: number; name: string };

export type MemberList = Array<User>;

export type Mode = 'split' | 'inAdvance';

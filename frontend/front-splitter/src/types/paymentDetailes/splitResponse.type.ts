interface Counterparty {
  userId: string;
  amount: number;
}

interface PaymentDetail {
  userId: string;
  createdAt: string;
  updatedAt: string;
  deletedAt: string | null;
  amount: number;
  from: Counterparty[];
}

export interface SplitResponse {
  lessees: PaymentDetail[];
  split: PaymentDetail[];
}

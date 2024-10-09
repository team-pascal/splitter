export interface Payment {
  id: string;
  title: string;
  amount: number;
  done: boolean;
  group_id: string;
  genre: 'split' | 'replacement';
  created_at: string;
  updated_at: string;
  deleted_at: string;
}

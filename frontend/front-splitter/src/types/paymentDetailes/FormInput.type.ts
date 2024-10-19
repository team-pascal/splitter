export type FormInput = {
  title: string;
  cost: string;
  paymentMembers: { name: string }[];
  paymentCost: { cost: string }[];
  splitMembers: { name: string }[];
  inAdvanceMembers: { name: string }[];
  inAdvanceCost: { cost: string }[];
};

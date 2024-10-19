import { GroupRequest } from '@/types/groups.type';

import { FormInput } from '../../types/singUp/top.type';

export const createRequestBody = (data: FormInput): GroupRequest => {
  return {
    name: data.teamName,
    users: data.members.map((user) => user.name),
  };
};

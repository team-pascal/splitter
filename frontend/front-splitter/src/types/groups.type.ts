interface Group {
  id: string;
  name: string;
  created_at: string;
  updated_at: string;
  deleted_at: string;
}

interface User {
  id: string;
  name: string;
  group_id: string;
  created_at: string;
  updated_at: string;
  deleted_at: string;
}

export interface GroupResponse {
  group: Group;
  users: Array<User>;
}

export interface GroupRequest {
  name: string;
  users: Array<string>;
}

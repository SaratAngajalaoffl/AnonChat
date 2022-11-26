export type Room = {
  id: number;
  topic: string;
  participants: null | [{ name: string }];
};

export type Message = {
  sender: string;
  message: string;
};

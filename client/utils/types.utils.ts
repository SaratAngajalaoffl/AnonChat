export type Room = {
  id: number;
  topic: string;
  Participants: null | [{ name: string }];
};

export type Message = {
  sender: string;
  message: string;
};

export type Room = {
  id: number;
  topic: string;
  participants: null | User[];
};

export type SocketData = MessageData | UserData;

export type MessageData = {
  type: "MESSAGE";
  data: Message;
};

export type UserData = {
  type: "PERSON_JOINED" | "PERSON_LEFT";
  data: User;
};

export type Message = {
  sender: string;
  message: string;
};

export type User = {
  name: string;
};

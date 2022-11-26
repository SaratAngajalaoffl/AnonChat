import { Room } from "../../utils/types.utils";

export type createRoomRequest = {
  topic: string;
};

export type createRoomResponse = {
  roomId: string;
};

export type joinRoomRequest = {
  roomId: number;
  uname: string;
};

export type getRoomRequest = string;

export type getRoomResponse = {
  participants: null | [{ name: string }];
  population: number;
  room: Room;
};

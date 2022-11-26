import axios from "axios";
import {
  CREATE_ROOM_URL,
  GET_ROOM_URL,
  JOIN_ROOM_URL,
} from "../../utils/url.utils";
import * as types from "./types";

export const createRoom = async (
  data: types.createRoomRequest
): Promise<{ data: types.createRoomResponse | null; error: null | Error }> => {
  const response = await axios.post(CREATE_ROOM_URL, data);

  if (response.status === 200) {
    return { data: response.data, error: null };
  } else {
    return { data: null, error: response.data };
  }
};

export const joinRoom = async ({ roomId, uname }: types.joinRoomRequest) => {
  const socket = new WebSocket(`${JOIN_ROOM_URL}/${roomId}/${uname}`);

  return socket;
};

export const getRoom = async (
  data: types.getRoomRequest
): Promise<{
  data: types.getRoomResponse | null;
  error: null | Error;
}> => {
  const response = await axios.get(`${GET_ROOM_URL}/${data}`);

  if (response.status === 202) {
    return { data: response.data, error: null };
  } else {
    return { data: null, error: response.data };
  }
};

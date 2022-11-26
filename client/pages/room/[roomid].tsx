import { useRouter } from "next/router";
import React, { KeyboardEventHandler, useEffect, useState } from "react";
import LoadingComponent from "../../components/loading/LoadingComponent";
import { getRoom, joinRoom } from "../../services/room";
import { Message, Room } from "../../utils/types.utils";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import LogoutIcon from "@mui/icons-material/Logout";
import SendIcon from "@mui/icons-material/Send";
import PersonIcon from "@mui/icons-material/PersonRounded";
import Image from "next/image";
import {
  Divider,
  Grid,
  Hidden,
  InputAdornment,
  OutlinedInput,
} from "@mui/material";
import {
  ChatBubblePrimary,
  ChatBubbleSecondary,
  RootContainer,
} from "./styles";

type Props = {};

function ChatRoomPage({}: Props) {
  const router = useRouter();
  const { roomid } = router.query;
  const [room, setRoom] = useState<Room | null>();
  const [population, setPopulation] = useState<number>(0);
  const [currMsg, setCurrMsg] = useState<string>("");
  const [connection, setConnection] = useState<{
    name: string;
    socket: WebSocket;
  } | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [messages, setMessages] = useState<Message[]>([]);

  const handleRecieveMessage = (message: string) => {
    const parsedMessage: Message = JSON.parse(message);

    setMessages((prevState) => prevState.concat(parsedMessage));
  };

  const handleSendMessage = () => {
    if (!connection) return;
    if (currMsg.length < 1) return;

    connection.socket.send(currMsg);
    setCurrMsg("");
  };

  const handleKeyDown: KeyboardEventHandler = (e) => {
    if (e.key === "Enter") handleSendMessage();
  };

  useEffect(() => {
    (async () => {
      setIsLoading(true);

      if (!roomid || typeof roomid !== "string") return console.log("No data");

      const { data, error } = await getRoom(roomid);

      if (error) return console.log(error);

      if (!data) return console.log("No Data");

      setRoom(data.room);
      setPopulation(data.population);

      const uname = prompt("Enter username");

      if (!uname) return console.log("Please Enter Name");

      const socket = await joinRoom({ roomId: data.room.id, uname: uname });

      socket.addEventListener("message", (e) => handleRecieveMessage(e.data));

      setConnection({ name: uname, socket });

      setIsLoading(false);
    })();
  }, [roomid]);

  if (isLoading || !room) return <LoadingComponent />;

  return (
    <RootContainer sx={{ flexGrow: 1 }} display="flex" flexDirection="column">
      <Grid
        container
        direction="row"
        style={{ backgroundColor: "#121212", height: "10vh" }}
        alignItems="center"
        padding={2}
      >
        <IconButton
          edge="start"
          color="inherit"
          aria-label="menu"
          sx={{ mr: 2 }}
        >
          <Image
            src="/images/logo.png"
            alt="Logo"
            width={512 / 16}
            height={512 / 16}
          />
        </IconButton>
        <Typography
          variant="h6"
          color="inherit"
          component="div"
          sx={{ flexGrow: 1 }}
        >
          Joined room {room.topic} as {connection?.name}
        </Typography>
        <IconButton
          size="large"
          aria-label="display more actions"
          edge="end"
          color="inherit"
          onClick={() => router.replace("/")}
        >
          <LogoutIcon />
        </IconButton>
      </Grid>
      <Grid container direction="row" flexGrow={1}>
        <Grid
          item
          lg={9}
          xs={12}
          container
          direction="column"
          justifyContent="flex-end"
          alignItems="stretch"
        >
          <Grid
            item
            xs={11}
            container
            direction="column"
            justifyContent="flex-end"
            overflow="scroll"
            alignItems="stretch"
            style={{ height: "80vh", overflow: "scroll" }}
          >
            {messages?.map((message, index) => (
              <Grid
                item
                key={index}
                container
                justifyContent={
                  message.sender === connection?.name
                    ? "flex-end"
                    : "flex-start"
                }
              >
                <Grid item xs={8}>
                  {message.sender === connection?.name ? (
                    <ChatBubblePrimary>{message.message}</ChatBubblePrimary>
                  ) : (
                    <ChatBubbleSecondary>{message.message}</ChatBubbleSecondary>
                  )}
                  <Typography variant="caption" style={{ paddingLeft: 20 }}>
                    {message.sender !== connection?.name && message.sender}
                  </Typography>
                </Grid>
              </Grid>
            ))}
          </Grid>
          <Grid item xs={1} style={{ height: "10vh" }}>
            <OutlinedInput
              fullWidth
              onKeyDown={handleKeyDown}
              color="primary"
              value={currMsg}
              onChange={(e) => setCurrMsg(e.target.value)}
              endAdornment={
                <InputAdornment position="end">
                  <IconButton
                    aria-label="toggle password visibility"
                    onClick={handleSendMessage}
                    edge="end"
                  >
                    <SendIcon />
                  </IconButton>
                </InputAdornment>
              }
            />
          </Grid>
        </Grid>
        <Hidden xsDown>
          <Divider flexItem orientation="vertical" light color="#222222" />
          <Grid
            item
            lg={2}
            style={{ padding: 10 }}
            container
            direction="column"
            spacing={2}
            alignItems="stretch"
          >
            <Grid
              item
              container
              direction="row"
              spacing={1}
              alignItems="center"
            >
              <Grid item>
                <Typography variant="h6">{population} participants</Typography>
              </Grid>
            </Grid>
            <Grid
              item
              container
              direction="row"
              spacing={1}
              alignItems="center"
            >
              <Grid item>
                <PersonIcon />
              </Grid>
              <Grid item>
                <Typography variant="body1">{connection?.name}(You)</Typography>
              </Grid>
            </Grid>
            {room.Participants?.map((participant) => (
              <Grid
                key={participant.name}
                item
                container
                direction="row"
                spacing={1}
                alignItems="center"
              >
                <Grid item>
                  <PersonIcon />
                </Grid>
                <Grid item>
                  <Typography variant="body1">{participant.name}</Typography>
                </Grid>
              </Grid>
            ))}
          </Grid>
        </Hidden>
      </Grid>
    </RootContainer>
  );
}

export default ChatRoomPage;

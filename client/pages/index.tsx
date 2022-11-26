import { Button, Divider, TextField } from "@mui/material";
import Grid from "@mui/material/Grid";
import Typography from "@mui/material/Typography";
import Head from "next/head";
import Image from "next/image";
import { useRouter } from "next/router";
import { useState } from "react";
import LoadingComponent from "../components/loading/LoadingComponent";
import { createRoom } from "../services/room";
import styles from "../styles/Home.module.css";

export default function Home() {
  const [topic, setTopic] = useState<string>("");
  const [roomId, setRoomId] = useState<string>("");
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const router = useRouter();

  const handleCreateRoom = async () => {
    setIsLoading(true);

    const { data, error } = await createRoom({ topic });

    if (error) return console.log(error);

    if (!data) return;

    router.push(`/room/${data.roomId}`);
  };

  const handleJoinRoom = () => {
    router.push(`/room/${roomId}`);
  };

  if (isLoading) {
    return (
      <Grid container justifyContent="center" alignItems="center">
        <LoadingComponent />
      </Grid>
    );
  }

  return (
    <>
      <Head>
        <title>AnonChat</title>
        <meta
          name="description"
          content="An simple, encrypted, and anonymous chat room application."
        />
        <link rel="icon" href="/meta/favicon.ico" />
      </Head>

      <Grid
        container
        className={styles.main}
        direction="column"
        alignItems="center"
        justifyContent="center"
        spacing={5}
      >
        <Grid item xs={3}>
          <Image
            src="/images/logo.png"
            alt="Logo"
            width={512 / 4}
            height={512 / 4}
          />
        </Grid>
        <Grid item xs={1}>
          <Typography variant="h3">Welcome to AnonChat!</Typography>
        </Grid>
        <Grid item xs={3} container direction="row" alignItems="center">
          <Grid
            item
            xs
            container
            direction="column"
            alignItems="center"
            spacing={5}
          >
            <Grid item xs>
              <Typography variant="h4">Create a Room</Typography>
            </Grid>
            <Grid item xs>
              <TextField
                variant="outlined"
                label="Topic"
                color="primary"
                value={topic}
                onChange={(e) => setTopic(e.target.value)}
              />
            </Grid>
            <Grid item xs>
              <Button
                variant="contained"
                color="primary"
                fullWidth
                onClick={handleCreateRoom}
              >
                Create
              </Button>
            </Grid>
          </Grid>
          <Divider flexItem orientation="vertical" light color="#222222" />
          <Grid
            item
            xs
            container
            direction="column"
            alignItems="center"
            spacing={5}
          >
            <Grid item xs>
              <Typography variant="h4">Join a Room</Typography>
            </Grid>
            <Grid item xs>
              <TextField
                value={roomId}
                onChange={(e) => setRoomId(e.target.value)}
                variant="outlined"
                label="Room ID"
                color="primary"
              />
            </Grid>
            <Grid item xs>
              <Button
                variant="contained"
                color="secondary"
                fullWidth
                onClick={handleJoinRoom}
              >
                Join
              </Button>
            </Grid>
          </Grid>
        </Grid>
      </Grid>
    </>
  );
}

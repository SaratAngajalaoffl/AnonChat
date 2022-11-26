import { Box, Typography } from "@mui/material";
import { styled } from "@mui/material/styles";

export const RootContainer = styled(Box)(({ theme }) => ({
  height: "100vh",
}));

export const ChatBubblePrimary = styled(Typography)(({ theme }) => ({
  padding: 10,
  backgroundColor: theme.palette.primary.dark,
  margin: 10,
  borderRadius: 10,
  borderBottomRightRadius: 0,
}));

export const ChatBubbleSecondary = styled(Typography)(({ theme }) => ({
  padding: 10,
  backgroundColor: theme.palette.secondary.dark,
  margin: 10,
  borderRadius: 10,
  borderBottomLeftRadius: 0,
}));

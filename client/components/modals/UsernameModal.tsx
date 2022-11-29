import Dialog from "@mui/material/Dialog";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import DialogActions from "@mui/material/DialogActions";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import { useState } from "react";

type Props = {
  isOpen: boolean;
  handleCancel: () => void;
  handleJoin: (username: string) => void;
};

function UsernameModal({ isOpen, handleCancel, handleJoin }: Props) {
  const [username, setUsername] = useState<string>("");

  return (
    <Dialog open={isOpen} fullWidth>
      <DialogTitle>Enter Username</DialogTitle>
      <DialogContent>
        <TextField
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          fullWidth
        />
      </DialogContent>
      <DialogActions>
        <Button onClick={handleCancel}>Cancel</Button>
        <Button onClick={() => handleJoin(username)}>Join</Button>
      </DialogActions>
    </Dialog>
  );
}

export default UsernameModal;

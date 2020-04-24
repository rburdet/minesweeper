import React, { useState } from "react";
import "./App.css";
import Board from "./component/board";
import Button from "@material-ui/core/Button";
import DialogTitle from "@material-ui/core/DialogTitle";
import DialogActions from "@material-ui/core/DialogActions";
import DialogContent from "@material-ui/core/DialogContent";
import Dialog from "@material-ui/core/Dialog";
import TextField from "@material-ui/core/TextField";
import { DEFAULT_ROWS, DEFAULT_COLUMNS, DEFAULT_MINES } from "./constants";

function ParametersDialog({ settedValues, open, onClose }) {
  const [values, setValues] = useState(settedValues);

  const handleClose = () => {
    onClose(values);
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setValues({ ...values, [name]: value });
  };

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      aria-labelledby="form-dialog-title"
    >
      <DialogTitle id="form-dialog-title">Start your game !</DialogTitle>
      <DialogContent>
        <TextField
          autoFocus
          margin="dense"
          name="mines"
          type="number"
          label="mines"
          value={values.mines}
          onChange={handleInputChange}
          fullWidth
        />
        <TextField
          autoFocus
          margin="dense"
          name="rows"
          type="number"
          label="rows"
          value={values.rows}
          onChange={handleInputChange}
          fullWidth
        />
        <TextField
          autoFocus
          margin="dense"
          name="columns"
          type="number"
          label="columns"
          value={values.columns}
          onChange={handleInputChange}
          fullWidth
        />
        <TextField
          autoFocus
          margin="dense"
          name="name"
          label="Game name"
          helperText="You can recover your game with this name! - There cant be 2 games with the same name"
          fullWidth
          onChange={handleInputChange}
          value={values.name}
        />
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose} color="primary">
          OK!
        </Button>
      </DialogActions>
    </Dialog>
  );
}

function App() {
  const [open, setOpen] = useState(true);
  const [selectedValues, setSelectedValues] = useState({
    rows: DEFAULT_ROWS,
    columns: DEFAULT_COLUMNS,
    mines: DEFAULT_MINES,
    name: "",
  });

  const handleClose = (value) => {
    setOpen(false);
    setSelectedValues(value);
  };

  return (
    <div className="App">
      <div>
        <h3> Minesweeep </h3>
        {!open && (
          <Board
            columns={selectedValues.columns}
            rows={selectedValues.rows}
            mines={selectedValues.mines}
            name={selectedValues.name}
          />
        )}
      </div>
      <ParametersDialog
        settedValues={selectedValues}
        open={open}
        onClose={handleClose}
      />
    </div>
  );
}

export default App;

import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles((theme) => ({
  cell: {
    background: "#7b7b7b",
    border: "1px solid #fff",
    float: "left",
    lineHeight: `${theme.spacing(4)}px`,
    padding: 0,
    height: `${theme.spacing(4)}px`,
    textAlign: "center",
    width: `${theme.spacing(4)}px`,
    cursor: "pointer",
    borderRadius: "5px",
    color: "#fff",
    fontWeight: 600,
  },
  clicked: {
    background: "#bbbbbb",
  },
}));

function Cell({ i, j, onClick, cell }) {
  const classes = useStyles();
  const [clicked, setClicked] = useState(false);

  const cellClick = (action) => {
    if (!action) setClicked(true);
    if (action === "flag" && cell.status === "flag") {
      onClick(i, j, "clean");
    } else {
      onClick(i, j, action);
    }
  };

  const openedCell = () => (
    <div key={i * j} className={`${classes.cell} ${classes.clicked}`}>
      {cell.hasMine ? "💣" : cell.adjancetMines || ""}
    </div>
  );

  const closedCell = () => (
    <div
      onClick={() => cellClick()}
      onContextMenu={(e) => {
        e.preventDefault();
        cellClick("flag");
        return false;
      }}
      key={i * j}
      className={classes.cell}
    ></div>
  );

  const flaggedCell = () => (
    <div
      key={i * j}
      onContextMenu={(e) => {
        e.preventDefault();
        cellClick("flag");
        return false;
      }}
      className={classes.cell}
    >
      🚩
    </div>
  );

  if (cell.status === "flag") return flaggedCell();
  if (cell.clicked || clicked) return openedCell();
  return closedCell();
}

export default Cell;

import React, { useEffect, useState } from "react";
import axios from "axios";
import { makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import Cell from "./cell";

const useStyles = makeStyles(() => ({
  root: {
    flexGrow: 1,
  },
}));

const URL = "https://rburdet-minesweeper-api.herokuapp.com/api";
// const URL = "http://localhost:8080/api";

const gameUrl = `${URL}/game`;

function Board({ mines, columns, rows, name }) {
  const classes = useStyles();
  const [boardName, setBoardName] = useState(name);
  const [board, setBoard] = useState();
  const [finished, setFinished] = useState(false);
  const gameParameters = {
    rows: +rows,
    columns: +columns,
    mines: +mines,
  }

  const restart = () => {
    setFinished(true);
    setFinished(false);
    setBoard();
  };

  const createBoard = async () => {
    const {
      data: { id },
    } = await axios.post(gameUrl, gameParameters);
    setBoardName(id);
  };

  useEffect(() => {
    if (!finished && !boardName) createBoard();
  }, [finished]);

  useEffect(() => {
    const getBoard = async () => {
      const { data } = await axios.get(`${gameUrl}/${boardName}`);
      if (!data) await createBoard()
      setBoard(data);
    };
    if (boardName) getBoard();
  }, [boardName]);

  const onClick = async (i, j, action) => {
    const { data } = await axios.post(`${gameUrl}/${boardName}/click`, {
      i,
      j,
      action,
    });
    if (data.status !== "playing") {
      alert(data.status);
      restart();
    } else setBoard(data);
  };

  return (
    <div className={classes.root}>
      <h4> You are playing on {boardName} </h4>
      <Grid container spacing={1}>
        {board &&
          board.board.map((row, i) => (
            <Grid container justify="center" item spacing={1}>
              {row.map((cell, j) => (
                <Grid item>
                  <Cell onClick={onClick} cell={cell} i={i} j={j} />
                </Grid>
              ))}
            </Grid>
          ))}
      </Grid>
    </div>
  );
}

export default Board;

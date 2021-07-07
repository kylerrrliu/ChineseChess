import React from "react";
import {Owner, Pieces, INIT_BOARD, INIT_OWNER} from "../types";

const defaultValue = {};
const BoardContext = React.createContext(defaultValue);

export class BoardProvider extends React.Component {
    state = {
        //this board is valued as board[row][col] and is counted from left to right, down to up.
        board: INIT_BOARD,
        owner: INIT_OWNER,
        //piece owner of the board.
    }
    render() {
        return (
            <BoardContext.Provider value = {this.state}>
                {this.props.children}
            </BoardContext.Provider>
        )
    }
}

export const BoardConsumer = BoardContext.Consumer





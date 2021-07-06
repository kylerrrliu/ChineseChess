import {Owner, Pieces} from "./pieces";
import {INIT_BOARD, INIT_OWNER} from "./constant";

export class Board {
    //this board is valued as board[row][col] and is counted from left to right, down to up.
    board: Pieces[][];
    //piece owner of the board.
    owner: Owner[][];
    constructor() {
        this.board = INIT_BOARD;
        this.owner = INIT_OWNER;
    }
}
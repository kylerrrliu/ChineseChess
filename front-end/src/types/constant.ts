import {Owner, Pieces} from "./pieces";

export const BOARD_WIDTH: number = 9;

export const BOARD_HEIGHT: number = 10;

export const INIT_BOARD: Pieces[][] =[
    [Pieces.CHE, Pieces.MA, Pieces.XIANG, Pieces.SHI, Pieces.SHUAI, Pieces.SHI, Pieces.XIANG, Pieces.MA, Pieces.CHE],
    new Array(9).fill(Pieces.EMPTY),
    [Pieces.EMPTY, Pieces.PAO,Pieces.EMPTY,Pieces.EMPTY,Pieces.EMPTY,Pieces.EMPTY,Pieces.EMPTY,Pieces.PAO,Pieces.EMPTY],
    [Pieces.ZU,Pieces.EMPTY,Pieces.ZU,Pieces.EMPTY,Pieces.ZU,Pieces.EMPTY,Pieces.ZU,Pieces.EMPTY,Pieces.ZU],
    new Array(9).fill(Pieces.EMPTY),
    new Array(9).fill(Pieces.EMPTY),
    [Pieces.ZU,Pieces.EMPTY,Pieces.ZU,Pieces.EMPTY,Pieces.ZU,Pieces.EMPTY,Pieces.ZU,Pieces.EMPTY,Pieces.ZU],
    [Pieces.EMPTY, Pieces.PAO,Pieces.EMPTY,Pieces.EMPTY,Pieces.EMPTY,Pieces.EMPTY,Pieces.EMPTY,Pieces.PAO,Pieces.EMPTY],
    new Array(9).fill(Pieces.EMPTY),
    [Pieces.CHE, Pieces.MA, Pieces.XIANG, Pieces.SHI, Pieces.SHUAI, Pieces.SHI, Pieces.XIANG, Pieces.MA, Pieces.CHE]
];

export const INIT_OWNER: Owner[][] = [
    new Array(9).fill(Owner.RED),
    new Array(9).fill(Owner.NONE),
    [Owner.NONE,Owner.RED,Owner.NONE,Owner.NONE,Owner.NONE,Owner.NONE,Owner.NONE,Owner.RED,Owner.NONE],
    [Owner.RED,Owner.NONE,Owner.RED,Owner.NONE,Owner.RED,Owner.NONE,Owner.RED,Owner.NONE,Owner.RED],
    new Array(9).fill(Owner.NONE),
    new Array(9).fill(Owner.NONE),
    [Owner.BLACK,Owner.NONE,Owner.BLACK,Owner.NONE,Owner.BLACK,Owner.NONE,Owner.BLACK,Owner.NONE,Owner.BLACK],
    [Owner.NONE,Owner.BLACK,Owner.NONE,Owner.NONE,Owner.NONE,Owner.NONE,Owner.NONE,Owner.BLACK,Owner.NONE],
    new Array(9).fill(Owner.NONE),
    new Array(9).fill(Owner.BLACK),
]
import React from "react";

const defaultValue = {};
const BoardContext = React.createContext(defaultValue);

export class BoardProvider extends React.Component {
    state = {
        board: true,
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





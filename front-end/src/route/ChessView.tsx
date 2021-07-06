import {BoardConsumer} from "./boardContext";

const ChessView = () => (
    <BoardConsumer>
        {() => <div>123</div>}
    </BoardConsumer>
)

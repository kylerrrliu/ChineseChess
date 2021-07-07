package server

import (
	"context"
	"fmt"
	"github.com/rabcat/cn-chess-backend/be/ai"
)

func (s *CNChessBackendServer) GetNextMove(
	ctx context.Context,
	board [][]string,
) ([][]string, error) {

	newBoard := [][]string{
		{board[0][1]},
		{board[0][0]},
	}

	cnChessAI := ai.NewAI(100)
	fmt.Println(cnChessAI.Plus(3))

	return newBoard, nil
}
package proxy

import (
	ctx "context"
)

type CNChessBackendServerProxy struct {
	Proxy *Proxy
}

func (p *CNChessBackendServerProxy) GetNextMove(
	context ctx.Context,
	board [][]string,
) ([][]string, error) {
	if context == nil {
		context = ctx.Background()
	}
	result, err := p.Proxy.Invoke("GetNextMove", context, board)
	resp := result[0].([][]string)
	return resp, err
}
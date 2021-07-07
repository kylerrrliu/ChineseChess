package be

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rabcat/cn-chess-backend/be/proxy"
	"github.com/rabcat/cn-chess-backend/be/server"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	getNextMovePattern = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"cn-chess-backend","get-next-move"}, "", runtime.AssumeColonVerbOpt(true)))
)

type CNChessBackendService struct {
	HTTPServer *http.Server
	Mux        *mux.Router
	HTTPPort       int
}

func NewCNChessBackendService(port int) *CNChessBackendService {
	// create grpc_gateway reverse proxy
	gwmux := runtime.NewServeMux()
	mx := mux.NewRouter()
	httpServer := &http.Server{Addr: fmt.Sprintf("localhost:%d", port), Handler: mx}

	s := &CNChessBackendService{
		HTTPServer: httpServer,
		Mux:        mx,
		HTTPPort:   port,
	}
	sv := server.CNChessBackendServer{}
	pxy := proxy.NewProxy(&sv, &proxy.MetricsProxy{})
	serverProxy := proxy.CNChessBackendServerProxy{Proxy: pxy}

	s.RegisterHTTPHandlers(gwmux, serverProxy)
	mx.PathPrefix("/").Handler(gwmux)

	return s
}

func (s *CNChessBackendService) RegisterHTTPHandlers(gwmux *runtime.ServeMux, cnChessBackendServer proxy.CNChessBackendServerProxy) {
	gwmux.Handle("POST", getNextMovePattern, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		badRequestLogger := func(errTag string, internalErr error) {
			http.Error(w, internalErr.Error(), http.StatusBadRequest)
		}
		internalErrorLogger := func(errTag string, internalErr error) {
			http.Error(w, internalErr.Error(), http.StatusInternalServerError)
		}

		requestBody, readErr := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if readErr != nil {
			badRequestLogger("read", readErr)
			return
		}

		var board [][]string
		unMarshalErr := json.Unmarshal(requestBody, &board)
		if unMarshalErr != nil {
			badRequestLogger("deserialize", unMarshalErr)
			return
		}

		newBoard, handleErr := cnChessBackendServer.GetNextMove(context.Background(), board)
		if handleErr != nil {
			internalErrorLogger("handle", handleErr)
		}

		respData, marshalErr := json.Marshal(newBoard)
		if marshalErr != nil {
			internalErrorLogger("serialize", marshalErr)
		}

		w.Header().Set("Content-Type", "application/json")
		_, writeErr := w.Write(respData)
		if writeErr != nil {
			internalErrorLogger("write", writeErr)
			return
		}
	})
}

func (s *CNChessBackendService) Start() error {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	quit := make(chan error)
	go func() {
		log.Printf("http server listen to :%d", s.HTTPPort)
		if err := s.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			quit <- err
			log.Fatalf("Failed to start http server: %v", err)
		}
	}()

	select {
	// wait on kill signal
	case <-ch:
		log.Println("Interrupt caught, gracefully shutting down")
	// wait on context cancel
	case err := <-quit:
		if err != http.ErrServerClosed {
			return err
		}
	}

	return nil
}
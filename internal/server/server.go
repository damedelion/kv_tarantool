package server

import (
	"fmt"
	"net/http"

	"github.com/damedelion/kv_tarantool/config"
	"github.com/damedelion/kv_tarantool/internal/middleware"
	"github.com/damedelion/kv_tarantool/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/tarantool/go-tarantool/v2"
)

type Server struct {
	config *config.Server
	db     *tarantool.Connection
	mux    *mux.Router
	Logger logger.Logger
}

func New(config *config.Server, db *tarantool.Connection, mux *mux.Router, logger logger.Logger) *Server {
	return &Server{config: config, db: db, mux: mux, Logger: logger}
}

func (s *Server) Run() {
	s.registerHandlers()

	loggedMux := middleware.LoggingMiddleware(s.Logger, s.mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: loggedMux,
	}

	fmt.Println("Server is listening on", server.Addr)
	server.ListenAndServe()
}

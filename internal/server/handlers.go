package server

import (
	"github.com/damedelion/kv_tarantool/internal/kv/delivery/http"
	"github.com/damedelion/kv_tarantool/internal/kv/repository/tarantool"
	"github.com/damedelion/kv_tarantool/internal/kv/usecase"
)

func (s *Server) registerHandlers() {
	repository := tarantool.New(s.db)
	usecase := usecase.New(repository)
	handlers := http.New(usecase)

	s.mux.HandleFunc("/kv/{key}", handlers.Get).Methods("GET")
	s.mux.HandleFunc("/kv", handlers.Post).Methods("POST")
	s.mux.HandleFunc("/kv/{key}", handlers.Put).Methods("PUT")
	s.mux.HandleFunc("/kv/{key}", handlers.Delete).Methods("DELETE")
}

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/damedelion/kv_tarantool/config"
	"github.com/damedelion/kv_tarantool/internal/server"
	"github.com/damedelion/kv_tarantool/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/tarantool/go-tarantool/v2"
	_ "github.com/tarantool/go-tarantool/v2/datetime"
	_ "github.com/tarantool/go-tarantool/v2/decimal"
	_ "github.com/tarantool/go-tarantool/v2/uuid"
)

func main() {
	time.Sleep(1 * time.Second)
	dialer := tarantool.NetDialer{
		Address:  "tarantool:3301",
		User:     "sampleuser",
		Password: "123456",
	}
	opts := tarantool.Opts{
		Timeout: time.Second,
	}

	conn, err := tarantool.Connect(context.Background(), dialer, opts)
	defer conn.CloseGraceful()
	if err != nil {
		fmt.Println("Connection refused:", err)
		return
	}

	config, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	logger := logger.New()

	server := server.New(&config.Server, conn, mux.NewRouter(), logger)
	server.Run()
}

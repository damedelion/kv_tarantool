package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/damedelion/kv_tarantool/internal/dto"
	"github.com/damedelion/kv_tarantool/internal/interrors"
	"github.com/damedelion/kv_tarantool/internal/kv"
	"github.com/damedelion/kv_tarantool/internal/models"
	"github.com/damedelion/kv_tarantool/internal/utils"
	"github.com/damedelion/kv_tarantool/pkg/logger"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type delivery struct {
	usecase kv.Usecase
	logger  logger.Logger
}

func New(usecase kv.Usecase) kv.Delivery {
	return &delivery{usecase: usecase}
}

func (d *delivery) Get(w http.ResponseWriter, r *http.Request) {
	requestID := r.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["key"])

	item, err := d.usecase.Get(key)
	if err != nil {
		switch err.(type) {
		case *interrors.KeyNotFound:
			w.WriteHeader(http.StatusNotFound)
			d.logger.Error("key not found", requestID, zap.String("msg", err.Error()))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			d.logger.Error("unknown error", requestID, zap.String("msg", err.Error()))
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		d.logger.Error("unknown error", requestID, zap.String("msg", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (d *delivery) Post(w http.ResponseWriter, r *http.Request) {
	requestID := r.Context().Value(utils.RequestIDKey{})
	item := &models.Item{}

	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		d.logger.Error("bad request", requestID, zap.String("msg", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key, _ := strconv.Atoi(item.Key)
	dtoItem := &dto.Item{Key: key, Value: item.Value}

	err := d.usecase.Post(dtoItem)
	if err != nil {
		switch err.(type) {
		case *interrors.DuplicateKey:
			w.WriteHeader(http.StatusConflict)
			d.logger.Error("duplicate key", requestID, zap.String("msg", err.Error()))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			d.logger.Error("unknown error", requestID, zap.String("msg", err.Error()))
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		d.logger.Error("unknown error", requestID, zap.String("msg", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (d *delivery) Put(w http.ResponseWriter, r *http.Request) {
	requestID := r.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["key"])

	item := &models.Item{}

	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		d.logger.Error("unknown error", requestID, zap.String("msg", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dtoItem := &dto.Item{Key: key, Value: item.Value}

	err := d.usecase.Put(dtoItem)
	if err != nil {
		switch err.(type) {
		case *interrors.KeyNotFound:
			w.WriteHeader(http.StatusNotFound)
			d.logger.Error("key not found", requestID, zap.String("msg", err.Error()))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			d.logger.Error("unknown error", requestID, zap.String("msg", err.Error()))
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		d.logger.Error("unknown error", requestID, zap.String("msg", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (d *delivery) Delete(w http.ResponseWriter, r *http.Request) {
	requestID := r.Context().Value(utils.RequestIDKey{})
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["key"])

	item, err := d.usecase.Delete(key)
	if err != nil {
		switch err.(type) {
		case *interrors.KeyNotFound:
			w.WriteHeader(http.StatusNotFound)
			d.logger.Error("key not found", requestID, zap.String("msg", err.Error()))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			d.logger.Error("unknown error", requestID, zap.String("msg", err.Error()))
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

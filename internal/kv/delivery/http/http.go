package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/damedelion/kv_tarantool/internal/dto"
	"github.com/damedelion/kv_tarantool/internal/kv"
	"github.com/damedelion/kv_tarantool/internal/models"
	"github.com/gorilla/mux"
)

type delivery struct {
	usecase kv.Usecase
}

func New(usecase kv.Usecase) kv.Delivery {
	return &delivery{usecase: usecase}
}

func (d *delivery) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["key"])

	item, err := d.usecase.Get(key)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
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

func (d *delivery) Post(w http.ResponseWriter, r *http.Request) {
	item := &models.Item{}

	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key, _ := strconv.Atoi(item.Key)
	dtoItem := &dto.Item{Key: key, Value: item.Value}

	err := d.usecase.Post(dtoItem)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		fmt.Println("failed to encode body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (d *delivery) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["key"])

	item := &models.Item{}

	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dtoItem := &dto.Item{Key: key, Value: item.Value}

	err := d.usecase.Put(dtoItem)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		fmt.Println("failed to encode body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (d *delivery) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["key"])

	item, err := d.usecase.Delete(key)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
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

package kv

import (
	"github.com/damedelion/kv_tarantool/internal/dto"
)

type Usecase interface {
	Get(key int) (*dto.Item, error)
	Post(item *dto.Item) error
	Put(item *dto.Item) error
	Delete(key int) (*dto.Item, error)
}

package tarantool

import (
	"github.com/damedelion/kv_tarantool/internal/dto"
	"github.com/damedelion/kv_tarantool/internal/interrors"
	"github.com/damedelion/kv_tarantool/internal/kv"
	"github.com/tarantool/go-tarantool/v2"
)

type repository struct {
	db *tarantool.Connection
}

func New(db *tarantool.Connection) kv.Repository {
	return &repository{db: db}
}

func (r *repository) Get(key int) (*dto.Item, error) {
	res := &[]dto.Item{}
	err := r.db.Do(
		tarantool.NewSelectRequest("kv").
			Key([]interface{}{key}),
	).GetTyped(res)

	if err != nil {
		return &dto.Item{}, &interrors.DuplicateKey{Err: err.Error()}
	}

	return &(*res)[0], nil
}

func (r *repository) Post(item *dto.Item) error {
	_, err := r.db.Do(
		tarantool.NewUpdateRequest("kv").
			Key([]interface{}{item.Key}).
			Operations(tarantool.NewOperations().Assign(1, item.Value)),
	).Get()

	if err != nil {
		return &interrors.DuplicateKey{Err: err.Error()}
	}

	return nil
}

func (r *repository) Put(item *dto.Item) error {
	_, err := r.db.Do(
		tarantool.NewReplaceRequest("kv").
			Tuple([]interface{}{item.Key, item.Value}),
	).Get()

	if err != nil {
		return &interrors.KeyNotFound{Err: err.Error()}
	}

	return nil
}

func (r *repository) Delete(key int) (*dto.Item, error) {
	res := &[]dto.Item{}
	err := r.db.Do(
		tarantool.NewDeleteRequest("kv").
			Key([]interface{}{key}),
	).GetTyped(res)

	if err != nil {
		return &dto.Item{}, &interrors.KeyNotFound{Err: err.Error()}
	}

	return &(*res)[0], nil
}

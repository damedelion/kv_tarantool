package usecase

import (
	"github.com/damedelion/kv_tarantool/internal/dto"
	"github.com/damedelion/kv_tarantool/internal/kv"
)

type usecase struct {
	repository kv.Repository
}

func New(repository kv.Repository) kv.Usecase {
	return &usecase{repository: repository}
}

func (u *usecase) Get(key int) (*dto.Item, error) {
	item, err := u.repository.Get(key)
	if err != nil {
		return &dto.Item{}, err
	}
	return item, nil
}

func (u *usecase) Post(item *dto.Item) error {
	if err := u.repository.Post(item); err != nil {
		return err
	}
	return nil
}

func (u *usecase) Put(item *dto.Item) error {
	if err := u.repository.Post(item); err != nil {
		return err
	}
	return nil
}

func (u *usecase) Delete(key int) (*dto.Item, error) {
	item, err := u.repository.Get(key)
	if err != nil {
		return &dto.Item{}, err
	}
	return item, nil
}

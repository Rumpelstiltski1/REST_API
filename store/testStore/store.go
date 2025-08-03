package testStore

import (
	"github.com/Rumpelstiltski1/restapi/internal/app/model"
	"github.com/Rumpelstiltski1/restapi/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
	return s.userRepository
}

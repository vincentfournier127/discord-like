package repository

import (
	"maps"
	"sync"

	"discord-like/internal/api/model"
)

type UserRepository struct {
	mu     sync.RWMutex
	users  map[int]model.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  make(map[int]model.User),
		nextID: 1,
	}
}

func (ur *UserRepository) getNextID() int {
	id := ur.nextID
	ur.nextID++
	return id
}

func (ur *UserRepository) Create(user *model.User) {
	ur.mu.Lock()
	defer ur.mu.Unlock()
	user.ID = ur.getNextID()
	ur.users[user.ID] = *user
}

func (ur *UserRepository) GetByID(id int) (model.User, bool) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()
	user, ok := ur.users[id]
	return user, ok
}

func (ur *UserRepository) GetAll() map[int]model.User {
	ur.mu.RLock()
	defer ur.mu.RUnlock()
	copy := maps.Clone(ur.users)
	return copy
}

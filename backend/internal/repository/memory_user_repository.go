package repository

import (
	"context"
	"errors"
	"maps"
	"slices"
	"sync"

	"discord-like/internal/api/model"
)

type MemoryUserRepository struct {
	mu     sync.RWMutex
	users  map[int]model.User
	nextID int
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users:  make(map[int]model.User),
		nextID: 1,
	}
}

func (ur *MemoryUserRepository) getNextID() int {
	id := ur.nextID
	ur.nextID++
	return id
}

func (ur *MemoryUserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()
	user.ID = ur.getNextID()
	ur.users[user.ID] = *user
	return user, nil
}

func (ur *MemoryUserRepository) GetByID(ctx context.Context, id int) (model.User, error) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()
	user, ok := ur.users[id]

	if ok {
		return user, nil
	} else {
		return model.User{}, errors.New("user not found")
	}
}

func (ur *MemoryUserRepository) GetAll(ctx context.Context) ([]model.User, error) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()
	copy := slices.Collect(maps.Values(ur.users))
	return copy, nil
}

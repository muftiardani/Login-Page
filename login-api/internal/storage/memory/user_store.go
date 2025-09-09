package memory

import (
	"errors"
	"login-api/internal/model"
	"sync"
)

type MemoryUserStore struct {
	users map[string]model.User
	lock  *sync.RWMutex
}

func NewMemoryUserStore() *MemoryUserStore {
	return &MemoryUserStore{
		users: make(map[string]model.User),
		lock:  &sync.RWMutex{},
	}
}

func (s *MemoryUserStore) GetUser(email string) (model.User, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	user, exists := s.users[email]
	return user, exists
}

func (s *MemoryUserStore) CreateUser(user model.User) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, exists := s.users[user.Email]; exists {
		return errors.New("email already exists")
	}
	s.users[user.Email] = user
	return nil
}
package models

import (
	"log"
	"sync"
)

type UserCollection struct {
	mu    sync.RWMutex
	users []*User
}

func (s *UserCollection) Add(user *User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	log.Println("Adding: ", user, len(s.users))
	s.users = append(s.users, user)
	log.Println("Added: ", user, len(s.users))
}

func (s *UserCollection) Remove(user *User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	log.Println("Removing: ", user, len(s.users))
	for index, u := range s.users {
		if u.C != user.C {
			continue
		}

		s.users = append(s.users[:index], s.users[index+1:]...)
		log.Println("Removed: ", user, len(s.users))
		break
	}
}

func (s *UserCollection) Get(user User) *User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for index, u := range s.users {
		if u.C == user.C {
			return s.users[index]
		}
	}
	return nil
}

func (s *UserCollection) GetAll() []*User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.users
}

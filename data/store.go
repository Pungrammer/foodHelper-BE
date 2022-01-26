package data

import (
	"foodHelper/shoppinglist"
	"foodHelper/users"
)

type Store struct {
}

func (s *Store) CreateOrUpdateUser(user *users.User) error {
	//TODO: Persist user and all API-Keys
	return nil
}

func (s *Store) GetUserByName(username string) (*users.User, error) {
	return &users.User{}, nil
}

func (s *Store) DeleteUser(userId string) error {
	return nil
}

func (s *Store) AddShoppingItem(userId string, item *shoppinglist.Item) error {
	return nil
}

func (s *Store) DeleteShoppingItem(itemId string) error {
	return nil
}

func (s *Store) CreateSession(userId string) (*Session, error) {
	return &Session{}, nil
}

func (s *Store) GetSession(sessionToken string) (*Session, error) {
	return &Session{}, nil
}

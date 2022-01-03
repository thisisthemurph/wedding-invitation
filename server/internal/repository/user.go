package repository

import (
	"errors"
	"wedding_api/internal/datastruct"
)

type UserQuery interface {
	CreateUser(user datastruct.Person) (*int64, error)
	GetUserById(userId int64) (*datastruct.Person, error)
}

type userQuery struct{}

func (u *userQuery) CreateUser(user datastruct.Person) (*int64, error) {
	// Code to add the user to the database

	var createdId int64 = 1
	return &createdId, nil
}

func (u *userQuery) GetUserById(userId int64) (*datastruct.Person, error) {
	// Logic to obtain the user
	mocks := []datastruct.Person{
		{
			ID: 1, 
			Name: "MiKe", 
			Email: "mike@email.com", 
			Password: "password",
			Verified: true,
			Role: datastruct.ADMIN,
		},
	}

	for _, user := range mocks {
		if user.ID == userId {
			return &user, nil
		}
	}

	return nil, errors.New("Could not find user with given ID.")
}
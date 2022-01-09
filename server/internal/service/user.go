package service

import (
	"wedding_api/internal/datastruct"
	"wedding_api/internal/repository"
)

type UserService interface {
	CreateUser(user datastruct.Person) (*datastruct.Person, error)
	GetUserById(userId int64) (*datastruct.Person, error)
}

type userService struct {
	dao repository.DAO
}

func NewUserService(dao repository.DAO) UserService {
	return &userService{dao: dao}
}

func (u *userService) CreateUser(user datastruct.Person) (*datastruct.Person, error) {
	usr, err := u.dao.NewUserQuery().CreateUser(user)
	if (err != nil) {
		return nil, err
	}
	
	person := &datastruct.Person{
		ID: usr.ID,
		Name: usr.Name,
		Email: usr.Email,
	}

	return person, nil
}

func (u *userService) GetUserById(userId int64) (*datastruct.Person, error) {
	user, err := u.dao.NewUserQuery().GetUserById(userId)

	if err != nil {
		return nil, err
	}

	person := &datastruct.Person{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}

	return person, nil
}
package service

import (
	"log"
	"wedding_api/internal/datastruct"
	"wedding_api/internal/repository"
)

type UserService interface {
	CreateUser(user datastruct.Person) (error)
	GetUserById(userId int64) (*datastruct.Person, error)
}

type userService struct {
	dao repository.DAO
}

func NewUserService(dao repository.DAO) UserService {
	return &userService{dao: dao}
}

func (u *userService) CreateUser(user datastruct.Person) error {
	err := u.dao.NewUserQuery().CreateUser(user)
	if (err != nil) {
		log.Println(err)
		return err
	}
	
	return nil
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
package service

import (
	"wedding_api/internal/datastruct"
	"wedding_api/internal/repository"
)

type UserService interface {
	GetUserById(userId int64) (*datastruct.Person, error)
}

type userService struct {
	dao repository.DAO
}

func NewUserService(dao repository.DAO) UserService {
	return &userService{dao: dao}
}

func (u *userService) GetUserById(userId int64) (*datastruct.Person, error) {
	user, err := u.dao.NewUserQuery().GetUserById(userId)

	if err != nil {
		return nil, err
	}

	p := &datastruct.Person{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}

	return p, nil
}
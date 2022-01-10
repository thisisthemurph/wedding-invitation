package repository

import (
	"context"
	"errors"
	"wedding_api/internal/datastruct"
	"wedding_api/internal/repository/models"
)

var ErrEmailAlreadyExists = errors.New("A user with that email address already exists")

type UserQuery interface {
	CreateUser(user datastruct.Person) (*models.User, error)
	GetUserById(userId int64) (*models.User, error)
	EmailExists(email string) (bool)
	DeleteUserById(userId int64) error
}

type userQuery struct{}

func (u *userQuery) CreateUser(user datastruct.Person) (*models.User, error) {
	userModel := &models.User{
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
		Verified: user.Verified,
		Role: user.Role,
	}

	emailExists := u.EmailExists(user.Email)
	if emailExists {
		return nil, ErrEmailAlreadyExists
	}

	_, err := DB.NewInsert().
		Model(userModel).
		Exec(context.Background())
	
	if err != nil {
		return nil, err
	}

	// Grab the newly created user
	newUser, err := u.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u *userQuery) GetUserById(userId int64) (*models.User, error) {
	user := new(models.User)
	err := DB.NewSelect().Model(user).Where("id = ?", userId).Scan(context.Background())

	if (err != nil) {
		return nil, err
	}

	return user, nil
}

func (u * userQuery) GetUserByEmail(email string) (*models.User, error) {
	userModel := new(models.User)
	err := DB.NewSelect().Model(userModel).Where("email = ?", email).Limit(1).Scan(context.Background())
	if err != nil {
		return userModel, err
	}

	return userModel, nil
}

func (u * userQuery) EmailExists(email string) bool {
	userModel := new(models.User)
	exists, _ := DB.NewSelect().Model(userModel).Where("email = ?", email).Limit(1).Exists(context.Background())
	return exists
}

func (u *userQuery) DeleteUserById(userId int64) error {
	userModel := new(models.User)
	_, err := DB.NewDelete().Model(userModel).Where("id = ?", userId).Exec(context.Background())
	return err
}
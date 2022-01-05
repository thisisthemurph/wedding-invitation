package repository

import (
	"context"
	"wedding_api/internal/datastruct"
	"wedding_api/internal/repository/models"
)

type UserQuery interface {
	CreateUser(user datastruct.Person) (error)
	GetUserById(userId int64) (*models.User, error)
}

type userQuery struct{}

func (u *userQuery) CreateUser(user datastruct.Person) error {
	userModel := &models.User{
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
		Verified: user.Verified,
		Role: user.Role,
	}

	_, err := DB.NewInsert().
		Model(userModel).
		Exec(context.Background())
	
	if err != nil {
		return err
	}

	return nil
}

func (u *userQuery) GetUserById(userId int64) (*models.User, error) {
	// Logic to obtain the user
	// mocks := []datastruct.Person{
	// 	{
	// 		ID: 1, 
	// 		Name: "MiKe", 
	// 		Email: "mike@email.com", 
	// 		Password: "password",
	// 		Verified: true,
	// 		Role: datastruct.ADMIN,
	// 	},
	// }

	// for _, user := range mocks {
	// 	if user.ID == userId {
	// 		return &user, nil
	// 	}
	// }

	user := new(models.User)
	ctx := context.Background()
	err := DB.NewSelect().Model(user).Where("id = ?", userId).Scan(ctx)

	if (err != nil) {
		return nil, err
	}

	return user, nil
}
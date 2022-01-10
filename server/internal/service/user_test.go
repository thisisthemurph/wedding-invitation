package service_test

import (
	"log"
	"testing"
	"wedding_api/internal/datastruct"
	"wedding_api/internal/repository"
	"wedding_api/internal/service"

	"github.com/uptrace/bun"
)

func TestCreateUser(t *testing.T) {
	userService := setup()
	p := makePerson()

	t.Run("creating user", func(t *testing.T) {
		created, err := userService.CreateUser(p)
		
		assertErrorIsNil(t, err)
		assertSamePerson(t, p, *created)
	})
	
	t.Run("cannot careate user with identical email", func(t *testing.T) {
		_, err := userService.CreateUser(p)

		assertHasError(t, err)
	})
}

func TestGetUserById(t *testing.T) {
	userService := setup()
	p := makePerson()

	t.Run("create and get user", func(t *testing.T) {
		created, _ := userService.CreateUser(p)
		user, err := userService.GetUserById(created.ID)
	
		assertErrorIsNil(t, err)
		assertSamePerson(t, p, *user)
	})

	t.Run("get unknown ID", func(t *testing.T) {
		_, err := userService.GetUserById(999)

		assertHasError(t, err)
	})
}

func TestDeleteUserById(t *testing.T) {
	userService := setup()
	p := makePerson()

	t.Run("create and delete user", func(t *testing.T) {
		created, _ := userService.CreateUser(p)
		deleted, err := userService.DeleteUserById(created.ID)

		assertErrorIsNil(t, err)
		assertSamePerson(t, p, *deleted)

		_, err = userService.GetUserById(created.ID)
		if err == nil {
			log.Fatal("User still present in the database following deletion")
		}
	})

	t.Run("cannot delete non-existant user", func(t *testing.T) {
		deleted, err := userService.DeleteUserById(999)

		assertHasError(t, err)
		if deleted != nil {
			log.Fatal("Expected returned user to be nil")
		}
	})
}

func assertHasError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		log.Fatal("Expected error, no error present.")
	}
}

func assertErrorIsNil(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		log.Fatalf(`Did not expect error response. Err: "%s"`, err)
	}
}

func assertSamePerson(t *testing.T, p datastruct.Person, o datastruct.Person) {
	t.Helper()
	if p.Name != o.Name {
		log.Fatalf("Expected %s got %s", p.Name, o.Name)
	}

	if p.Email != o.Email {
		log.Fatalf("Expected %s got %s", p.Email, o.Email)
	}
}

func makePerson() datastruct.Person {
	return datastruct.Person{
		Name: "Mike",
		Email: "mike@email.com",
		Password: "secretpassword",
		Verified: false,
		Role: datastruct.USER,
	}
}

func connectToDb() *bun.DB {
	db, err := repository.NewDevDB()
	if err != nil {
		log.Fatal("Cannot connect to database.")
	}

	return db
}

func setup() service.UserService {
	db := connectToDb()
	dao := repository.NewDAO(db)
	return service.NewUserService(dao)
}

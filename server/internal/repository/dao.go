package repository

type DAO interface {
	NewUserQuery() UserQuery
}

type TempType string

type dao struct{}

var DB TempType

func NewDAO(db TempType) DAO {
	DB = db
	return &dao{}
}

func NewDB() (TempType, error) {
	// Code for connecting to the database
	return TempType(""), nil
}

func (d *dao) NewUserQuery() UserQuery {
	return &userQuery{}
}
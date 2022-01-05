package repository

import (
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DAO interface {
	NewUserQuery() UserQuery
	NewEventQuery() EventQuery
}

type dao struct{}

var DB *bun.DB

func NewDAO(db *bun.DB) DAO {
	DB = db
	return &dao{}
}

func makePostgresConnectionString() string {
	uname    := "postgres"
	password := "GreenAli3n2001"
	host     := "localhost"
	port	 := "5432"
	name     := "wedding"

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", uname, password, host, port, name)
}

func NewDB() (*bun.DB, error) {
	dsn := makePostgresConnectionString()
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	DB := bun.NewDB(sqldb, pgdialect.New())

	// Create the tables
	// ctx := context.Background()
	// DB.NewCreateTable().Model((*models.User)(nil)).Exec(ctx)
	// DB.NewCreateTable().Model((*models.Event)(nil)).Exec(ctx)

	// fmt.Println("We have created the table")
	// fmt.Println(DB)

	return DB, nil
}

func (d *dao) NewUserQuery() UserQuery {
	return &userQuery{}
}

func (d *dao) NewEventQuery() EventQuery {
	return &eventQuery{}
}

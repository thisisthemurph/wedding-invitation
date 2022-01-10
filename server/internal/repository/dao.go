package repository

import (
	"context"
	"database/sql"
	"fmt"
	"wedding_api/internal/config"
	"wedding_api/internal/repository/models"

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

func makePostgresConnectionString(confif config.Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", 
		confif.DatabaseUsername,
		confif.DatabasePassword, 
		confif.DatabaseHost, 
		confif.DatabasePort, 
		confif.DatabaseName,
	)
}

func makeDB(config config.Config) (*bun.DB, error) {
	dsn := makePostgresConnectionString(config)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	DB := bun.NewDB(sqldb, pgdialect.New())

	return DB, nil
}

func NewDB() (*bun.DB, error) {
	config := config.LoadConfig()
	return makeDB(config)
}

func NewDevDB() (*bun.DB, error) {
	config := config.LoadDevConfig()
	DB, err := makeDB(config)
	
	// Drop the tables and create them again each time
	ctx := context.Background()
	DB.NewDropTable().Model((*models.User)(nil)).Exec(ctx)
	DB.NewCreateTable().Model((*models.User)(nil)).Exec(ctx)

	return DB, err
}

func (d *dao) NewUserQuery() UserQuery {
	return &userQuery{}
}

func (d *dao) NewEventQuery() EventQuery {
	return &eventQuery{}
}

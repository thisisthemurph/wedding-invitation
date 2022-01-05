package models

import (
	"github.com/uptrace/bun"
)

type Event struct {
	bun.BaseModel

	ID 			int64  `bun:"id,pk,autoincrement"`
	Name 		string `bun:",notnull"`
	Description string 

	// The owner of the event
	UserId 		int64
}
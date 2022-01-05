package models

import (
	"time"
	"wedding_api/internal/datastruct"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel

	ID 			int64 			`bun:"id,pk,autoincrement"`
	Name 		string 			`bun:",notnull"`
	Email 		string 			`bun:",notnull"`
	Password 	string 			`bun:",notnull"`
	Verified 	bool 			`bun:",nullzero,notnull,default:false"`
	Role 		datastruct.Role `bun:",nullzero,notnull,default:'user'"`
	CreatedAt 	time.Time 		`bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt 	time.Time 		`bun:",nullzero,notnull,default:current_timestamp"`

	// User has many events
	Events		[]*Event		`bun:"rel:has-many,join:id=user_id"`
}
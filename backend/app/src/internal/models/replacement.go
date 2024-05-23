package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Replacement struct {
	bun.BaseModel `bun:"table:replacements"`

	ID        uuid.UUID `bun:"id,pk,type:uuid"`
	Title     string    `bun:"title,notnull"`
	Amount    uint      `bun:"amount,notnull,type:integer"`
	Done      bool      `bun:"done,notnull"`
	GroupID   uuid.UUID `bun:"group_id,type:uuid"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime

	Group *Group `bun:"rel:belongs-to,join:group_id=id"`
}

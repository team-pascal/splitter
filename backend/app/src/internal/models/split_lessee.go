package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SplitLessee struct {
	bun.BaseModel `bun:"table:split_lessees"`

	SplitID   uuid.UUID `bun:"split_id,type:uuid"`
	UserID    uuid.UUID `bun:"user_id,type:uuid"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime

	User  *User  `bun:"rel:belongs-to,join:user_id=id"`
	Split *Split `bun:"rel:belongs-to,join:split_id=id"`
}

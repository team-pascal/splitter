package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ReplacementLessor struct {
	bun.BaseModel `bun:"table:replacement_lessors"`

	ReplacementID uuid.UUID `bun:"replacement_id,type:uuid"`
	UserID        uuid.UUID `bun:"user_id,type:uuid"`
	Amount        uint      `bun:"amount,type:integer"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt     bun.NullTime

	User        *User        `bun:"rel:belongs-to,join:user_id=id"`
	Replacement *Replacement `bun:"rel:belongs-to,join:replacement_id=id"`
}

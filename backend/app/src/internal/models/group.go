package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Group struct {
	bun.BaseModel `bun:"table:groups"`

	ID        uuid.UUID `bun:"id,pk,type:uuid"`
	Name      string    `bun:"name,notnull"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime
}

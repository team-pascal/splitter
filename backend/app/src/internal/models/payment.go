package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Payment struct {
	bun.BaseModel `bun:"table:payments"`

	ID        uuid.UUID `bun:"id,pk,type:uuid"`
	Title     string    `bun:"title,notnull"`
	Amount    uint      `bun:"amount,notnull,type:integer"`
	Done      bool      `bun:"done,notnull"`
	GroupID   uuid.UUID `bun:"group_id,type:uuid"`
	Genre     string    `bun:"genre"`
	CreatedAt time.Time `bun:",nullzero,notnull"`
	UpdatedAt time.Time `bun:",nullzero,notnull"`
	DeletedAt bun.NullTime
}

func (p Payment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID        string       `json:"id"`
		Title     string       `json:"title"`
		Amount    uint         `json:"amount"`
		Done      bool         `json:"done"`
		GroupID   string       `json:"group_id"`
		Genre     string       `json:"genre"`
		CreatedAt time.Time    `json:"created_at"`
		UpdatedAt time.Time    `json:"updated_at"`
		DeletedAt bun.NullTime `json:"deleted_at"`
	}{
		ID:        p.ID.String(),
		Title:     p.Title,
		Amount:    p.Amount,
		Done:      p.Done,
		GroupID:   p.GroupID.String(),
		Genre:     p.Genre,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: p.DeletedAt,
	})
}

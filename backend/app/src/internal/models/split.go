package models

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Split struct {
	bun.BaseModel `bun:"table:splits"`

	ID        uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	Title     string    `bun:"title,notnull"`
	Amount    uint      `bun:"amount,notnull,type:integer"`
	Done      bool      `bun:"done,notnull"`
	GroupID   uuid.UUID `bun:"group_id,type:uuid"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime

	Group *Group `bun:"rel:belongs-to,join:group_id=id"`
}

func (s Split) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID        string       `json:"id"`
		Title     string       `json:"title"`
		Amount    uint         `json:"amount"`
		Done      bool         `json:"done"`
		GroupID   string       `json:"group_id"`
		CreatedAt time.Time    `json:"created_at"`
		UpdatedAt time.Time    `json:"updated_at"`
		DeletedAt bun.NullTime `json:"deleted_at"`
	}{
		ID:        s.ID.String(),
		Title:     s.Title,
		Amount:    s.Amount,
		Done:      s.Done,
		GroupID:   s.GroupID.String(),
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		DeletedAt: s.DeletedAt,
	})
}

var _ bun.BeforeAppendModelHook = (*Split)(nil)

func (s *Split) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		s.CreatedAt = time.Now()
		s.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		s.UpdatedAt = time.Now()
	case *bun.DeleteQuery:
		s.DeletedAt = bun.NullTime{Time: time.Now()}
	}
	return nil
}

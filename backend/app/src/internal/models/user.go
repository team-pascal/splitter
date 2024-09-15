package models

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	Name      string    `bun:"name,notnull"`
	GroupID   uuid.UUID `bun:"group_id,type:uuid"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime

	Group *Group `bun:"rel:belongs-to,join:group_id=id"`
}

func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID        string       `json:"id"`
		Name      string       `json:"name"`
		GroupID   string       `json:"group_id"`
		CreatedAt time.Time    `json:"created_at"`
		UpdatedAt time.Time    `json:"updated_at"`
		DeletedAt bun.NullTime `json:"deleted_at"`
	}{
		ID:        u.ID.String(),
		Name:      u.Name,
		GroupID:   u.GroupID.String(),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	})
}

var _ bun.BeforeAppendModelHook = (*User)(nil)

func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		u.UpdatedAt = time.Now()
	case *bun.DeleteQuery:
		u.DeletedAt = bun.NullTime{Time: time.Now()}

	}
	return nil
}

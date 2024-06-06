package models

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SplitLessor struct {
	bun.BaseModel `bun:"table:split_lessors"`

	SplitID   uuid.UUID `bun:"split_id,type:uuid"`
	UserID    uuid.UUID `bun:"user_id,type:uuid"`
	Amount    uint      `bun:"amount,type:integer"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime

	User  *User  `bun:"rel:belongs-to,join:user_id=id"`
	Split *Split `bun:"rel:belongs-to,join:split_id=id"`
}

func (sl SplitLessor) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SplitID   string       `json:"split_id"`
		UserID    string       `json:"user_id"`
		Amount    uint         `json:"amount"`
		CreatedAt time.Time    `json:"created_at"`
		UpdatedAt time.Time    `json:"updated_at"`
		DeletedAt bun.NullTime `json:"deleted_at"`
	}{
		SplitID:   sl.SplitID.String(),
		UserID:    sl.UserID.String(),
		Amount:    sl.Amount,
		CreatedAt: sl.CreatedAt,
		UpdatedAt: sl.UpdatedAt,
		DeletedAt: sl.DeletedAt,
	})
}

var _ bun.BeforeAppendModelHook = (*SplitLessor)(nil)

func (sl *SplitLessor) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		sl.CreatedAt = time.Now()
		sl.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		sl.UpdatedAt = time.Now()
	case *bun.DeleteQuery:
		sl.DeletedAt = bun.NullTime{Time: time.Now()}
	}
	return nil
}

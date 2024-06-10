package models

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ReplacementLessee struct {
	bun.BaseModel `bun:"table:replacement_lessees"`

	ReplacementID uuid.UUID `bun:"replacement_id,pk,type:uuid"`
	UserID        uuid.UUID `bun:"user_id,pk,type:uuid"`
	Amount        uint      `bun:"amount,type:integer"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt     bun.NullTime

	User        *User        `bun:"rel:belongs-to,join:user_id=id"`
	Replacement *Replacement `bun:"rel:belongs-to,join:replacement_id=id"`
}

func (rl ReplacementLessee) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SplitID   string       `json:"split_id"`
		UserID    string       `json:"user_id"`
		Amount    uint         `json:"amount"`
		CreatedAt time.Time    `json:"created_at"`
		UpdatedAt time.Time    `json:"updated_at"`
		DeletedAt bun.NullTime `json:"deleted_at"`
	}{
		SplitID:   rl.ReplacementID.String(),
		UserID:    rl.UserID.String(),
		Amount:    rl.Amount,
		CreatedAt: rl.CreatedAt,
		UpdatedAt: rl.UpdatedAt,
		DeletedAt: rl.DeletedAt,
	})
}

var _ bun.BeforeAppendModelHook = (*ReplacementLessee)(nil)

func (rl *ReplacementLessee) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		rl.CreatedAt = time.Now()
		rl.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		rl.UpdatedAt = time.Now()
	case *bun.DeleteQuery:
		rl.DeletedAt = bun.NullTime{Time: time.Now()}
	}
	return nil
}

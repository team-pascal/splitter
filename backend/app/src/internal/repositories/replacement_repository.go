package repositories

import (
	"context"
	"splitter/internal/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ReplacementRepository struct {
	TX *bun.Tx
}

func (rr *ReplacementRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.Replacement, error) {
	replacement := new(models.Replacement)
	err := rr.TX.NewSelect().Model(replacement).Where("id = ?", id).Scan(ctx)
	return replacement, err
}

func (rr *ReplacementRepository) Create(ctx context.Context, title string, amount uint, groupID uuid.UUID) (*models.Replacement, error) {
	replacement := models.Replacement{Title: title, GroupID: groupID, Amount: amount}

	_, err := rr.TX.NewInsert().Model(&replacement).Exec(ctx)
	return &replacement, err
}

func (rr *ReplacementRepository) Update(ctx context.Context, title string, amount uint, id uuid.UUID) (*models.Replacement, error) {
	replacement := new(models.Replacement)
	replacement.Title = title
	replacement.Amount = amount

	column := []string{"amount", "updated_at"}

	if title != "" {
		column = append(column, "title")
	}

	_, err := rr.TX.NewUpdate().Model(replacement).Column(column...).Where("id = ?", id).Exec(ctx)

	return replacement, err
}

func (rr *ReplacementRepository) UpdateDone(ctx context.Context, done bool, id uuid.UUID) error {
	replacement := models.Replacement{Done: done}
	_, err := rr.TX.NewUpdate().Model(&replacement).Column("done", "updated_at").Where("id = ?", id).Exec(ctx)
	return err
}

func (rr *ReplacementRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := rr.TX.NewDelete().Model((*models.Replacement)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}

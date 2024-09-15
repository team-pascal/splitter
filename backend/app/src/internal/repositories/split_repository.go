package repositories

import (
	"context"
	"splitter/internal/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SplitRepository struct {
	TX *bun.Tx
}

func (sr *SplitRepository) FindByID(ctx context.Context, id string) (*models.Split, error) {
	split := new(models.Split)
	err := sr.TX.NewSelect().Model(split).Where("id = ?", id).Scan(ctx)
	return split, err
}

func (sr *SplitRepository) Create(ctx context.Context, title string, group_id string, amount uint) (*models.Split, error) {
	uuidGroupID, err := uuid.Parse(group_id)
	if err != nil {
		return nil, err
	}

	split := models.Split{Title: title, GroupID: uuidGroupID, Amount: amount}

	_, err = sr.TX.NewInsert().Model(&split).Exec(ctx)

	return &split, err
}

func (sr *SplitRepository) Update(ctx context.Context, title string, amount uint, id string) (*models.Split, error) {
	split := new(models.Split)
	split.Title = title
	split.Amount = amount

	column := []string{"amount", "updated_at"}

	if title != "" {
		column = append(column, "title")
	}

	_, err := sr.TX.NewUpdate().Model(split).Column(column...).Where("id = ?", id).Exec(ctx)

	return split, err
}

func (sr *SplitRepository) UpdateDone(ctx context.Context, done bool, id string) error {
	split := models.Split{Done: done}

	if _, err := sr.TX.NewUpdate().Model(&split).Column("done", "updated_at").Where("id = ?", id).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (sr *SplitRepository) Delete(ctx context.Context, splitID string) error {
	if _, err := sr.TX.NewDelete().Model((*models.Split)(nil)).Where("id = ?", splitID).Exec(ctx); err != nil {
		return err
	}
	return nil
}

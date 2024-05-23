package repositories

import (
	"context"
	"splitter/internal/models"

	"github.com/uptrace/bun"
)

type GroupRepository struct {
	DB *bun.DB
}

func (gr *GroupRepository) FindByID(ctx context.Context, id string) (*models.Group, error) {
	group := new(models.Group)
	err := gr.DB.NewSelect().Model(group).Where("id = ?", id).Scan(ctx)
	return group, err
}

package repositories

import (
	"context"
	"splitter/internal/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type GroupRepository struct {
	TX *bun.Tx
}

func (gr *GroupRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.Group, error) {
	group := new(models.Group)
	err := gr.TX.NewSelect().Model(group).Where("id = ?", id).Scan(ctx)
	return group, err
}

func (gr *GroupRepository) Create(ctx context.Context, name string) (*models.Group, error) {
	group := models.Group{Name: name}
	_, err := gr.TX.NewInsert().Model(&group).Exec(ctx)
	return &group, err
}

func (gr *GroupRepository) UpdateName(ctx context.Context, id uuid.UUID, name string) (*models.Group, error) {
	group := new(models.Group)
	group.Name = name

	_, err := gr.TX.NewUpdate().Model(group).Column("name", "updated_at").Where("id = ?", id).Exec(ctx)

	return group, err
}

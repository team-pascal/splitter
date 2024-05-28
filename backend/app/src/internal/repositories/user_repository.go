package repositories

import (
	"context"
	"splitter/internal/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	TX *bun.Tx
}

func (ur *UserRepository) FindByGroupID(ctx context.Context, group_id string) ([]models.User, error) {
	users := make([]models.User, 0)
	err := ur.TX.NewSelect().Model(&users).Where("group_id = ?", group_id).Scan(ctx)
	return users, err
}

func (ur *UserRepository) Create(ctx context.Context, names []string, groupID uuid.UUID) ([]models.User, error) {
	users := make([]models.User, len(names))

	for i, name := range names {
		users[i] = models.User{
			Name:    name,
			GroupID: groupID,
		}
	}
	_, err := ur.TX.NewInsert().Model(&users).Exec(ctx)
	return users, err
}

func (ur *UserRepository) UpdateName(ctx context.Context, name string, id string) (*models.User, error) {
	user := new(models.User)
	user.Name = name

	_, err := ur.TX.NewUpdate().Model(user).Column("name", "updated_at").Where("id = ?", id).Exec(ctx)

	return user, err
}

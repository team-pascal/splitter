package repositories

import (
	"context"
	"errors"
	"splitter/internal/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ReplacementLesseeRepository struct {
	TX *bun.Tx
}

func compareReplacementLessees(newReplacementLessees, oldReplacementLessees []models.ReplacementLessee) (createdLessees, updatedLessees, deletedLessees []models.ReplacementLessee) {
	newIDs := make(map[string]models.ReplacementLessee)
	oldIDs := make(map[string]models.ReplacementLessee)

	for _, newLessee := range newReplacementLessees {
		newIDs[newLessee.UserID.String()] = newLessee
	}

	for _, oldLessee := range oldReplacementLessees {
		oldIDs[oldLessee.UserID.String()] = oldLessee
	}

	for _, newLessee := range newReplacementLessees {
		if oldLessee, exists := oldIDs[newLessee.UserID.String()]; exists && newLessee.Amount != oldLessee.Amount {
			updatedLessees = append(updatedLessees, newLessee)
		} else if !exists {
			createdLessees = append(createdLessees, newLessee)
		}
	}

	for _, oldLessee := range oldReplacementLessees {
		if _, exists := newIDs[oldLessee.UserID.String()]; !exists {
			deletedLessees = append(deletedLessees, oldLessee)
		}
	}

	return
}

func (rlr *ReplacementLesseeRepository) FindByReplacementID(ctx context.Context, replacementID uuid.UUID) ([]models.ReplacementLessee, error) {
	replacementLessees := make([]models.ReplacementLessee, 0)
	err := rlr.TX.NewSelect().Model(&replacementLessees).Where("replacement_id = ?", replacementID).Order("amount DESC").Scan(ctx)
	return replacementLessees, err
}

func (rlr *ReplacementLesseeRepository) Create(ctx context.Context, lessees []models.ReplacementLessee) ([]models.ReplacementLessee, error) {
	_, err := rlr.TX.NewInsert().Model(&lessees).Exec(ctx)
	return lessees, err
}

func (rlr *ReplacementLesseeRepository) Update(ctx context.Context, newLessees []models.ReplacementLessee, replacementID uuid.UUID) error {
	// Process
	// 1. Get 'replacement_lessees' table where 'replacement_id' = replacementID
	// 2. If the pairs of 'replacement_id' and 'user_id' in request body exist in 'replacement_lessees' table, update the pairs in 'replacement_lessees' table.
	// 3. If the pairs of 'replacement_id' and 'user_id' in request body' do not exist in 'replacement_lessees' table, create the pairs in 'replacement_lessees' table.
	// 4. If the pairs of 'replacement_id' and 'user_id' in 'replacement_lessees' table do not exist in request body, delete the pairs in 'replacement_lessees' table.

	// 1. Get 'replacement_lessees' table where 'replacement_id' = replacementID
	oldLessees, err := rlr.FindByReplacementID(ctx, replacementID)
	if err != nil {
		return err
	}

	createdLessees, updatedLessees, deletedLessees := compareReplacementLessees(newLessees, oldLessees)

	// 2. If the pairs of 'replacement_id' and 'user_id' in request body exist in 'replacement_lessees' table, update the pairs in 'replacement_lessees' table.
	if len(createdLessees) > 0 {
		if _, err := rlr.Create(ctx, createdLessees); err != nil {
			createError := errors.New(err.Error() + " create lessees")
			return createError
		}
	}

	// 3. If the pairs of 'replacement_id' and 'user_id' in request body' do not exist in 'replacement_lessees' table, create the pairs in 'replacement_lessees' table.
	if len(updatedLessees) > 0 {
		if _, err := rlr.TX.NewUpdate().Model(&updatedLessees).Column("amount", "updated_at").Bulk().Exec(ctx); err != nil {
			updateError := errors.New(err.Error() + " update lessees")
			return updateError
		}
	}

	// 4. If the pairs of 'replacement_id' and 'user_id' in 'replacement_lessees' table do not exist in request body, delete the pairs in 'replacement_lessees' table.
	if len(deletedLessees) > 0 {
		if _, err := rlr.TX.NewDelete().Model(&deletedLessees).WherePK().Exec(ctx); err != nil {
			deleteError := errors.New(err.Error() + " delete lessees")
			return deleteError
		}
	}

	return nil
}

func (rlr *ReplacementLesseeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := rlr.TX.NewDelete().Model((*models.ReplacementLessee)(nil)).Where("replacement_id = ?", id).Exec(ctx)
	return err
}

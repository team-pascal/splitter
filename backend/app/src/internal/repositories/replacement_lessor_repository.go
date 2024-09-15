package repositories

import (
	"context"
	"errors"
	"splitter/internal/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ReplacementLessorRepository struct {
	TX *bun.Tx
}

func compareReplacementLessors(newReplacementLessors, oldReplacementLessors []models.ReplacementLessor) (createdLessors, deletedLessors, updatedLessors []models.ReplacementLessor) {
	newIDs := make(map[string]models.ReplacementLessor)
	oldIDs := make(map[string]models.ReplacementLessor)

	for _, newLessor := range newReplacementLessors {
		newIDs[newLessor.UserID.String()] = newLessor
	}

	for _, oldLessor := range oldReplacementLessors {
		oldIDs[oldLessor.UserID.String()] = oldLessor
	}

	for _, newLessor := range newReplacementLessors {
		if oldLessor, exists := oldIDs[newLessor.UserID.String()]; exists && newLessor.Amount != oldLessor.Amount {
			updatedLessors = append(updatedLessors, newLessor)
		} else if !exists {
			createdLessors = append(createdLessors, newLessor)
		}
	}

	for _, oldLessor := range oldReplacementLessors {
		if _, exists := newIDs[oldLessor.UserID.String()]; !exists {
			deletedLessors = append(deletedLessors, oldLessor)
		}
	}

	return
}

func (rlr *ReplacementLessorRepository) FindByReplacementID(ctx context.Context, replacementID uuid.UUID) ([]models.ReplacementLessor, error) {
	replacementLessors := make([]models.ReplacementLessor, 0)
	err := rlr.TX.NewSelect().Model(&replacementLessors).Where("replacement_id = ?", replacementID).Order("amount DESC").Scan(ctx)
	return replacementLessors, err
}

func (rlr *ReplacementLessorRepository) Create(ctx context.Context, lessors []models.ReplacementLessor) ([]models.ReplacementLessor, error) {
	_, err := rlr.TX.NewInsert().Model(&lessors).Exec(ctx)
	return lessors, err
}

func (rlr *ReplacementLessorRepository) Update(ctx context.Context, newLessors []models.ReplacementLessor, replacementID uuid.UUID) error {
	// Process
	// 1. Get 'replacement_lessors' table where 'replacement_id' = replacementID
	// 2. If the pairs of 'replacement_id' and 'user_id' in request body exist in 'replacement_lessors' table, update the pairs in 'replacement_lessors' table.
	// 3. If the pairs of 'replacement_id' and 'user_id' in request body' do not exist in 'replacement_lessors' table, create the pairs in 'replacement_lessors' table.
	// 4. If the pairs of 'replacement_id' and 'user_id' in 'replacement_lessors' table do not exist in request body, delete the pairs in 'replacement_lessors' table.

	// 1. Get 'replacement_lessors' table where 'replacement_id' = replacementID
	oldLessors := make([]models.ReplacementLessor, 0)
	if err := rlr.TX.NewSelect().Model(&oldLessors).Where("replacement_id = ?", replacementID).Scan(ctx); err != nil {
		return err
	}

	createdLessors, deletedLessors, updatedLessors := compareReplacementLessors(newLessors, oldLessors)

	// 2. If the pairs of 'replacement_id' and 'user_id' in request body exist in 'replacement_lessors' table, update the pairs in 'replacement_lessors' table.
	if len(createdLessors) > 0 {
		if _, err := rlr.Create(ctx, createdLessors); err != nil {
			createError := errors.New(err.Error() + " create lessors")
			return createError
		}
	}

	// 3. If the pairs of 'replacement_id' and 'user_id' in request body' do not exist in 'replacement_lessors' table, create the pairs in 'replacement_lessors' table.
	if len(updatedLessors) > 0 {
		if _, err := rlr.TX.NewUpdate().Model(&updatedLessors).Column("amount", "updated_at").Bulk().Exec(ctx); err != nil {
			updateError := errors.New(err.Error() + " update lessors")
			return updateError
		}
	}

	// 4. If the pairs of 'replacement_id' and 'user_id' in 'replacement_lessors' table do not exist in request body, delete the pairs in 'replacement_lessors' table.
	if len(deletedLessors) > 0 {
		if _, err := rlr.TX.NewDelete().Model(&deletedLessors).WherePK().Exec(ctx); err != nil {
			deleteError := errors.New(err.Error() + " delete lessors")
			return deleteError
		}
	}

	return nil
}

func (rlr *ReplacementLessorRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := rlr.TX.NewDelete().Model((*models.ReplacementLessor)(nil)).Where("replacement_id = ?", id).Exec(ctx)
	return err
}

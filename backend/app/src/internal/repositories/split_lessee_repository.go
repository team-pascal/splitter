package repositories

import (
	"context"
	"errors"
	"slices"
	"splitter/internal/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SplitLesseeRepository struct {
	TX *bun.Tx
}

func (slr *SplitLesseeRepository) FindBySplitID(ctx context.Context, split_id string) ([]models.SplitLessee, error) {
	splitLessee := make([]models.SplitLessee, 0)
	err := slr.TX.NewSelect().Model(&splitLessee).Where("split_id = ?", split_id).Scan(ctx)
	return splitLessee, err
}

func (slr *SplitLesseeRepository) Create(ctx context.Context, userIDs []string, splitID uuid.UUID) ([]models.SplitLessee, error) {
	lessees := make([]models.SplitLessee, len(userIDs))

	var err error

	for i, userID := range userIDs {
		uuidUserID, uuidErr := uuid.Parse(userID)
		if uuidErr != nil {
			err = uuidErr
			break
		}
		lessees[i] = models.SplitLessee{
			UserID:  uuidUserID,
			SplitID: splitID,
		}
	}

	if err != nil {
		return nil, err
	}

	_, err = slr.TX.NewInsert().Model(&lessees).Exec(ctx)

	return lessees, err
}

func (slr *SplitLesseeRepository) Update(ctx context.Context, lesseeIDs []string, splitID string) error {

	// 1. If the pairs of user_id in request body do not exist in lessees, create the data in lessees table.
	// 2. If the pairs of user_id in lessees do not exist in request body, delete the data in lessees table.

	// Initialization
	splitUUID, err := uuid.Parse(splitID)
	if err != nil {
		return err
	}

	lessees := make([]models.SplitLessee, 0)
	if err := slr.TX.NewSelect().Model(&lessees).Where("split_id = ?", splitUUID).Scan(ctx); err != nil {
		return err
	}

	oldLesseeIDs := make([]string, len(lessees))
	for i, oldLessee := range lessees {
		oldLesseeIDs[i] = oldLessee.UserID.String()
	}

	// 1. If the pairs of user_id in request body do not exist in lessees, create the data in lessees table.
	newLesseesIDs := make([]string, 0)

	for _, lesseeID := range lesseeIDs {
		if !slices.Contains(oldLesseeIDs, lesseeID) {
			newLesseesIDs = append(newLesseesIDs, lesseeID)
		}
	}

	if len(newLesseesIDs) > 0 {
		_, err = slr.Create(ctx, newLesseesIDs, splitUUID)
		if err != nil {
			return errors.New(err.Error() + " Create in Update")
		}
	}

	// 2. If the pairs of user_id in lessees do not exist in request body, delete the data in lessees table.
	deleteLesseeIDs := make([]uuid.UUID, 0)

	for _, oldLesseeID := range oldLesseeIDs {
		if !slices.Contains(lesseeIDs, oldLesseeID) {
			oldLesseeUUID, err := uuid.Parse(oldLesseeID)
			if err != nil {
				return err
			}
			deleteLesseeIDs = append(deleteLesseeIDs, oldLesseeUUID)
		}
	}

	if len(deleteLesseeIDs) > 0 {
		_, err := slr.TX.NewDelete().
			Model((*models.SplitLessee)(nil)).
			Where("split_id = ?", splitID).
			Where("user_id IN (?)", bun.In(deleteLesseeIDs)).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (slr *SplitLesseeRepository) Delete(ctx context.Context, splitID string) error {
	if _, err := slr.TX.NewDelete().Model((*models.SplitLessee)(nil)).Where("split_id = ?", splitID).Exec(ctx); err != nil {
		return err
	}
	return nil
}

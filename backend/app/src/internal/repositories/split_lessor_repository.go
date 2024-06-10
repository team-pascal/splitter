package repositories

import (
	"context"
	"errors"
	"splitter/internal/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SplitLessorRepository struct {
	TX *bun.Tx
}

type tempLessors struct {
	bun.BaseModel `bun:"table:temp_lessors"`

	SplitID uuid.UUID `bun:"split_id,type:uuid"`
	UserID  uuid.UUID `bun:"user_id,type:uuid"`
	Amount  uint      `bun:"amount,type:integer"`
}

func fromSplitLessor(sl models.SplitLessor) tempLessors {
	return tempLessors{
		SplitID: sl.SplitID,
		UserID:  sl.UserID,
		Amount:  sl.Amount,
	}
}

func (slr *SplitLessorRepository) FindBySplitID(ctx context.Context, split_id string) ([]models.SplitLessor, error) {
	splitLessors := make([]models.SplitLessor, 0)
	err := slr.TX.NewSelect().Model(&splitLessors).Where("split_id = ?", split_id).Order("amount ASC").Scan(ctx)
	return splitLessors, err
}

func (slr *SplitLessorRepository) Create(ctx context.Context, lessors []models.SplitLessor) ([]models.SplitLessor, error) {
	_, err := slr.TX.NewInsert().Model(&lessors).Exec(ctx)
	return lessors, err
}

func (slr *SplitLessorRepository) Update(ctx context.Context, lessors []models.SplitLessor) error {

	// 1. Create temporary table 'temp_lessors' from request data
	// 2. If the pairs of split_id and user_id in temp_lessors exist in lessors table, update the pairs in lessors table.
	// 3. If the pairs of split_id and user_id in temp_lessors do not exist in lessors table, create the pairs in lessors table.
	// 4. If the pairs of split_id and user_id in lessors table do not exist in temp_lessors, delete the pairs in lessors table.

	// Initialization
	defer slr.TX.NewDropTable().Model((*tempLessors)(nil)).IfExists().Exec(ctx)

	tmpl := make([]tempLessors, len(lessors))

	for i, lessor := range lessors {
		tmpl[i] = fromSplitLessor(lessor)
	}

	// 1. Create temporary table 'temp_lessors'

	// SQL
	// CREATE TEMPORARY TABLE temp_lessors ( split_id UUID, group UUID, amount INTEGER);

	if _, err := slr.TX.NewCreateTable().Temp().Model((*tempLessors)(nil)).Exec(ctx); err != nil {
		return err
	}
	if _, err := slr.TX.NewInsert().Model(&tmpl).Exec(ctx); err != nil {
		return err
	}

	// 2. If the pairs of split_id and user_id in temp_lessors exist in lessors table, update the pairs in lessors table.

	// SQL

	// UPDATE split_lessors
	// SET amount = temp_lessors.amount,
	// 	updated_at = NOW(),
	// FROM temp_lessors
	// WHERE split_lessors.split_id = temp_lessors.split_id
	// AND split_lessors.user_id = temp_lessors.user_id;

	_, err := slr.TX.NewRaw(`
			UPDATE split_lessors
			SET amount = temp_lessors.amount,
				updated_at = NOW()
			FROM temp_lessors
			WHERE split_lessors.split_id = temp_lessors.split_id
			AND split_lessors.user_id = temp_lessors.user_id;
		`).
		Exec(ctx)
	if err != nil {
		return errors.New(err.Error() + " Update")
	}

	// 3. If the pairs of split_id and user_id in temp_lessors do not exist in lessors table, create the pairs in lessors table.

	// SQL

	// INSERT INTO split_lessors (split_id, amount, user_id, created_at, updated_at)
	// SELECT temp_lessors.split_id, temp_lessors.amount, temp_lessors.user_id, NOW(), NOW()
	// FROM temp_new_data
	// LEFT JOIN split_lessors ON split_lessors.id = temp_lessors.id AND split_lessors.user_id = temp_lessors.user_id
	// WHERE split_lessors.split_id IS NULL;

	_, err = slr.TX.NewRaw(`
			INSERT INTO split_lessors (split_id, amount, user_id, created_at, updated_at)
			SELECT temp_lessors.split_id, temp_lessors.amount, temp_lessors.user_id, NOW(), NOW()
			FROM temp_lessors
			LEFT JOIN split_lessors ON split_lessors.split_id = temp_lessors.split_id AND split_lessors.user_id = temp_lessors.user_id
			WHERE split_lessors.split_id IS NULL;
		`).Exec(ctx)
	if err != nil {
		return errors.New(err.Error() + " INSERT")
	}

	// 4. If the pairs of split_id and user_id in lessors table do not exist in temp_lessors, delete the pairs in lessors table.

	// SQL

	// DELETE FROM split_lessors
	// WHERE NOT EXISTS (
	// 	SELECT 1
	// 	FROM temp_lessors
	// 	WHERE temp_lessors.user_id = split_lessors.user_id
	// 	AND temp_lessors.split_id = split_lessors.split_id
	// );

	_, err = slr.TX.NewRaw(`
		DELETE FROM split_lessors
		WHERE NOT EXISTS (
			SELECT 1
			FROM temp_lessors
			WHERE temp_lessors.user_id = split_lessors.user_id
			AND temp_lessors.split_id = split_lessors.split_id
		);
	`).Exec(ctx)
	if err != nil {
		return errors.New(err.Error() + " DELETE")
	}

	return nil
}

func (slr *SplitLessorRepository) Delete(ctx context.Context, splitID string) error {
	if _, err := slr.TX.NewDelete().Model((*models.SplitLessor)(nil)).Where("split_id = ?", splitID).Exec(ctx); err != nil {
		return nil
	}
	return nil
}

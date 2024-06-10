package repositories

import (
	"context"
	"splitter/internal/models"

	"github.com/uptrace/bun"
)

type PaymentRepository struct {
	TX *bun.Tx
}

func (pr *PaymentRepository) FindByGroupID(ctx context.Context, group_id string) ([]models.Payment, error) {

	// SQL

	// CREATE TEMPORARY TABLE payments AS
	// SELECT *, 'split' AS genre
	// FROM splits
	// WHERE group_id = group_id
	// UNION ALL
	// SELECT *, 'replacement' AS genre
	// FROM replacements
	// WHERE group_id = group_id;

	defer pr.TX.NewDropTable().Model((*models.Payment)(nil)).IfExists().Exec(ctx)

	selectSplitQuery := pr.TX.NewSelect().Model((*models.Split)(nil)).ColumnExpr("*, 'split' AS genre").Where("group_id = ?", group_id)
	selectRepacementQuery := pr.TX.NewSelect().Model((*models.Replacement)(nil)).ColumnExpr("*, 'replacement' AS genre").Where("group_id = ?", group_id)

	unionQuery := selectSplitQuery.Union(selectRepacementQuery)

	_, err := pr.TX.NewRaw(`CREATE TEMPORARY TABLE payments AS (?)`, unionQuery).Exec(ctx)
	if err != nil {
		return nil, err
	}

	payments := make([]models.Payment, 0)
	err = pr.TX.NewSelect().Model(&payments).Order("created_at ASC").Scan(ctx)

	return payments, err
}

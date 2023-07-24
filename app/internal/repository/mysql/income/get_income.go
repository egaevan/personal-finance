package income

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/personal-finance/app/domain/entity"
	"github.com/personal-finance/app/internal/repository/mapper"
	"github.com/personal-finance/app/internal/repository/models"
	log "github.com/sirupsen/logrus"
)

func (i *IncomeRepository) GetIncome(ctx context.Context, userId int) ([]*entity.Income, error) {
	query := fmt.Sprintf(`select id, user_id, income_information, total_income, created_at, updated_at from %s where user_id = $1`, models.Income{}.GetTableName())

	rows, err := i.db.QueryContext(ctx, query, userId)

	if err != nil {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("data not found")
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error(errRow)
		}
	}()

	resRow := make([]*models.Income, 0)
	for rows.Next() {
		t := &models.Income{}
		err = rows.Scan(
			&t.ID,
			&t.UserId,
			&t.IncomeInformation,
			&t.TotalIncome,
			&t.CreatedDate,
			&t.UpdatedDate,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		resRow = append(resRow, t)
	}

	result := mapper.ToArrEntityIncome(resRow)

	return result, nil
}

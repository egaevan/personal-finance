package income

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/personal-finance/app/domain/entity"
	"github.com/personal-finance/app/internal/repository/mapper"
	"github.com/personal-finance/app/internal/repository/models"
)

func (i *IncomeRepository) GetIncomeById(ctx context.Context, userId, id int) (*entity.Income, error) {
	query := fmt.Sprintf(`select id, user_id, income_information, total_income, created_at, updated_at from %s where user_id = $1 AND id = $2`, models.Income{}.GetTableName())

	income := &models.Income{}

	err := i.db.QueryRowContext(ctx, query, userId, id).Scan(
		&income.ID,
		&income.UserId,
		&income.IncomeInformation,
		&income.TotalIncome,
		&income.CreatedDate,
		&income.UpdatedDate,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("data not found")
	}

	if err != nil {
		return nil, err
	}

	result := mapper.ToEntityIncome(income)

	return result, nil
}

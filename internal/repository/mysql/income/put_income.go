package income

import (
	"context"
	"fmt"
	"github.com/personal-finance/domain/entity"
	"github.com/personal-finance/internal/repository/mapper"
	"github.com/personal-finance/internal/repository/models"
)

func (i *IncomeRepository) PutIncomeById(ctx context.Context, ett *entity.Income, userId, id int) error {
	req := mapper.ToModelIncome(ett)
	query := fmt.Sprintf(`UPDATE %s SET total_income = $1, income_information = $2 where user_id = $3 AND id = $4`, models.Income{}.GetTableName())

	_, err := i.db.ExecContext(ctx, query, req.TotalIncome, req.IncomeInformation, userId, id)

	if err != nil {
		return err
	}

	return nil
}

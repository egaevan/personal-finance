package income

import (
	"context"
	"fmt"
	"github.com/personal-finance/domain/entity"
	"github.com/personal-finance/internal/repository/mapper"
	"github.com/personal-finance/internal/repository/models"
)

func (i *IncomeRepository) InsertIncome(ctx context.Context, ett *entity.Income) error {
	req := mapper.ToModelIncome(ett)
	query := fmt.Sprintf(`INSERT INTO %s(user_id, total_income, income_information, created_at) VALUES ($1, $2, $3, NOW())`, models.Income{}.GetTableName())

	_, err := i.db.ExecContext(ctx, query, req.UserId, req.TotalIncome, req.IncomeInformation)

	if err != nil {
		return err
	}

	return nil
}

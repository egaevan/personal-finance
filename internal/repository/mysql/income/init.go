package income

import "database/sql"

type IncomeRepository struct {
	db *sql.DB
}

func NewIncomeRepository(db *sql.DB) IncomeRepository {
	return IncomeRepository{db: db}
}

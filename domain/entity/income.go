package entity

import "time"

type Income struct {
	ID                int       `json:"id"`
	UserId            int       `json:"user_id"`
	TotalIncome       int       `json:"total_income"`
	IncomeInformation string    `json:"income_information"`
	CreatedDate       time.Time `json:"created_at"`
	UpdatedDate       time.Time `json:"updated_at"`
}

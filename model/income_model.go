package model

import (
	"time"

	"github.com/andrewesteves/finfi/storage"
)

type IncomeModel struct {
	ID           int         `json:"id"`
	Title        string      `json:"title"`
	Client       ClientModel `json:"client"`
	Description  string      `json:"description"`
	Status       string      `json:"status"`
	Installments int         `json:"installments"`
	Total        float64     `json:"total"`
	ExpiredAt    time.Time   `json:"expired_at"`
	PaidAt       time.Time   `json:"paid_at"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

func (i IncomeModel) All() []IncomeModel {
	var incomes []IncomeModel
	db := storage.Connection()
	defer db.Close()

	rs, err := db.Query("SELECT id, client_id, title, description, status, installments, total, expired_at, paid_at, created_at, updated_at FROM incomes")
	if err != nil {
		panic(err.Error())
	}

	for rs.Next() {
		var income IncomeModel
		err = rs.Scan(&income.ID, &income.Client.ID, &income.Title, &income.Description, &income.Status, &income.Installments, &income.Total, &income.ExpiredAt, &income.PaidAt, &income.CreatedAt, &income.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		income.Client = ClientModel{}.Find(income.Client.ID)
		incomes = append(incomes, income)
	}

	return incomes
}

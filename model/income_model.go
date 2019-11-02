package model

import "time"

type IncomeModel struct {
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
	return []IncomeModel{
		IncomeModel{
			Title: "Item 01",
			Client: ClientModel{
				Name:  "Bill Gates",
				Email: "bill@microsoft.com",
			},
			Description:  "Lorem ipsum...",
			Status:       "Paid",
			Installments: 0,
			Total:        100.00,
		},
	}
}

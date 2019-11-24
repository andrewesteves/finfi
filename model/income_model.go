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

func (i IncomeModel) Find(id int) IncomeModel {
	var income IncomeModel
	db := storage.Connection()
	defer db.Close()
	err := db.QueryRow("SELECT id, client_id, title, description, status, installments, total, expired_at, paid_at, created_at, updated_at FROM incomes WHERE id = ?", id).Scan(&income.ID, &income.Client.ID, &income.Title, &income.Description, &income.Status, &income.Installments, &income.Total, &income.ExpiredAt, &income.PaidAt, &income.CreatedAt, &income.UpdatedAt)
	if err != nil {
		panic(err.Error())
	}
	income.Client = ClientModel{}.Find(income.Client.ID)
	return income
}

func (i IncomeModel) Insert() (IncomeModel, []Errors) {
	if hasErrors := incomeValidate(i); len(hasErrors) > 0 {
		return i, hasErrors
	}

	db := storage.Connection()
	defer db.Close()
	rs, err := db.Exec("INSERT INTO incomes (client_id, title, description, status, installments, total, expired_at, paid_at, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?,now())", i.Client.ID, i.Title, i.Description, i.Status, i.Installments, i.Total, i.ExpiredAt, i.PaidAt, i.CreatedAt)
	if err != nil {
		panic(err.Error())
	}
	id, err := rs.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	i.ID = int(id)
	return i, nil
}

func (i IncomeModel) Update(id int) (IncomeModel, []Errors) {
	if hasErrors := incomeValidate(i); len(hasErrors) > 0 {
		return i, hasErrors
	}

	db := storage.Connection()
	defer db.Close()
	rs, err := db.Prepare("UPDATE incomes SET title = ?, description = ?, status = ?, installments = ?, total = ?, expired_at = ?, paid_at = ?, updated_at = now() WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	rs.Exec(i.Title, i.Description, i.Status, i.Installments, i.Total, i.ExpiredAt, i.PaidAt, id)
	i.ID = id
	return i, nil
}

func (i IncomeModel) Destroy(id int) (IncomeModel, []Errors) {
	if id < 1 {
		var errs []Errors
		errs = append(errs, Errors{"id", "The id field is required"})
		return i, errs
	}

	db := storage.Connection()
	defer db.Close()
	rs, err := db.Prepare("DELETE FROM incomes WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	rs.Exec(id)
	i.ID = id
	return i, nil
}

func incomeValidate(i IncomeModel) []Errors {
	var errs []Errors
	if i.Title == "" {
		errs = append(errs, Errors{"title", "The title field is required"})
	}
	if i.Client.ID > 0 {
		errs = append(errs, Errors{"client_id", "The client_id field is required"})
	}
	if i.Total < 0 {
		errs = append(errs, Errors{"total", "The total field must be greater than 0"})
	}
	return errs
}

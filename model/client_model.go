package model

import "github.com/andrewesteves/finfi/storage"

type ClientModel struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}

func (c ClientModel) All() []ClientModel {
	var clients []ClientModel
	db := storage.Connection()
	defer db.Close()

	rs, err := db.Query("SELECT id, name, email, phone, description FROM clients")
	if err != nil {
		panic(err.Error())
	}

	for rs.Next() {
		var client ClientModel
		err = rs.Scan(&client.ID, &client.Name, &client.Email, &client.Phone, &client.Description)
		if err != nil {
			panic(err.Error())
		}
		clients = append(clients, client)
	}

	return clients
}

func (c ClientModel) Find(id int) ClientModel {
	var client ClientModel
	db := storage.Connection()
	defer db.Close()
	err := db.QueryRow("SELECT id, name, email, phone, description FROM clients WHERE id = ?", id).Scan(&client.ID, &client.Name, &client.Email, &client.Phone, &client.Description)
	if err != nil {
		panic(err.Error())
	}
	return client
}

func (c ClientModel) Insert() ClientModel {
	db := storage.Connection()
	defer db.Close()
	rs, err := db.Exec("INSERT INTO clients (name, email, phone, description) VALUES (?,?,?,?)", c.Name, c.Email, c.Phone, c.Description)
	if err != nil {
		panic(err.Error())
	}
	id, err := rs.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	c.ID = int(id)
	return c
}

func (c ClientModel) Update(id int) ClientModel {
	db := storage.Connection()
	defer db.Close()
	rs, err := db.Prepare("UPDATE clients SET name = ?, email = ?, phone = ?, description = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	rs.Exec(c.Name, c.Email, c.Phone, c.Description, id)
	c.ID = id
	return c
}

func (c ClientModel) Destroy(id int) ClientModel {
	db := storage.Connection()
	defer db.Close()
	rs, err := db.Prepare("DELETE FROM clients WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	rs.Exec(id)
	c.ID = id
	return c
}

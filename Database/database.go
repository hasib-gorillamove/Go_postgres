package database

import (
	"Go_postgres/model"
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	connectionString := "user=postgres password=hasib dbname=Go_postgres sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func CreateItem(item model.Item) (int, error) {
	query := "INSERT INTO item(name, description) VALUES($1, $2) RETURNING id"
	var id int
	err := db.QueryRow(query, item.Name, item.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func GetAllItems() ([]model.Item, error) {
	query := "SELECT id, name, description FROM item WHERE deleted_at IS NULL"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var items []model.Item
	defer rows.Close()
	for rows.Next() {
		var item model.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
func GetItem(id int) (model.Item, error) {
	var item model.Item
	query := "SELECT id, name, description FROM item WHERE id=$1 AND deleted_at IS NULL"
	err := db.QueryRow(query, id).Scan(&item.ID, &item.Name, &item.Description)
	if err != nil {
		return item, err
	}
	return item, nil
}

func UpdateItem(item model.Item) error {
	query := "UPDATE item SET name=$1, description=$2 WHERE id=$3"
	_, err := db.Exec(query, item.Name, item.Description, item.ID)
	return err
}

func DeleteItem(id int) error {
	query := "UPDATE item SET deleted_at= '2023-10-17 15:30:00' WHERE id=$1"
	_, err := db.Exec(query, id)
	return err
}

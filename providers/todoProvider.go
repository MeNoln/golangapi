package providers

import (
	"log"
	"time"

	"github.com/MeNoln/golangapi/db"
	"github.com/MeNoln/golangapi/models"
)

//GetAllTodos ...
func GetAllTodos() ([]models.Todo, error) {
	db := db.GetDb()
	defer db.Close()

	var todos []models.Todo
	rows, err := db.Queryx("select * from todo")
	if err != nil {
		log.Fatalln("Failed tp query")
		return nil, err
	}
	for rows.Next() {
		var todo models.Todo
		err := rows.StructScan(&todo)
		if err != nil {
			log.Fatalln("Failed")
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

//GetCurrentTodo ...
func GetCurrentTodo(id int) (models.Todo, error) {
	db := db.GetDb()
	defer db.Close()

	var todo models.Todo
	err := db.QueryRowx("select todo.id, todo.description from todo where id = $1", id).StructScan(&todo)
	if err != nil {
		log.Fatalln("Failed")
		return models.Todo{}, err
	}
	return todo, nil
}

//CreateTodo ...
func CreateTodo(data *models.Todo) error {
	data.DateCreated = time.Now()
	db := db.GetDb()
	defer db.Close()

	queryArgs := []interface{}{
		data.Description,
		data.DateCreated,
	}

	_, err := db.Exec("insert into todo (description, datecreated) values ($1, $2)", queryArgs...)
	if err != nil {
		log.Fatalln("Failed to insert")
		return err
	}
	return nil
}

//UpdateTodo ...
func UpdateTodo(data *models.Todo) error {
	data.DateCreated = time.Now()
	db := db.GetDb()
	defer db.Close()

	queryArgs := []interface{}{
		data.Description,
		data.DateCreated,
		data.ID,
	}

	_, err := db.Exec("update todo set description = $1, datecreated = $2 where id = $3", queryArgs...)
	if err != nil {
		log.Fatalln("Failed to update")
		return err
	}
	return nil
}

//DeleteTodo ...
func DeleteTodo(id int) error {
	db := db.GetDb()
	defer db.Close()

	_, err := db.Exec("delete from todo where id = $1", id)
	if err != nil {
		log.Fatalln("Failed to delete")
		return err
	}
	return nil
}

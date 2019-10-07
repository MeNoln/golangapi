package providers

import (
	"log"

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

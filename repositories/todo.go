package repositories

import (
	"github.com/jblim0125/golang-web-platform/infrastructures/datastore"
	"github.com/jblim0125/golang-web-platform/models"
)

// Todo is struct of todo.
type Todo struct {
	*datastore.DataStore
}

// New is constructor that creates TodoRepository
func (Todo) New(handler *datastore.DataStore) *Todo {
	return &Todo{handler}
}

// FindAll will return all recode of todo table.
func (repo *Todo) FindAll() (todos *models.Todos, err error) {
	result := repo.Orm.Find(&todos)
	if result.Error != nil {
	}
	if result.RowsAffected <= 0 {
	}
	return nil, nil
}

// FindByID will return todo whoes ID mathces
func (repo *Todo) FindByID(id string) (todos *models.Todos, err error) {
	result := repo.Orm.Where("ID = ?", id).Find(&todos)
	if result.Error != nil {
	}
	if result.RowsAffected <= 0 {
	}
	return nil, nil
}

// Store add a new todo recode.
func (repo *Todo) Store(todo models.Todo) (int, error) {
	//_, err := repo.DB.repo.Exec("insert into todo (ID, Title, Content ,Status) values ($1,$2,$3,$4)", todo.ID, todo.Title, todo.Content, todo.Status)
	// _, err := repo.Exec("insert into todo (ID, Title, Content ,Status) values (4,'todo4','todo-sample','0')")
	// if err != nil {
	// 	return 1, err
	// }
	return 0, nil
}

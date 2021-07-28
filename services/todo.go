package services

import (
	"github.com/jblim0125/golang-web-platform/models"
)

// Todo service - repository - interactor for Todo entity.
type Todo struct {
	Repo TodoRepository
}

// TodoRepository ...
type TodoRepository interface {
	FindAll() (*models.Todos, error)
	FindByID(string) (*models.Todos, error)
	Store(models.Todo) (int, error)
}

// New is constructor that creates Todo service
func (Todo) New(repo TodoRepository) *Todo {
	return &Todo{repo}
}

// GetAll returns All of todos.
func (service *Todo) GetAll() (*models.Todos, error) {
	return service.Repo.FindAll()
}

// GetByID returns todo whoes that ID mathces.
func (service *Todo) GetByID(id string) (*models.Todos, error) {
	return service.Repo.FindByID(id)
}

// Create create a new todo.
func (service *Todo) Create(todo models.Todo) (int, error) {
	return service.Repo.Store(todo)
}

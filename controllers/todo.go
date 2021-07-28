package controllers

import (
	"net/http"

	"github.com/jblim0125/golang-web-platform/models"
	"github.com/labstack/echo/v4"
)

// TodoUsecase Todo Usecase
type TodoUsecase interface {
	GetAll() (*models.Todos, error)
	GetByID(string) (*models.Todos, error)
	Create(models.Todo) (int, error)
}

// Todo Controller
type Todo struct {
	Usecase TodoUsecase
}

// New create Todo instance.
func (Todo) New(usecase TodoUsecase) *Todo {
	return &Todo{usecase}
}

// GetAll returns all of todos as JSON object.
func (controller *Todo) GetAll(c echo.Context) error {
	// res := model.HTTPResponse{}
	todos, err := controller.Usecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, todos)
	}
	return c.JSON(http.StatusOK, todos)
}

// GetByID return todos whoes ID mathces
func (controller *Todo) GetByID(c echo.Context) error {
	todo, err := controller.Usecase.GetByID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, todo)
}

// Create create a new todo.
func (controller *Todo) Create(c echo.Context) error {
	u := models.Todo{}
	c.Bind(&u)
	todo, err := controller.Usecase.Create(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, todo)
}

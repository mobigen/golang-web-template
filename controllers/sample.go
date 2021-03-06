package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/mobigen/golang-web-template/common"
	"github.com/mobigen/golang-web-template/models"
)

// SampleUsecase sample usecase(services)
type SampleUsecase interface {
	GetAll() (*[]models.Sample, error)
	GetByID(int) (*models.Sample, error)
	Create(*models.Sample) (*models.Sample, error)
	Update(*models.Sample) (*models.Sample, error)
	Delete(int) (*models.Sample, error)
}

// Sample Controller
type Sample struct {
	Log     *logrus.Logger
	Usecase SampleUsecase
}

// New create Sample instance.
func (Sample) New(usecase SampleUsecase) *Sample {
	return &Sample{
		Log:     common.Logger{}.GetInstance().Logger,
		Usecase: usecase,
	}
}

// GetAll returns all of sample as JSON object.
func (controller *Sample) GetAll(c echo.Context) error {
	samples, err := controller.Usecase.GetAll()
	if err != nil {
		if err == common.ErrNoHaveResult {
			return c.JSON(http.StatusOK, samples)
		}
		return c.JSON(http.StatusBadRequest, samples)
	}
	return c.JSON(http.StatusOK, samples)
}

// GetByID return sample whoes ID mathces
func (controller *Sample) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	sample, err := controller.Usecase.GetByID(id)
	if err != nil {
		if err == common.ErrNoHaveResult {
			return c.JSON(http.StatusOK, sample)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, sample)
}

// Create create a new ...
func (controller *Sample) Create(c echo.Context) error {
	input := new(models.Sample)
	c.Bind(input)
	sample, err := controller.Usecase.Create(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, sample)
}

// Update update from input
func (controller *Sample) Update(c echo.Context) error {
	input := new(models.Sample)
	c.Bind(input)
	sample, err := controller.Usecase.Update(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, sample)
}

// Delete delete sample from id
func (controller *Sample) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	sample, err := controller.Usecase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, sample)
}

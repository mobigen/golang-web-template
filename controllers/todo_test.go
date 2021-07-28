package controllers

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"github.com/jblim0125/golang-web-platform/common"
	"github.com/jblim0125/golang-web-platform/mocks"
	"github.com/jblim0125/golang-web-platform/models"
)

func TestGetByID(t *testing.T) {
	log := common.MakeTestLogger(t)
	mockUser := &models.Todo{
		ID:      1,
		Title:   "test",
		Content: "test",
		Status:  "test",
	}
	byteData, err := json.Marshal(mockUser)
	if err != nil {
		log.Error(err)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/todos/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)

	mockService := mocks.TodoUsecase{}
	mockService.On("GetByID", "1").Return(mockUser, nil)
	todoController := Todo{}.New(&mockService)
	// log.Debugf("Create Mock Service")

	if assert.NoError(t, todoController.GetByID(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(byteData), strings.Trim(rec.Body.String(), "\n"))
	}
}

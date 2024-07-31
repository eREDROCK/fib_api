package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"api/controllers"
)

func TestGetFib(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest("GET", "/fib?n=100", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.GetFib(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"result": 3736710778780434371}`, rec.Body.String())
	}
}

func TestGetFibInvalidInput_minus(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest("GET", "/fib?n=-6", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.GetFib(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"message":"Invalid request", "status":"400"}`, rec.Body.String())
	}
}

func TestGetFibInvalidInput_string(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest("GET", "/fib?n=数字", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.GetFib(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"message":"Invalid request", "status":"400"}`, rec.Body.String())
	}
}

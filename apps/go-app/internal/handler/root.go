package handler

import (
	"go-app/internal/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// RootHandler returns a "Helo World" message
func RootHandler(c echo.Context) error {
	response := model.SuccessResponse[map[string]string]{
		BaseAPIResponse: model.BaseAPIResponse{
			Status:  http.StatusOK,
			Success: true,
			Message: "Helo World",
		},
		Data: map[string]string{},
	}
	return c.JSON(http.StatusOK, response)
}

// TODO: use health check package like: https://github.com/alexliesenfeld/health
func HealthCheckHandler(c echo.Context) error {
	response := model.HealthCheckResponse{
		BaseAPIResponse: model.BaseAPIResponse{
			Status:  http.StatusOK,
			Success: true,
			Message: "All is well",
		},
		Data: model.HealthCheckData{
			Uptime:    0000000000,
			Timestamp: time.Now().Unix(),
		},
	}
	return c.JSON(http.StatusOK, response)
}

package http

import (
	"net/http"

	"github.com/github-profile/go-boilerplate/domain/health"
	SharedContext "github.com/github-profile/go-boilerplate/shared/context"
	"github.com/labstack/echo"
)

type handlerHealthCheck struct {
	usecase health.Usecase
}

func HealthCheckHandler(e *echo.Echo, usecase health.Usecase) {
	handler := handlerHealthCheck{
		usecase: usecase,
	}

	e.GET("/api/health/check", handler.DoHealthCheck)
}

func (h handlerHealthCheck) DoHealthCheck(c echo.Context) error {
	ac := c.(*SharedContext.ApplicationContext)
	res, err := h.usecase.HealthCheck(c)
	if err != nil {
		return err
	}

	return ac.WithSuccess("100% Healthy", http.StatusOK, res)
}

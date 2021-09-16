package health

import (
	"github.com/labstack/echo"
)

type Usecase interface {
	HealthCheck(c echo.Context) (interface{}, error)
}

type Repository interface {
	MySqlHealthCheck() (bool, error)
}

package context

import (
	"github.com/labstack/echo"

	SharedConfig "github.com/github-profile/go-boilerplate/shared/config"
)

type ApplicationContext struct {
	echo.Context
	Config SharedConfig.Config
}

func (c *ApplicationContext) WithSuccess(message string, status int, data interface{}) error {
	return c.JSON(status, &SuccessResponse{
		Message: message,
		Data:    data,
		Status:  status,
	})
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
}

type ErrorResponse struct {
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Status    int         `json:"status"`
	ErrorCode string      `json:"error_code"`
}

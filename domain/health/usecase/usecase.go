package usecase

import (
	"github.com/github-profile/go-boilerplate/domain/health"
	"github.com/labstack/echo"
)

type usecase struct {
	repository health.Repository
}

func NewHealthCheckUsecase(repository health.Repository) health.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u usecase) HealthCheck(c echo.Context) (interface{}, error) {
	var err error
	// - mysql HealthCheck
	_, err = u.repository.MySqlHealthCheck()
	if err != nil {
		return false, err
	}

	return true, nil
}

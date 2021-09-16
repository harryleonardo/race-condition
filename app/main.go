package main

import (
	// - common
	"fmt"

	"github.com/labstack/echo"

	// - shared
	SharedConfig "github.com/github-profile/go-boilerplate/shared/config"
	SharedContainer "github.com/github-profile/go-boilerplate/shared/container"
	SharedContext "github.com/github-profile/go-boilerplate/shared/context"
	SharedDatabase "github.com/github-profile/go-boilerplate/shared/database"

	// - repository
	HealthRepository "github.com/github-profile/go-boilerplate/domain/health/repository"

	// - usecase
	HealthUsecase "github.com/github-profile/go-boilerplate/domain/health/usecase"

	// - delivery
	HealthHandler "github.com/github-profile/go-boilerplate/domain/health/delivery/http"
)

func main() {
	// - initialize echo labstack as a framework
	e := echo.New()

	// - get default dependency injection container
	container := SharedContainer.GetDefaultContainer()
	getConfig := container.MustGet("shared.config").(*SharedConfig.Config)

	// - initialize config
	conf := *getConfig
	mysql := container.MustGet("shared.database").(SharedDatabase.MysqlInterface)

	mysqlSess, err := mysql.OpenMysqlConn()
	if err != nil {
		// panic(msgError)
	}

	// - initialize customize context;
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ac := &SharedContext.ApplicationContext{
				Context: c,
				Config:  conf,
			}
			return h(ac)
		}
	})

	// - repository
	healthCheckRepo := HealthRepository.NewHealthCheckRepository(mysqlSess)

	// - usecase
	healthCheckUcase := HealthUsecase.NewHealthCheckUsecase(healthCheckRepo)

	// - delivery
	HealthHandler.HealthCheckHandler(e, healthCheckUcase)

	e.Start(fmt.Sprintf(":%d", conf.Port))
}

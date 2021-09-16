package container

import (
	"github.com/fgrosse/goldi"
	SharedConfig "github.com/github-profile/go-boilerplate/shared/config"
	SharedDatabase "github.com/github-profile/go-boilerplate/shared/database"
)

func GetDefaultContainer() *goldi.Container {
	registry := goldi.NewTypeRegistry()
	config := make(map[string]interface{})
	container := goldi.NewContainer(registry, config)

	container.RegisterType("shared.config", SharedConfig.GetDefaultImmutableConfig)
	container.RegisterType("shared.database", SharedDatabase.NewMysql, "@shared.config")

	return container
}

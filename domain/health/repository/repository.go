package repository

import (
	"github.com/github-profile/go-boilerplate/domain/health"
	"github.com/jinzhu/gorm"
)

type repoHandler struct {
	mySqlSess *gorm.DB
}

func NewHealthCheckRepository(mySqlSess *gorm.DB) health.Repository {
	return &repoHandler{
		mySqlSess: mySqlSess,
	}
}

func (r repoHandler) MySqlHealthCheck() (bool, error) {
	// err := r.mySqlSess.DB().Ping()
	// if err != nil {
	// 	return false, err
	// }

	return true, nil
}

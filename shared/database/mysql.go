package database

import (
	"fmt"
	"sync"
	"time"

	"github.com/github-profile/go-boilerplate/shared/config"
	"github.com/jinzhu/gorm"
)

var (
	once sync.Once
)

type (
	// MysqlInterface ...
	MysqlInterface interface {
		OpenMysqlConn() (*gorm.DB, error)
	}

	database struct {
		SharedConfig config.Config
	}
)

func (d *database) OpenMysqlConn() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		d.SharedConfig.DATABASE.Username,
		d.SharedConfig.DATABASE.Password,
		d.SharedConfig.DATABASE.Host,
		d.SharedConfig.DATABASE.Name)

	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	db.BlockGlobalUpdate(true) //to prevent update/delete without where clause
	db.DB().SetMaxIdleConns(64)
	db.DB().SetMaxIdleConns(128)
	db.DB().SetConnMaxLifetime(2 * time.Second)
	db.LogMode(true)
	return db, nil
}

// NewMysql is an factory that implement of mysql database configuration
func NewMysql(config *config.Config) MysqlInterface {
	if config == nil {
		panic("[CONFIG] immutable config is required")
	}

	return &database{*config}
}

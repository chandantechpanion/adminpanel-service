package storage

import (
	"errors"
	"fmt"

	"github.com/chandantechpanion/adminpanel-service/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func NewDb(c *config.Config) (*gorm.DB, error) {
	if c.DB.Use == "mssql" {
		return newMSSQL(c)
	}
	return nil, errors.New("not supported db")
}

func newMSSQL(c *config.Config) (*gorm.DB, error) {
	mssqlInfo := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		c.DB.Mssql.UserName,
		c.DB.Mssql.Password,
		c.DB.Mssql.Host,
		c.DB.Mssql.Port,
		c.DB.Mssql.Database)

	db, err := gorm.Open("mssql", mssqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

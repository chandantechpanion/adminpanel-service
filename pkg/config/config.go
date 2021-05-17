package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

type Config struct {
	AppName string `default:"adminpanle"`
	Port    string `default:"8283"`
	Logger  struct {
		Use         string `default:"zapLogger"`
		Environment string `default:"prod"`
		LogLevel    string `default:"info"`
		FileName    string `default:"adminpanle.log"`
	}
	DB struct {
		Use   string `default:"mssql"`
		Mssql struct {
			Enabled  bool   `default:"true"`
			Host     string `default:"localhost"`
			Port     string `default:"1433"`
			UserName string `default:"sa"`
			Password string `default:"sa"`
			Database string `default:"Nimble_SP_NewDevelopment"`
		}
	}
	Contacts struct {
		Name  string `default:"Akbar Shaikh"`
		Email string `default:"csharma@techpanion.com"`
	}
	Laserfiche struct {
		Applicationurl string `default:"localhost"`
	}
}

func NewConfig() (*Config, error) {
	c := &Config{}
	err := configor.Load(c, "pkg/config/config.yml")
	if err != nil {
		fmt.Println("Config Not loaded")
		fmt.Println(err)
		return nil, err
	}
	return c, nil
}

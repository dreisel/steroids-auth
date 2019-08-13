package db

import (
	"fmt"
	"github.com/dreisel/steroids-auth/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/*
	connect to a pg and returns a pg instance
*/

type Options struct {
	User     string
	Host     string
	Port     string
	Password string
	Database string
	SSLMode  string
}

var dbOptions = Options{
	User:     utils.GetEnv("POSTGRES_USER", "postgres"),
	Host:     utils.GetEnv("POSTGRES_HOST", "localhost"),
	Port:     utils.GetEnv("POSTGRES_PORT", "5432"),
	Password: utils.GetEnv("POSTGRES_PASSWORD", ""),
	Database: utils.GetEnv("POSTGRES_DATABASE", "postgres"),
	SSLMode:  utils.GetEnv("POSTGRES_SSL", "disable"),
}

func (o *Options) String() string {

	return fmt.Sprintf("sslmode=%s host=%s port=%s user=%s dbname=%s password=%s", o.SSLMode, o.Host, o.Port, o.User, o.Database, o.Password)
}

func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", dbOptions.String())
	if err != nil {
		panic(err.Error())
	}
	return db
}

package config

import (
	"fmt"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type IConfig interface {
	Database() IDatabase
}

type config struct {
	db *database
}

func (c *config) Database() IDatabase {
	return c.db
}

func LoadConfig() IConfig {
	envMap, err := godotenv.Read(".env")
	if err != nil {
		logrus.Error("Error loading .env file", err.Error())
	}

	return &config{
		db: &database{
			host: envMap["DB_HOST"],
			port: func() int {
				port, err := strconv.Atoi(envMap["DB_PORT"])
				if err != nil {
					logrus.Error("Error converting port to int", err.Error())
					return 0
				}
				return port
			}(),
			username: envMap["DB_USERNAME"],
			password: envMap["DB_PASSWORD"],
			sslmode:  "disable",
			dbname:   envMap["DB_NAME"],
		},
	}
}

type IDatabase interface {
	EndPoint() string
	Host() string
	Port() int
	SSLMode() string
	UserName() string
	PassWord() string
	DatabaseName() string
}

type database struct {
	host     string
	port     int
	username string
	password string
	sslmode  string
	dbname   string
}

func (c *database) EndPoint() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.username, c.password, c.host, c.port, c.dbname, c.sslmode)
}

func (c *database) Host() string {
	return c.host
}

func (c *database) Port() int {
	return c.port
}

func (c *database) SSLMode() string {
	return c.sslmode
}

func (c *database) UserName() string {
	return c.username
}

func (c *database) PassWord() string {
	return c.password
}

func (c *database) DatabaseName() string {
	return c.dbname
}

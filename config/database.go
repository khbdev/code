package config

import (
	"database/sql"
	"fmt"
)




type Postgress struct {
	Host string
	Port int64
	User string
	Password string
	DBname  string
	SSLMode  string
}

func ConfigPostgress(cfg Postgress) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s")
}
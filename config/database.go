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
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBname, cfg.SSLMode)


	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return  nil, err
	}
	if err := db.Ping(); err != nil {
		return  nil, err
	}
	return db, nil
}
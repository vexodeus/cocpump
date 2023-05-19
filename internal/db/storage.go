package db

import (
	"database/sql"
	"fmt"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

type Storage struct {
	DB     *sql.DB
	Config Config
}
type Config struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
}

func NewStorage() (*Storage, error) {
	config := Config{}
	toml.DecodeFile("db.toml", &config)
	db, err := getDB(config)
	if err != nil {
		return nil, err
	}
	return &Storage{
		DB:     db,
		Config: config,
	}, nil
}
func getDB(config Config) (*sql.DB, error) {
	DBURI := getURI(config)
	db, err := sql.Open("postgres", DBURI)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func getURI(config Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Name,
	)
}

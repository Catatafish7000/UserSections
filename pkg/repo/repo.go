package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	Host     = "localhost"
	Port     = 5432
	Username = "kirill"
	Password = "avito"
	DBName   = "db"
)

type repo struct {
	DB *sqlx.DB
}

func NewRepo() *repo {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, Port, Username, Password, DBName)
	db, err := sqlx.Open("postgres", psqlConn)
	if err != nil {
		log.Fatal("Failed to connect to db")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Failed to ping db")
	}
	return &repo{DB: db}
}

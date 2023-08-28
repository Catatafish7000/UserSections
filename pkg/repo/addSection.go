package repo

import (
	"errors"
	"log"
	"strings"
	"time"
)

func (r *repo) AddSection(userID int, sectionName string) error {
	var sect string
	if err := r.DB.QueryRow("select section from list where section=$1", sectionName).Scan(&sect); err != nil {
		return errors.New("No such section created")
	}
	_, err := r.DB.Exec("insert into sections (user_id,section,created_at) values ($1,$2,$3)", userID, sectionName, time.Now())
	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		_, err1 := r.DB.Exec("update sections set created_at=$1 where user_id=$2 and section=$3", time.Now(), userID, sectionName)
		if err1 != nil {
			log.Println(err1.Error())
		}
	}
	return err
}

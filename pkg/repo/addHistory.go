package repo

import (
	"log"
	"strings"
	"time"
)

func (r *repo) AddHistory(userID int, sectionName string) error {
	_, err := r.DB.Exec("insert into history (user_id,section,created_at) values ($1,$2,$3)", userID, sectionName, time.Now())
	if err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		_, err1 := r.DB.Exec("update history set created_at=$1 where user_id=$2 and section=$3", time.Now(), userID, sectionName)
		if err1 != nil {
			log.Println(err1.Error())
			return err1
		}
	}
	return err
}

package repo

import (
	"log"
	"time"
)

func (r *repo) DeleteHistory(userID int, sectionName string) error {
	_, err := r.DB.Exec("update history set deleted_at=$1 where user_id=$2 and section=$3", time.Now(), userID, sectionName)
	if err != nil {
		log.Println(err.Error())
	}
	return err
}

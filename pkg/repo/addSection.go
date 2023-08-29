package repo

import (
	"errors"
	"log"
	"strings"
	"time"
)

func (r *repo) AddSection(userID int, sectionName string, ttl int) error {
	var sect string
	if err := r.DB.QueryRow("select section from list where section=$1", sectionName).Scan(&sect); err != nil {
		return errors.New("no such section created")
	}
	_, err := r.DB.Exec("insert into sections (user_id,section,created_at, ttl) values ($1,$2,$3,$4)", userID, sectionName, time.Now(), time.Now().Add(time.Hour*time.Duration(ttl*24)))
	if err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		_, err1 := r.DB.Exec("update sections set created_at=$1, ttl=$2 where user_id=$3 and section=$4", time.Now(), time.Now().Add(time.Hour*time.Duration(ttl*24)), userID, sectionName)
		if err1 != nil {
			log.Println(err1.Error())
			return err1
		}
	}
	return err
}

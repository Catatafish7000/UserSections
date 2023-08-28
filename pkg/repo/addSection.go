package repo

import (
	"errors"
	"time"
)

func (r *repo) AddSection(userID int, sectionName string) error {
	var sect string
	if err := r.DB.QueryRow("select section from list where section=$1", sectionName).Scan(&sect); err != nil {
		return errors.New("No such section created")
	}
	_, err := r.DB.Exec("insert into sections (user_id,section,created_at) values ($1,$2,$3)", userID, sectionName, time.Now())
	return err
}

package repo

import "errors"

func (r *repo) CheckExistence(userID int, sectionName string) error {
	rows, err := r.DB.Query("select * from sections where user_id=$1 and section=$2;", userID, sectionName)
	if rows.Next() == false {
		return errors.New("no assignment")
	}
	return err
}

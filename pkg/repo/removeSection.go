package repo

import "errors"

func (r *repo) RemoveSection(userID int, sectionName string) error {
	var sect string
	if err := r.DB.QueryRow("select section from list where section=$1", sectionName).Scan(&sect); err != nil {
		return errors.New("No such section created")
	}
	_, err := r.DB.Exec("delete from sections where user_id=$1 and section=$2", userID, sectionName)
	return err
}

package repo

func (r *repo) CheckExistence(userID int, sectionName string) error {
	_, err := r.DB.Query("select * from sections where user_id=$1 ans section=$2;", userID, sectionName)
	return err
}

package repo

func (r *repo) DeleteSection(sectionName string) error {
	_, err := r.DB.Exec("DELETE from list where section=$1", sectionName)
	r.DB.Exec("delete from sections where section = $1", sectionName)
	return err
}

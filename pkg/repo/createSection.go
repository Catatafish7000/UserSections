package repo

func (r *repo) CreateSection(sectionName string) error {
	_, err := r.DB.Exec("INSERT INTO list (section) VALUES ($1)", sectionName)
	return err
}

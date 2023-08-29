package repo

func (r *repo) CreateSection(sectionName string, percentage int) error {
	_, err := r.DB.Exec("INSERT INTO list (section, percentage) VALUES ($1,$2)", sectionName, percentage)
	return err
}

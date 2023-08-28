package handlers

type Repo interface {
	CreateSection(sectionName string) error
	DeleteSection(sectionName string) error
	AddSection(userID int, sectionName string) error
	RemoveSection(userID int, sectionName string) error
	ShowSections(userID int) ([]byte, error)
}

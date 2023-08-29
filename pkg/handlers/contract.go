package handlers

type Repo interface {
	CreateSection(sectionName string, percentage int) error
	DeleteSection(sectionName string) error
	AddSection(userID int, sectionName string, ttl int) error
	RemoveSection(userID int, sectionName string) error
	ShowSections(userID int) ([]byte, error)
	Clear()
	GetSectionList() ([]byte, error)
	ShowHistory(year int, month int) ([]byte, error)
	AddHistory(userID int, sectionName string) error
	DeleteHistory(userID int, sectionName string) error
}

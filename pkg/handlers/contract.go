//go:generate mockgen -source ${GOFILE} -destination mocks_test.go -package ${GOPACKAGE}_test
package handlers

type repo interface {
	CreateSection(sectionName string, percentage int) error
	DeleteSection(sectionName string) error
	AddSection(userID int, sectionName string, ttl int) error
	RemoveSection(userID int, sectionName string) error
	ShowSections(userID int) ([]byte, error)
	GetSectionList() ([]byte, error)
	ShowHistory(year int, month int) ([]byte, error)
	AddHistory(userID int, sectionName string) error
	DeleteHistory(userID int, sectionName string) error
	CheckExistence(userID int, sectionName string) error
}

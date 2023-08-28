package repo

import (
	"encoding/json"
	"log"
)

type SectionList struct {
	UserID   int
	Sections []string
}

func (r *repo) ShowSections(userID int) ([]byte, error) {
	sects := make([]string, 0)
	rows, err := r.DB.Query("select section from sections where user_id=$1", userID)
	if err != nil {
		log.Println(err.Error())
	}
	for rows.Next() {
		var sect string
		if err := rows.Scan(&sect); err != nil {
			log.Println(err.Error())
		}
		sects = append(sects, sect)
	}
	if err1 := rows.Err(); err1 != nil {
		log.Println(err1.Error())
	}
	list := SectionList{userID, sects}
	ans, err2 := json.Marshal(list)
	if err2 != nil {
		log.Fatal(err2)
	}
	return ans, err
}

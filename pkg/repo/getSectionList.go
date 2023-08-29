package repo

import (
	"encoding/json"
	"log"
)

func (r *repo) GetSectionList() ([]byte, error) {
	rows, err := r.DB.Query("select * from list")
	if err != nil {
		return nil, err
	}
	ans := make(map[string]string, 0)
	for rows.Next() {
		var sect, percentage string
		if err := rows.Scan(&sect, &percentage); err != nil {
			log.Println(err.Error())
		}
		ans[sect] = percentage
	}
	result, err1 := json.Marshal(ans)
	return result, err1
}

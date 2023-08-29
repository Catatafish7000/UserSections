package repo

import (
	"encoding/json"
	"strings"
)

type Info struct {
	CreatedAt string
	DeletedAt string
}

func (r *repo) ShowHistory(year int, month int) ([]byte, error) {
	resp := make(map[string]Info, 0)
	rows, err := r.DB.Query("select * from history where date_part ('year', created_at)=$1 and date_part('month',created_at)=$2 or date_part('year',deleted_at)=$3 and date_part('month',deleted_at)=$4", year, month, year, month)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id, section, createdAt, deletedAt string
		err := rows.Scan(&id, &section, &createdAt, &deletedAt)
		if err != nil && strings.Contains(err.Error(), "Scan error on column index 3, name \"deleted_at\": converting NULL to string") {
			deletedAt = "not deleted yet"
		} else if err != nil {
			return nil, err
		}
		resp[id+" assignment to "+section] = Info{createdAt, deletedAt}
	}
	ans, err1 := json.Marshal(resp)
	return ans, err1
}

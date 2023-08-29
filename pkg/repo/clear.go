package repo

import (
	"log"
	"strconv"
)

func (r *repo) Clear() {
	rows, err1 := r.DB.Query("select (user_id,section) from sections where ttl>CURRENT_TIMESTAMP")
	if err1 != nil {
		log.Println(err1.Error())
	}
	for rows.Next() {
		var id, sect string
		rows.Scan(&id, &sect)
		userID, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err.Error())
		}
		err = r.DeleteHistory(userID, sect)
		if err != nil {
			log.Println(err.Error())
		}
	}
	_, err := r.DB.Exec("delete from sections where ttl>CURRENT_TIMESTAMP;")
	if err != nil {
		log.Println("Failed to clear memory. Error:", err.Error())
	} else {
		log.Println("Memory has been successfully cleared")
	}
}

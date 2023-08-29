package repo

import "log"

func (r *repo) Clear() {
	_, err := r.DB.Exec("delete from sections where ttl>CURRENT_TIMESTAMP;")
	if err != nil {
		log.Println("Failed to clear memory. Error:", err.Error())
	} else {
		log.Println("Memory has been successfully cleared")
	}
}

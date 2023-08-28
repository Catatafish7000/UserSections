package handlers

import (
	"encoding/json"
	"net/http"
)

func jsonError(w http.ResponseWriter, msg string, status int) {
	resp, _ := json.Marshal(map[string]string{
		"message": msg,
	})
	w.WriteHeader(status)
	w.Write(resp)
}

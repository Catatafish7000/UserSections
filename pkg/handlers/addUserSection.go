package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) AddUserSection(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user"]
	sectionName := mux.Vars(r)["section"]
	ttlS := mux.Vars(r)["ttl"]
	ttl, _ := strconv.Atoi(ttlS)
	id, _ := strconv.Atoi(userID)
	err := h.Repo.AddSection(id, sectionName, ttl)
	var resp []byte
	if err != nil && !strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		jsonError(w, err.Error(), http.StatusInternalServerError)
	} else if err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		resp, _ = json.Marshal(map[string]string{
			"userID":  userID,
			"section": sectionName,
			"msg":     "time of section assignment has been updated",
		})
	} else {
		resp, _ = json.Marshal(map[string]string{
			"userID":  userID,
			"section": sectionName,
			"msg":     "section has been assigned to user",
		})
	}
	w.Write(resp)
}

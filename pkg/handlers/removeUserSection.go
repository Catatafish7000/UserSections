package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) RemoveUserSection(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user"]
	sectionName := mux.Vars(r)["section"]
	id, _ := strconv.Atoi(userID)
	err := h.Repo.RemoveSection(id, sectionName)
	var resp []byte
	if err != nil {
		jsonError(w, err.Error(), http.StatusInternalServerError)
	} else {
		resp, _ = json.Marshal(sectionName)
	}
	w.Write(resp)
}

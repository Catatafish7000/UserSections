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
	id, errId := strconv.Atoi(userID)
	if errId != nil {
		JsonError(w, errId.Error(), 400)
		return
	}
	errExist := h.repo.CheckExistence(id, sectionName)
	if errExist != nil {
		JsonError(w, "no such assignment", 400)
		return
	}
	err := h.repo.RemoveSection(id, sectionName)
	var resp []byte
	if err != nil {
		JsonError(w, err.Error(), http.StatusInternalServerError)
	} else {
		resp, _ = json.Marshal(map[string]string{
			"userID":  userID,
			"section": sectionName,
			"msg":     "section has been unassigned from user",
		})
	}
	err = h.repo.DeleteHistory(id, sectionName)
	if err != nil {
		JsonError(w, err.Error()+" failed to update history", http.StatusInternalServerError)
	}
	w.Write(resp)
}

package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	sectionName := mux.Vars(r)["section"]
	err := h.Repo.DeleteSection(sectionName)
	var resp []byte
	if err != nil {
		jsonError(w, err.Error(), http.StatusInternalServerError)
	} else {
		resp, _ = json.Marshal(sectionName)
	}
	w.Write(resp)
}

package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) NewSection(w http.ResponseWriter, r *http.Request) {
	sectionName := mux.Vars(r)["section"]
	err := h.Repo.CreateSection(sectionName)
	var resp []byte
	if err != nil {
		jsonError(w, err.Error(), http.StatusBadRequest)
	} else {
		resp, _ = json.Marshal(sectionName)
	}
	w.Write(resp)
}

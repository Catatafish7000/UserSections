package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) ShowUserSection(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user"]
	id, _ := strconv.Atoi(userID)
	resp, err := h.Repo.ShowSections(id)
	if err != nil {
		jsonError(w, err.Error(), 0)
	}
	w.Write(resp)
}

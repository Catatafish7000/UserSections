package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) ShowUserSection(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user"]
	id, errId := strconv.Atoi(userID)
	if errId != nil {
		jsonError(w, errId.Error(), 400)
		return
	}
	resp, err := h.repo.ShowSections(id)
	if err != nil {
		jsonError(w, err.Error(), 0)
	}
	w.Write(resp)
}

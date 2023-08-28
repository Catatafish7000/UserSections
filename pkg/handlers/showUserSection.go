package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) ShowUserSection(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user"]
	id, _ := strconv.Atoi(userID)
	resp, err := h.Repo.ShowSections(id)
	if err != nil {
		log.Println(err.Error())
	}
	w.Write(resp)
}

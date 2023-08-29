package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) History(w http.ResponseWriter, r *http.Request) {
	year, _ := strconv.Atoi(mux.Vars(r)["year"])
	month, _ := strconv.Atoi(mux.Vars(r)["month"])
	resp, err := h.Repo.ShowHistory(year, month)
	if err != nil {
		jsonError(w, err.Error(), 500)
	}
	w.Write(resp)
}

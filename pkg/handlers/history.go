package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) History(w http.ResponseWriter, r *http.Request) {
	year, errYear := strconv.Atoi(mux.Vars(r)["year"])
	if errYear != nil {
		JsonError(w, errYear.Error(), 400)
		return
	}
	month, errMonth := strconv.Atoi(mux.Vars(r)["month"])
	if errMonth != nil {
		JsonError(w, errMonth.Error(), 400)
		return
	}
	resp, err := h.repo.ShowHistory(year, month)
	if err != nil {
		JsonError(w, err.Error(), 500)
	}
	w.Write(resp)
}

package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) NewSection(w http.ResponseWriter, r *http.Request) {
	sectionName := mux.Vars(r)["section"]
	percentage, err1 := strconv.Atoi(mux.Vars(r)["percentage"])
	if err1 != nil {
		jsonError(w, err1.Error(), http.StatusBadRequest)
		return
	}
	err := h.Repo.CreateSection(sectionName, percentage)
	var resp []byte
	if err != nil {
		jsonError(w, err.Error(), http.StatusBadRequest)
	} else {
		resp, _ = json.Marshal(map[string]string{
			"section": sectionName,
			"msg":     "section has been created",
		})
	}
	w.Write(resp)
}

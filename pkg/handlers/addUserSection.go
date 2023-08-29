package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (h *Handler) AddUserSection(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user"]
	ttlS := mux.Vars(r)["ttl"]
	ttl, _ := strconv.Atoi(ttlS)
	id, _ := strconv.Atoi(userID)
	list, err := h.Repo.GetSectionList()
	if err != nil {
		jsonError(w, err.Error(), 400)
		return
	}
	sects := make(map[string]string)
	if err := json.Unmarshal(list, &sects); err != nil {
		jsonError(w, err.Error(), 500)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	for sect, percentage := range sects {
		var resp []byte
		perc, err := strconv.Atoi(percentage)
		if err != nil {
			jsonError(w, err.Error(), 500)
			continue
		}
		throw := rand.Intn(100)
		if throw <= perc {
			if err := h.Repo.AddSection(id, sect, ttl); err != nil && !strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				jsonError(w, err.Error(), http.StatusInternalServerError)
			} else if err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				resp, _ = json.Marshal(map[string]string{
					"userID":  userID,
					"section": sect,
					"msg":     "time of section assignment has been updated",
				})
			} else {
				resp, _ = json.Marshal(map[string]string{
					"userID":  userID,
					"section": sect,
					"msg":     "section has been assigned to user",
				})
			}
			w.Write(resp)
		}
	}
}

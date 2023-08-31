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
	ttl, errTtl := strconv.Atoi(ttlS)
	if errTtl != nil {
		jsonError(w, errTtl.Error(), 400)
		return
	}
	id, _ := strconv.Atoi(userID)
	list, err := h.repo.GetSectionList()
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
			err := h.repo.AddSection(id, sect, ttl)
			if err != nil && !strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				jsonError(w, err.Error(), http.StatusInternalServerError)
			}
			if err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				resp, _ = json.Marshal(map[string]string{
					"userID":  userID,
					"section": sect,
					"msg":     "time of section assignment has been updated",
				})
			}
			if err == nil {
				resp, _ = json.Marshal(map[string]string{
					"userID":  userID,
					"section": sect,
					"msg":     "section has been assigned to user",
				})
				w.Write(resp)
			}
			if err := h.repo.AddHistory(id, sect); err != nil && !strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				jsonError(w, err.Error(), 500)
			} else if err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				resp, _ = json.Marshal(map[string]string{
					"msg": "history has been updated",
				})
				w.Write(resp)
			}

		}
	}
}

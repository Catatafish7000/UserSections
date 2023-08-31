package handlers

import (
	"encoding/csv"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

type info struct {
	CreatedAt string
	DeletedAt string
}

func (h *Handler) DownloadHistory(w http.ResponseWriter, r *http.Request) {
	year, errYear := strconv.Atoi(mux.Vars(r)["year"])
	if errYear != nil {
		jsonError(w, "wrong format", 400)
		return
	}
	month, errMonth := strconv.Atoi(mux.Vars(r)["month"])
	resp, err := h.repo.ShowHistory(year, month)
	if err != nil {
		jsonError(w, err.Error(), 500)
	}
	if errMonth != nil {
		jsonError(w, "wrong format", 400)
		return
	}
	data := make(map[string]info, 0)
	err = json.Unmarshal(resp, &data)
	if err != nil {
		jsonError(w, "error retrieving history", 500)
		return
	}
	filePath := "history.csv"
	err = os.Truncate(filePath, 0)
	if err != nil {
		jsonError(w, "error updating history file", 500)
		return
	}
	file, err := os.Create(filePath)
	if err != nil {
		jsonError(w, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"UserID+Section", "CreatedAt", "DeletedAt"}
	writer.Write(headers)

	for i, j := range data {
		row := []string{i, j.CreatedAt, j.DeletedAt}
		err := writer.Write(row)
		if err != nil {
			jsonError(w, "error writing down history", 500)
			return
		}
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename="+file.Name())
	http.ServeFile(w, r, filePath)
}

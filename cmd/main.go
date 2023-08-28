package main

import (
	"Segments/pkg/handlers"
	"Segments/pkg/middleware"
	"Segments/pkg/repo"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	var handler *handlers.Handler
	repo := repo.NewRepo()
	handler = handlers.NewHandler(repo)
	r := mux.NewRouter()
	r.HandleFunc("/new/{section}", handler.NewSection).Methods(http.MethodPost)
	r.HandleFunc("/delete/{section}", handler.Delete).Methods(http.MethodPost)
	r.HandleFunc("/add/{user}/{section}", handler.AddUserSection).Methods(http.MethodPost)
	r.HandleFunc("/remove/{user}/{section}", handler.RemoveUserSection).Methods(http.MethodPost)
	r.HandleFunc("/show/{user}", handler.ShowUserSection).Methods(http.MethodGet)
	mux := middleware.Panic(r)
	port := ":8080"
	log.Println("Запуск сервера на localhost" + port)
	http.ListenAndServe(port, mux)
}

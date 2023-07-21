package router

import (
	"github.com/mikeleo03/Course-Scheduler_Backend/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/mk", middleware.GetAllMataKuliah).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/jurusan", middleware.GetAllJurusan).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/fakul", middleware.GetAllFakultas).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/schedule", middleware.GenerateSchedule).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newmk", middleware.AddMataKuliah).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/newjurusan", middleware.AddJurusan).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/newfakul", middleware.AddFakultas).Methods("POST", "OPTIONS")

	return router
}
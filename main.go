package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mikeleo03/Course-Scheduler_Backend/router"
	"github.com/mikeleo03/Course-Scheduler_Backend/middleware"
	"github.com/mikeleo03/Course-Scheduler_Backend/repository"
)

func main() {
	// Activate router
	r := router.Router()

	// Create DB connection
	db, err := repository.CreateConnection()
	if err != nil {
		log.Fatalf("Failed to create DB connection: %v", err)
	}

	// Set the instance of Repo in middleware package
	middleware.SetRepo(db)

	// Defer close
	defer func() {
		err := db.CloseConnection()
		if err != nil {
			log.Fatal(err)
		}
	}()

    // Start server
	// Use `PORT` provided in environment or default to 8080
  	var port = envPortOr("8080")
	fmt.Println("Starting server...")
	fmt.Println("Listening from" + port)
	log.Fatal(http.ListenAndServe(port, r))
}

// PORT handling
// Returns PORT from environment if found, defaults to
// value in `port` parameter otherwise. The returned port
// is prefixed with a `:`, e.g. `":8080"`.
func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}
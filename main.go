package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mohamedveron/phone_numbers_validation/api"
	"github.com/mohamedveron/phone_numbers_validation/persistence"
	"github.com/mohamedveron/phone_numbers_validation/service"
	"net/http"
	"strconv"
)


func main() {

	database, err := sql.Open("sqlite3", "./sample.db")
	if err != nil {
		fmt.Print(err)
	}

	rows, _ := database.Query("SELECT id, name, phone FROM customer")
	var id int
	var name string
	var phone string
	for rows.Next() {
		rows.Scan(&id, &name, &phone)
		fmt.Println(strconv.Itoa(id) + ": " + name + " " + phone)
	}

	persistenceLayer := persistence.NewPersistence("./sample.db")
	serviceLayer := service.NewService(persistenceLayer)
	server := api.NewServer(serviceLayer)

	// prepare router
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Mount("/", api.Handler(server))
	})

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	fmt.Println("server starting on port 8080...")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("server failed to start", "error", err)
	}

}

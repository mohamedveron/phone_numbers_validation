package main

import (
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mohamedveron/phone_numbers_validation/api"
	"github.com/mohamedveron/phone_numbers_validation/persistence"
	"github.com/mohamedveron/phone_numbers_validation/service"
	"net/http"
)

func main() {

	countriesCodeMap := make(map[string]service.Country)

	countriesCodeMap["Cameroon"] = service.Country{"+237", "\\(237\\)\\ ?[2368]\\d{7,8}$"}

	countriesCodeMap["Ethiopia"] = service.Country{
		"+251",
		"\\(251\\)\\ ?[1-59]\\d{8}$",
	}

	countriesCodeMap["Morocco"] = service.Country{
		"+212",
		"\\(212\\)\\ ?[5-9]\\d{8}$",
	}

	countriesCodeMap["Mozambique"] = service.Country{
		"+258",
		"\\(258\\)\\ ?[28]\\d{7,8}$",
	}

	countriesCodeMap["Uganda"] = service.Country{
		"+256",
		"\\(256\\)\\ ?\\d{9}$",
	}

	//initialize the service layers
	persistenceLayer := persistence.NewPersistence("./sample.db")
	serviceLayer := service.NewService(persistenceLayer, countriesCodeMap)
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

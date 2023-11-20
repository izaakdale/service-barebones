package main

import (
	"net/http"
	"os"
	"time"

	"github.com/izaakdale/lib/response"
	"github.com/izaakdale/lib/router"
	"github.com/izaakdale/lib/server"
)

func main() {
	s, err := server.New(
		"service-barebones",
		router.New(
			router.WithRoute(
				http.MethodGet,
				"/_/alive",
				func(w http.ResponseWriter, r *http.Request) {
					response.WriteJson(w, http.StatusOK, "alive")
					return
				},
			),
		),
		server.WithHost(os.Getenv("HOST")),
		server.WithPort(os.Getenv("PORT")),
		server.WithTimeoutHandler(time.Second, "timedout!"),
	)
	if err != nil {
		panic(err)
	}
	s.Run()
}

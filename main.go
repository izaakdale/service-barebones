package main

import (
	"fmt"
)

func main() {
	fmt.Println("testing lambda func!")
	// s, err := server.New(
	// 	"service-barebones",
	// 	router.New(
	// 		router.WithRoute(
	// 			http.MethodGet,
	// 			"/_/alive",
	// 			func(w http.ResponseWriter, r *http.Request) {
	// 				response.WriteJson(w, http.StatusOK, "alive")
	// 				return
	// 			},
	// 		),
	// 	),
	// 	server.WithHost(os.Getenv("HOST")),
	// 	server.WithPort(os.Getenv("PORT")),
	// 	server.WithTimeoutHandler(time.Second, "timedout!"),
	// )
	// if err != nil {
	// 	panic(err)
	// }
	// s.Run()
}

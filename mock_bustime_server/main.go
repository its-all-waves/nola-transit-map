package main

import (
	"log"
	"net/http"
)

const port = ":8081"

func main() {
	http.HandleFunc(
		"/getvehicles", // ignores query params
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./mock_vehicles.json")
			log.Println("GET /getvehicles :: Served Vehicles")
		},
	)
        http.HandleFunc(
                "/vehicles_jp",
                func(w http.ResponseWriter, r *http.Request) {
                        http.ServeFile(w, r, "./mock_vehicles_jp.protobuf")
                        log.Println("GET /vehicles_jp :: Served Vehicles")
                },
        )
	log.Printf("\n\nMock Clever Devices Bustime server is running at: http://localhost%s \n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

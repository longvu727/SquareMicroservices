package routes

import (
	"fmt"
	"log"
	"net/http"
)

type Handler = func(writer http.ResponseWriter, request *http.Request)

func Register() {
	log.Println("Registering routes")
	routes := map[string]Handler{
		"/":          home,
		"/GetSquare": GetSquare,
	}

	for route, handler := range routes {
		http.HandleFunc(route, handler)
	}
}

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "{\"Acknowledged\": true}")
}

func GetSquare(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Received request for %s\n", request.URL.Path)
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(`GetSquare Service Acknowledged`))
}

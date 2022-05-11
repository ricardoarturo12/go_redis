package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	handlers "github.com/ricardoarturo12/go_redis/handlers"
	"github.com/ricardoarturo12/go_redis/models"
	server "github.com/ricardoarturo12/go_redis/server"
)

func main() {
	serv := server.NewServer()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello)
	mux.HandleFunc("/albums/", handlers.ShowAlbum(serv))
	mux.HandleFunc("/album/", handlers.SetAlbum(serv))
	log.Println("Listening on :4000...")
	http.ListenAndServe(":4000", mux)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hola mundo")
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.MessageResponse{
		Message: "Hola mundo",
		Status:  true,
	})
}

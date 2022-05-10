package main

import (
	"fmt"
	"log"
	"net/http"

	handlers "github.com/ricardoarturo12/go_redis/handlers"
	server "github.com/ricardoarturo12/go_redis/server"
)

func main() {
	serv := server.NewServer()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello)
	mux.HandleFunc("/album/", handlers.ShowAlbum(serv))
	mux.HandleFunc("/albums/", handlers.SetAlbum(serv))
	log.Println("Listening on :4000...")
	http.ListenAndServe(":4000", mux)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hola mundo")

}

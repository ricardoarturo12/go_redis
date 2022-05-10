package album

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/ricardoarturo12/go_redis/models"
	server "github.com/ricardoarturo12/go_redis/server"
)

var ErrNoAlbum = errors.New("no album found")

func ShowAlbum(serv *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, http.StatusText(405), 405)
			return
		}

		fmt.Println(r.URL.Query().Get("id"))
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		if _, err := strconv.Atoi(id); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		bk, err := findAlbum(serv.GetConnect(), id)
		if err == ErrNoAlbum {
			http.NotFound(w, r)
			return
		} else if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.Album{
			Title:  bk.Title,
			Artist: bk.Artist,
			Price:  bk.Price,
			Likes:  bk.Likes,
		})

		// fmt.Fprintf(w, "%s by %s: £%.2f [%d likes] \n", bk.Title, bk.Artist, bk.Price, bk.Likes)

	}
}

func SetAlbum(serv *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			http.Error(w, http.StatusText(405), 405)
			return
		}
		decoder := json.NewDecoder(r.Body)

		var data models.Album
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}

		value, _ := Increment(serv.GetConnect())
		// convertir una interface a string
		value1 := fmt.Sprintf("%v", value)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.MessageResponse{
			Message: "Correcto",
			Status:  true,
		})
		setAlbums(serv.GetConnect(), value1, data.Title, data.Artist, data.Price, data.Likes)
	}
}

func findAlbum(conn redis.Conn, id string) (*models.Album, error) {
	defer conn.Close()

	values, err1 := redis.Values(conn.Do("HGETALL", "album:"+id))
	if err1 != nil {
		return nil, err1
	} else if len(values) == 0 {
		return nil, ErrNoAlbum
	}

	var album models.Album
	err1 = redis.ScanStruct(values, &album)
	if err1 != nil {
		return nil, err1
	}

	log.Println(&album)
	return &album, nil
}

// Incrementa el valor
func Increment(conn redis.Conn) (interface{}, error) {

	value, err1 := conn.Do("INCR", "album")
	if err1 != nil {
		log.Fatal(err1)
		return nil, err1
	}
	return value, nil
}

func setAlbums(conn redis.Conn, id, title, artist string, price float64, likes int) {

	album := "album:" + id

	_, err1 := conn.Do("HMSET", album, "title", title, "artist", artist, "price", price, "likes", likes)
	if err1 != nil {
		log.Fatal(err1)
	}

	defer conn.Close()

}

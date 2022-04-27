package main

import (
	"errors"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

// Declare a pool variable to hold the pool of Redis connections.
var pool *redis.Pool

var ErrNoAlbum = errors.New("no album found")

// Define a custom struct to hold Album data.
type Album struct {
	Title  string  `redis:"title"`
	Artist string  `redis:"artist"`
	Price  float64 `redis:"price"`
	Likes  int     `redis:"likes"`
}

func FindAlbum(id string) (*Album, error) {
	conn := pool.Get()

	defer conn.Close()

	values, err := redis.Values(conn.Do("HGETALL", "album:"+id))
	if err != nil {
		return nil, err
	} else if len(values) == 0 {
		return nil, ErrNoAlbum
	}

	var album Album
	err = redis.ScanStruct(values, &album)
	if err != nil {
		return nil, err
	}

	return &album, nil
}

func SetCount() {
	// SET album "0"
}

// Incrementa el valor
func Increment() (interface{}, error) {
	conn, err := redis.Dial("tcp", "localhost:8080",
		redis.DialPassword(os.Getenv("PASSWORD")))

	value, err1 := conn.Do("INCR", "album")
	if err1 != nil {
		log.Fatal(err)
		return nil, err
	}
	return value, nil
}

func SetAlbums(id, title, artist string, price float64, likes int) {
	conn, err := redis.Dial("tcp", "localhost:8080",
		redis.DialPassword(os.Getenv("PASSWORD")))

	if err != nil {
		log.Fatal(err)
	}

	_, err1 := conn.Do("HMSET", "album:", "title", title, "artist", artist, "price", price, "likes", likes)
	if err1 != nil {
		log.Fatal(err)
	}

	defer conn.Close()

}

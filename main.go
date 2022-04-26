package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

type Album struct {
	Title  string  `redis:"title"`
	Artist string  `redis:"artist"`
	Price  float64 `redis:"price"`
	Likes  int     `redis:"likes"`
}


func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	conn, err := redis.Dial("tcp", "localhost:8080",
		// redis.DialUsername("username"),
		redis.DialPassword("12345"))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	values, err := redis.Values(conn.Do("HGETALL", "album:1"))
	if err != nil {
		log.Fatal(err)
	}

	var album Album

	err = redis.ScanStruct(values, &album)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%+v", album)
}

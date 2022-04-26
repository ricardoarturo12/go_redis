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

	_, err = conn.Do("HMSET", "album:2", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Do("HMSET", "album:3", "title", "TESTS", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("added!")

	price, err := redis.Float64(conn.Do("HGET", "album:1", "price"))
	if err != nil {
		log.Fatal(err)
	}

	likes, err := redis.Int(conn.Do("HGET", "album:1", "likes"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Price: %.2f [likes: %d] \n", price, likes)
}

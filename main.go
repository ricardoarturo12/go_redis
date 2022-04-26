package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	conn, err := redis.Dial("tcp", "localhost:8080",
		// redis.DialUsername("username"),
		redis.DialPassword(os.Getenv("PASSWORD")))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	_, err = conn.Do("HMSET", "album:2", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Electric Ladyland added!")

	likes, err := redis.Int(conn.Do("HGET", "album:2", "likes"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[%d likes]\n", likes)
}

package server

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

type Server struct {
	pool *redis.Pool
}

func NewServer() *Server {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	fmt.Println("inicia servidor")
	server := &Server{
		pool: &redis.Pool{
			MaxIdle:     10,
			MaxActive:   12000,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", "redis_db:6379",
					redis.DialPassword(os.Getenv("PASSWORD")))
			},
		},
	}
	return server
}

func (server *Server) GetConnect() redis.Conn {
	// fmt.Println("obtiene getconnect")
	return server.pool.Get()
}

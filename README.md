# Acceso a datos

```


docker-compose exec redis_db /bin/sh

/data # redis-cli
127.0.0.1:6379> auth password
OK
127.0.0.1:6379> keys *
(empty array)
127.0.0.1:6379>

```

# Docker-compose

```

version: '3'
services:
  redis:
    image: redis:latest
    command: ["redis-server", "/etc/redis/redis.conf"]
    volumes:
      - ./redis.conf:/etc/redis/redis.conf
    ports:
      - "6379:6379"

```

# Conexión y comandos básicos
```
conn, err := redis.Dial("tcp", "localhost:8080",
		redis.DialPassword(os.Getenv("PASSWORD")))
	if err != nil {
		log.Fatal(err)
	}


conn.Do("HMSET", "album:3", "title", "TESTS", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)

conn.Do("HMSET", "album:2", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)

price, _ := redis.Float64(conn.Do("HGET", "album:1", "price"))


values, err := redis.Values(conn.Do("HGETALL", "album:1"))
	if err != nil {
		log.Fatal(err)
	}


var album Album
redis.ScanStruct(values, &album)
fmt.Printf("%+v", album)


```

# Postman
- Method: POST
- Body -> raw
```

{"title": "ricardo", "artist": "arturo", "price": 5.90, "likes": 9}

```

# Para countd
https://redis.io/commands/incr/

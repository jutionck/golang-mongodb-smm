package main

import "enigmacamp.com/golang-with-mongodb/delivery"

func main() {
	delivery.NewServer().Run()
}

// RUN
// MONGO_HOST=207.148.125.99 MONGO_PORT=27017 MONGO_DB=enigma MONGO_USER=jutionck MONGO_PASSWORD=password API_HOST=localhost API_PORT=8888 go run .

package main

import (
	"log"
	"url-shortener/internal/routes"
	"url-shortener/internal/storage"
)

func main() {
	db, err := storage.InitDB()
	if err != nil {
		log.Fatal("Gagal koneksi database:", err)
	}

	rdb, _ := storage.InitRedis()

	r := routes.SetupRouter(db, rdb)

	log.Println("Server running on :8080")
	r.Run(":8080")
}

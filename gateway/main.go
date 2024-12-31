package main

import (
	"log"
	"net/http"

	"gihtub.com/sandy088/common"
	_ "github.com/joho/godotenv/autoload"
)

var (
	httpAddress = common.EnvString("HTTP_ADDRESS", ":8080")
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	if err := http.ListenAndServe(httpAddress, mux); err != nil {
		log.Fatal("Failed to start server", err)
	}
}

//Gateway is the main entry point for the application

// - It is the most simple one, because only acts as a transport layer, Have no business logic, stores, etc.

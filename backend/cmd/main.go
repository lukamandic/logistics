package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
	"github.com/lukamandic/logistics/backend/internal/api/handlers"
	"github.com/lukamandic/logistics/backend/internal/api/middleware"
	"github.com/lukamandic/logistics/backend/internal/config"
	"github.com/lukamandic/logistics/backend/internal/repository"
	"github.com/lukamandic/logistics/backend/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	ctx := context.Background()
	repo, err := repository.NewDynamoClient(ctx, cfg.TableName)
	if err != nil {
		log.Fatalf("Failed to initialize DynamoDB client: %v", err)
	}

	parcelService := service.NewParcelService(repo)

	parcelHandler := handlers.NewParcelHandler(parcelService)

	mux := http.NewServeMux()

	handler := middleware.CORSMiddleware(cfg)(mux)

	mux.HandleFunc("/api/v1/parcel-sizes", parcelHandler.HandleParcelRoutes)
	mux.HandleFunc("/api/v1/calculate", parcelHandler.HandleCalculateDistribution)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}
		handlers.ServeStaticFiles(w, r)
	})

	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	fmt.Printf("Starting HTTP server on %s...\n", serverAddr)
	if err := http.ListenAndServe(serverAddr, handler); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
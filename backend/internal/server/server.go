package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	connectcors "connectrpc.com/cors"
	"github.com/ZanzyTHEbar/goflexpro/internal/adapters/secondary/persistence"
	"github.com/ZanzyTHEbar/goflexpro/internal/core/services"
	"github.com/ZanzyTHEbar/goflexpro/internal/dto/db"
	"github.com/ZanzyTHEbar/goflexpro/pkgs/config"
	productv1connect "github.com/ZanzyTHEbar/goflexpro/pkgs/gen/product/v1/v1connect"
	goflexprologger "github.com/ZanzyTHEbar/goflexpro/pkgs/logger"
	_ "github.com/jpfuentes2/go-env/autoload"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	db     *db.PrismaClient
	config config.Config
}

func NewServer(config config.Config) *Server {

	// Handle the location of the db_url
	err := os.Setenv("DATABASE_URL", config.DBHost)
	log.Printf("Config: %s", config.DBHost)

	log.Printf("DATABASE_URL: %s", os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("Failed to set DATABASE_URL", "error", err)
	}

	dbClient := db.NewClient()
	if err := dbClient.Prisma.Connect(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return &Server{
		db:     dbClient,
		config: config,
	}
}

func (server *Server) Start() {
	goflexprologger.InitLogger(server.config)

	productRepo := persistence.NewPrismaProductAdapter(server.db)

	mux := http.NewServeMux()

	// Register services here
	// We will need to instantiate and add the various services, etc.
	productService := services.NewProductService(productRepo)

	// Instantiate server and handlers
	productPath, productHandler := productv1connect.NewProductServiceHandler(productService)

	// Register handlers
	mux.Handle(productPath, productHandler)

	// Add CORS middleware
	corsHandler := withCORS(mux)

	slog.Info("Starting server", "addr", fmt.Sprintf(":%d", server.config.HttpPort))

	http.ListenAndServe(
		fmt.Sprintf(":%d", server.config.HttpPort),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(corsHandler, &http2.Server{}),
	)

	// Create a new signal.NotifyContext to handle graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	slog.Info("Shutting down server")

	if err := server.db.Prisma.Disconnect(); err != nil {
		slog.Error("Failed to disconnect from database", "error", err)
	}

	slog.Info("Server shutdown complete")
}

// withCORS adds CORS support to a Connect HTTP handler.
func withCORS(h http.Handler) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})
	return middleware.Handler(h)
}

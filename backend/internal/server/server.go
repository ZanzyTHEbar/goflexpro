package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jpfuentes2/go-env/autoload"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/ZanzyTHEbar/goflexpro/internal/adapters/secondary/persistence"
	"github.com/ZanzyTHEbar/goflexpro/internal/core/services"
	"github.com/ZanzyTHEbar/goflexpro/internal/dto/db"
	"github.com/ZanzyTHEbar/goflexpro/pkgs/config"
	productv1connect "github.com/ZanzyTHEbar/goflexpro/pkgs/gen/product/v1/v1connect"
	goflexprologger "github.com/ZanzyTHEbar/goflexpro/pkgs/logger"
)

type Server struct {
	db     *db.PrismaClient
	config config.Config
}

func NewServer(config config.Config) *Server {

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

	// TODO: Register services here
	// We will need to instantiate and add the various services, etc.
	productService := services.NewProductService(productRepo)

	// Instantiate server and handlers
	productPath, productHandler := productv1connect.NewProductServiceHandler(productService)

	// Register handlers
	mux.Handle(productPath, productHandler)

	// Create a new listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", server.config.HttpPort))
	if err != nil {
		slog.Error("Failed to create listener", "error", err)
		return
	}

	// Create a new signal.NotifyContext to handle graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Start the server in a goroutine so that it doesn't block
	go func() {
		slog.Info("Starting server", "addr", listener.Addr().String())

		if err := http.Serve(
			listener,
			h2c.NewHandler(mux, &http2.Server{}),
		); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Server failed to start", "error", err)
		}
	}()

	<-ctx.Done()

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	slog.Info("Shutting down server")

	if err := listener.Close(); err != nil {
		slog.Error("Failed to close listener", "error", err)
	}

	if err := server.db.Prisma.Disconnect(); err != nil {
		slog.Error("Failed to disconnect from database", "error", err)
	}

	slog.Info("Server shutdown complete")
}

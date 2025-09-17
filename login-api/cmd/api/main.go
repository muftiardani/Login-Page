package main

import (
	"context"
	"errors"
	"login-api/internal/config"
	"login-api/internal/handler"
	"login-api/internal/router"
	"login-api/internal/service"
	"login-api/internal/storage/postgres"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfg := config.New()

	dbpool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Tidak dapat membuat koneksi pool")
	}
	defer dbpool.Close()

	if err := dbpool.Ping(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("Tidak dapat melakukan ping ke database")
	}
	log.Info().Msg("Database berhasil terhubung!")

	// Inisialisasi lapisan penyimpanan (storage)
	userStore := postgres.NewPostgresUserStore(dbpool)
	paymentStore := postgres.NewPostgresPaymentStore(dbpool)

	jwtKey := []byte(cfg.JWTSecretKey)
	addr := cfg.ServerAddress

	// Inisialisasi lapisan layanan (service)
	authService := service.NewAuthService(userStore, jwtKey)

	// Suntikkan service ke dalam handler, bukan store langsung
	authHandler := handler.NewAuthHandler(authService, jwtKey)
	paymentHandler := handler.NewPaymentHandler(paymentStore)
	dashboardHandler := handler.NewDashboardHandler(paymentStore)

	// Buat router dengan handler yang sudah diinisialisasi
	r := router.NewRouter(authHandler, paymentHandler, dashboardHandler)

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// Jalankan server dalam goroutine agar tidak memblokir
	go func() {
		log.Info().Msgf("Server Go API dengan JWT berjalan di http://localhost%s", addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Server gagal memulai")
		}
	}()

	// Menunggu sinyal interupsi untuk mematikan server secara graceful
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Server sedang dimatikan...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server gagal dimatikan secara graceful")
	}

	log.Info().Msg("Server berhasil dimatikan.")
}
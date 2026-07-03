package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/yudhisrana/go-hospital/internal/application/patient/usecase"
	"github.com/yudhisrana/go-hospital/internal/infra/config"
	"github.com/yudhisrana/go-hospital/internal/infra/persistence/relational"
	"github.com/yudhisrana/go-hospital/internal/infra/persistence/relational/sqlite/patient_repo"
	"github.com/yudhisrana/go-hospital/internal/interface/http"
	"github.com/yudhisrana/go-hospital/internal/interface/http/handler/patient_handler"
	"github.com/yudhisrana/go-hospital/internal/interface/http/routes"
)

func main() {
	cfg := config.Load()

	db, err := relational.NewDatabaseService(cfg.DBCfg)
	if err != nil {
		log.Fatalf("Gagal tersambung ke database: %v", err)
	}

	patientRepo := patient_repo.NewRepository(db.DB())
	patientUsecase := usecase.NewPatientUseCase(patientRepo)
	patientHandler := patient_handler.NewPatientHandler(patientUsecase)

	srv := http.NewServer(cfg.AppCfg)

	routes.RegisterPatientRoutes(srv.App(), patientHandler)

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Jalankan server di port %s", cfg.AppCfg.AppPort)
		if err := srv.Start(cfg.AppCfg.AppPort); err != nil {
			serverErrors <- err
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case err := <-serverErrors:
		log.Fatalf("Server error: %v", err)
	case <-ctx.Done():
		log.Println("Menghentikan server...")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Menghentikan server gagal: %v", err)
	}

	if err := db.Close(); err != nil {
		log.Fatalf("Penutupan database gagal: %v", err)
	}

	log.Println("Server berhasil dihentikan")
}

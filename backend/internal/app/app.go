package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"organizational-structure-visualization-service/internal/config"
	controller "organizational-structure-visualization-service/internal/controllers/employees"
	http_lib "organizational-structure-visualization-service/internal/lib/http"
	model "organizational-structure-visualization-service/internal/models/employees"
	"organizational-structure-visualization-service/pkg/logger/sl"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	db  *pgxpool.Pool
	srv *http.Server
	log *slog.Logger
}

func New(cfg *config.Config, log *slog.Logger) *App {
	pool, err := pgxpool.New(context.Background(), cfg.Database.URL)
	if err != nil {
		log.Error("failed to init database connection", sl.Err(err))
		os.Exit(1)
	}

	model := model.NewModel(pool)

	r := chi.NewRouter()

	r.Use(http_lib.TraceID)
	r.Use(middleware.RealIP)
	r.Use(http_lib.Logging(log))
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		ctr := controller.New(
			controller.ControllerConfig{
				Model: model,
			},
		)
		r.Mount("/employees", ctr.Register())
	})

	srv := &http.Server{
		Addr:        cfg.HTTPServer.Address,
		Handler:     r,
		IdleTimeout: cfg.HTTPServer.IDLETimeout,
	}

	return &App{
		db:  pool,
		srv: srv,
		log: log,
	}
}

func (a *App) Start() {
	const op = "app.Start"

	go func() {
		if err := a.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.log.With(slog.String("op", op)).
				Error("failed to start server", sl.Err(err))
		}
	}()
}

func (a *App) Stop() {
	const op = "app.Stop"

	a.db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		a.log.With(slog.String("op", op)).
			Error("failed to shutdown server", sl.Err(err))
	}
}

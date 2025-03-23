package api

import (
	"context"
	"github.com/KinitaL/testovoye/config"
	"github.com/KinitaL/testovoye/internal/infrastructure/controllers"
	"github.com/KinitaL/testovoye/internal/infrastructure/repositories/books"
	"github.com/KinitaL/testovoye/internal/server"
	"github.com/KinitaL/testovoye/internal/usecases"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	config   *config.Config
	logger   *zap.Logger
	stopFunc context.CancelFunc
}

func NewApp(
	config *config.Config,
	logger *zap.Logger,
) *App {
	// create instance of application
	return &App{
		config: config,
		logger: logger,
	}
}

func (app *App) Configure(ctx context.Context) error {
	// initialize all components of application:
	// - controllers
	// - repositories
	// - usecases and etc.
	app.logger.Info("api app configuration started", zap.String("time", time.Now().Format("2006-01-02 15:04:05")))

	return nil
}

func (app *App) Run() error {
	// run all components in sufficient sequence
	// in case of 'fatal' errors, gracefully stop the application
	app.logger.Info("api app started", zap.String("time", time.Now().Format("2006-01-02 15:04:05")))

	ctx := context.Background()

	ctx, app.stopFunc = signal.NotifyContext(ctx, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM) // signals to graceful shutdown
	defer app.stopFunc()

	s := server.BuildServer(app.config.Service,
		middleware.RequestID(),
		server.ZapLogger(app.logger),
	)

	repsRegistry := usecases.NewRepositoriesRegistry(books.NewInMemoryRepo())
	ucRegistry := usecases.NewRegistry(repsRegistry)

	controllers.Register(s, ucRegistry)

	appErrors := make(chan error, 1)
	go func() {
		appErrors <- s.Start("")
	}()

	shutdown := func(ctx context.Context) error {
		if err := s.Shutdown(ctx); err != nil {
			app.logger.Error("server.shutdown", zap.Error(err))
			return err
		}
		return nil
	}

	select {
	case err := <-appErrors:
		if err != nil {
			app.logger.Error("server.listen", zap.Error(err))
			return err
		}
	case <-ctx.Done():
		app.logger.Info("start shutdown", zap.String("reason", ctx.Err().Error()))
		defer app.stopFunc()
		err := shutdown(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

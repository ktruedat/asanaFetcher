package app

import (
	"encoding/json"
	"os"
	"time"

	"github.com/ktruedat/recoAssignment/config"
	"github.com/ktruedat/recoAssignment/internal/domain/entities"
	"github.com/ktruedat/recoAssignment/internal/infra/http"
	"github.com/ktruedat/recoAssignment/internal/infra/http/client"
	"github.com/ktruedat/recoAssignment/internal/infra/http/limiter"
	"github.com/ktruedat/recoAssignment/internal/infra/service"
	"github.com/ktruedat/recoAssignment/pkg/log"
	"github.com/pkg/errors"
)

type App struct {
	cfg           *config.Config
	logger        log.Logger
	userSvc       service.ResourceGetService[entities.User]
	projectSvc    service.ResourceGetService[entities.Project]
	limiter       http.Limiter
	pullingTicker *time.Ticker
	startChan     chan struct{}
	stopChan      chan struct{}
}

const configPath string = "./config.json"

func Run() error {
	app, err := NewApp()
	if err != nil {
		return err
	}

	return app.run()
}

func NewApp() (*App, error) {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		return nil, err
	}

	logger := log.NewLogger()

	logger.Debug("Config", "cfg", *cfg)

	lm := limiter.NewLimiter()

	cl := client.NewClient(cfg.BaseURL, cfg.APIToken, lm, logger)
	userSvc := service.NewUserGetService(cl, cfg)
	projectSvc := service.NewProjectGetService(cl, cfg)

	return &App{
		cfg:           cfg,
		logger:        logger,
		userSvc:       userSvc,
		projectSvc:    projectSvc,
		limiter:       lm,
		pullingTicker: time.NewTicker(cfg.ExtractionRateDuration),
		startChan:     make(chan struct{}),
		stopChan:      make(chan struct{}),
	}, nil
}

func (a *App) run() error {
	go func() {
		a.startChan <- struct{}{}
	}()
	for {
		select {
		case <-a.startChan:
			if err := a.fetchData(); err != nil {
				a.logger.Debug("Starting fetching")
				return errors.Wrap(err, "failed to fetch data")
			}
		case <-a.pullingTicker.C:
			a.logger.Debug("Tick")
			a.stopChan <- struct{}{}
			go func() {
				time.Sleep(a.cfg.ExtractionRateDuration)
				a.startChan <- struct{}{}
			}()
		}
	}
}

func (a *App) fetchData() error {
	for {
		select {
		case <-a.limiter.Tries():
			users, err := a.userSvc.Get()
			if err != nil {
				return errors.Wrap(err, "failed to fetch users")
			}

			projects, err := a.projectSvc.Get()
			if err != nil {
				return errors.Wrap(err, "failed to fetch projects")
			}

			if err := a.saveUsersInFile(users); err != nil {
				return errors.Wrap(err, "failed to save user data")
			}

			if err := a.saveProjectsInFile(projects); err != nil {
				return errors.Wrap(err, "failed to save project data")
			}
		case <-a.stopChan:
			a.logger.Debug("Stopping fetching")
			return nil
		}
	}
}

func (a *App) saveUsersInFile(users []entities.User) error {
	a.logger.Debug("saving user data")
	marshalled, err := json.Marshal(users)
	if err != nil {
		return errors.Wrap(err, "failed to marshal user info")
	}

	if err := os.WriteFile("./users.json", marshalled, 0o600); err != nil {
		return errors.Wrap(err, "failed to save user data")
	}

	return nil
}

func (a *App) saveProjectsInFile(projects []entities.Project) error {
	a.logger.Debug("saving project data")

	marshalled, err := json.Marshal(projects)
	if err != nil {
		return errors.Wrap(err, "failed to marshal project info")
	}

	if err := os.WriteFile("./projects.json", marshalled, 0o600); err != nil {
		return errors.Wrap(err, "failed to save project data")
	}

	return nil
}

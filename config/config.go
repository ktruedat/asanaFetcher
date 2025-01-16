package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

// Config is the app configuration structure. It holds the APIToken needed to access the Asana API, the WorkspaceGID,
// and the extractionRate which represents the interval between data fetches from the API.
type Config struct {
	APIToken               string `yaml:"apiToken"`
	WorkspaceGID           string `yaml:"workspaceGID"`
	ExtractionRateString   string `yaml:"extractionRateString"`
	ExtractionRateDuration time.Duration
	BaseURL                string `yaml:"baseURL"`
}

// NewConfig parses the configuration file specified in filePath and returns a Config instance initialized with the
// values from the config file.
func NewConfig(filePath string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(filePath, &cfg); err != nil {
		return nil, errors.Wrap(err, "failed to parse configuration file")
	}

	var err error
	cfg.ExtractionRateDuration, err = time.ParseDuration(cfg.ExtractionRateString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse duration")
	}

	return &cfg, nil
}

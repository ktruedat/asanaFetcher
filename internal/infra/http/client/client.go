package client

import (
	"encoding/json"
	nethttp "net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/ktruedat/recoAssignment/internal/infra/http"
	"github.com/ktruedat/recoAssignment/pkg/log"
	"github.com/pkg/errors"
)

const (
	acceptHeader              = "Accept"
	authorizationHeader       = "Authorization"
	applicationJSONAcceptType = "application/json"
)

type client struct {
	rc      *resty.Client
	logger  log.Logger
	token   string
	limiter http.Limiter
}

func NewClient(baseURL, token string, limiter http.Limiter, logger log.Logger) http.Client {
	rc := resty.New().
		SetHeaders(
			map[string]string{
				acceptHeader:        applicationJSONAcceptType,
				authorizationHeader: "Bearer " + token,
			},
		).
		SetDebug(false).
		SetBaseURL(baseURL)

	return &client{
		rc:      rc,
		logger:  logger,
		token:   token,
		limiter: limiter,
	}
}

func (c *client) Get(url string, response any) error {
	body, err := c.GetRaw(url)
	if err != nil {
		return errors.Wrap(err, "failed to perform GET request")
	}

	if err := json.Unmarshal(body, response); err != nil {
		return errors.Wrap(err, "failed to unmarshal data")
	}

	return nil
}

func (c *client) GetRaw(url string) ([]byte, error) {
	c.logger.Debug("Performing GET request", "url", url)

	resp, err := c.rc.R().
		Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform GET request")
	}

	if resp.StatusCode() == nethttp.StatusTooManyRequests {
		c.logger.Warning("Rate limit exceeded")
		retryAfter := resp.Header().Get("Retry-After")
		delay, err := parseRetryAfter(retryAfter)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse retry after header")
		}

		c.logger.Debug("limit exceeded, sleeping or delay", "delay", delay)
		time.Sleep(delay)
	}

	return resp.Body(), nil
}

func parseRetryAfter(header string) (time.Duration, error) {
	if header == "" {
		return 0, errors.New("header is missing ")
	}

	seconds, err := strconv.Atoi(header)
	if err != nil {
		return 0, errors.Wrap(err, "failed to convert time")
	}

	return time.Duration(seconds), nil
}

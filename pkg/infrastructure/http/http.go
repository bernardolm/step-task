package http

import (
	"context"
	errorsnative "errors"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/bernardolm/step-task/pkg/infrastructure/errors"
)

const (
	userAgent = "BernardoLm/StepTask"
)

type client struct {
	baseURL string
	client  http.Client
}

func (c client) Get(req *http.Request) (*http.Response, error) {
	log.WithField("url", req.URL.String()).Debug("http client: requesting from")

	req.Header.Set("user-agent", userAgent)
	req.Method = http.MethodGet

	res, err := c.client.Do(req)
	if err != nil {
		return nil, errorsnative.Join(err, fmt.Errorf("http client: failed to do the request"))
	}

	if res.StatusCode >= http.StatusBadRequest {
		err := errors.New("http client: failed to get from server: %s", res.Status)
		err.WithField("status_code", res.StatusCode)
		err.WithField("url", req.URL.String())
		return nil, err
	}

	return res, nil
}

func (c client) NewRequest(ctx context.Context) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "", c.baseURL, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func New(alias string) (*client, error) {
	alias = strings.ToUpper(alias)

	baseURL := fmt.Sprintf("%s_BASE_API", alias)
	baseURL = viper.GetString(baseURL)

	hdc := http.DefaultClient
	hdc.Timeout = viper.GetDuration("HTTP_TIMEOUT")

	c := client{
		baseURL: baseURL,
		client:  *hdc,
	}

	return &c, nil
}

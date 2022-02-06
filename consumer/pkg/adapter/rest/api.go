package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/isutare412/istio-playground/consumer/pkg/config"
	"github.com/isutare412/istio-playground/consumer/pkg/core/user"
	"github.com/opentracing/opentracing-go"
)

type apiRest struct {
	addr   string
	client *http.Client
}

func (a *apiRest) GetUser(ctx context.Context, name string) (*user.User, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "apiRest.GetUser")
	defer span.Finish()

	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/api/v1/hello/%s", a.addr, name), nil)
	if err != nil {
		return nil, fmt.Errorf("apiRest.GetUser: %w", err)
	}

	err = span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)
	if err != nil {
		return nil, fmt.Errorf("apiRest.GetUser: %w", err)
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("apiRest.GetUser: %w", err)
	}
	defer resp.Body.Close()

	var user user.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("apiRest.GetUser: %w", err)
	}
	return &user, nil
}

func NewApiRest(cfg *config.ApiServerConfig) *apiRest {
	return &apiRest{
		addr: cfg.Addr,
		client: &http.Client{
			Timeout: time.Millisecond * time.Duration(cfg.Timeout),
		},
	}
}

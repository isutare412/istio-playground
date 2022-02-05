package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/isutare412/istio-playground/consumer/pkg/config"
	"github.com/isutare412/istio-playground/consumer/pkg/core/user"
)

type api struct {
	addr   string
	client *http.Client
}

func (a *api) GetUser(ctx context.Context, name string) (*user.User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/api/v1/hello/%s", a.addr, name), nil)
	if err != nil {
		return nil, fmt.Errorf("api.GetUser: %w", err)
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("api.GetUser: %w", err)
	}
	defer resp.Body.Close()

	var user user.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("api.GetUser: %w", err)
	}
	return &user, nil
}

func NewApi(cfg *config.ApiServerConfig) *api {
	return &api{
		addr: cfg.Addr,
		client: &http.Client{
			Timeout: time.Millisecond * time.Duration(cfg.Timeout),
		},
	}
}

package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/isutare412/istio-playground/api-server/pkg/config"
	"github.com/isutare412/istio-playground/api-server/pkg/core/user"
)

type userRest struct {
	addr   string
	client *http.Client
}

func (ur *userRest) Name() string {
	return "userRest"
}

func (ur *userRest) IsHealthy() error {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/readiness", ur.addr), nil)
	if err != nil {
		return fmt.Errorf("userRest.IsHealthy: %w", err)
	}

	resp, err := ur.client.Do(req)
	if err != nil {
		return fmt.Errorf("userRest.IsHealthy: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("response[%d] is not ready", resp.StatusCode)
	}
	return nil
}

func (ur *userRest) GetUser(ctx context.Context, name string) (*user.User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/api/v1/users/%s", ur.addr, name), nil)
	if err != nil {
		return nil, fmt.Errorf("userRest.GetUser: %w", err)
	}

	resp, err := ur.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("userRest.GetUser: %w", err)
	}
	defer resp.Body.Close()

	var user user.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("userRest.GetUser: %w", err)
	}
	return &user, nil
}

func NewUserRest(cfg *config.UserServerConfig) *userRest {
	return &userRest{
		addr: cfg.Addr,
		client: &http.Client{
			Timeout: time.Millisecond * time.Duration(cfg.Timeout),
		},
	}
}

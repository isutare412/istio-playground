package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/isutare412/istio-playground/api-server/pkg/config"
)

type Service interface {
	GetUser(ctx context.Context, name string) (*User, error)
}

type service struct {
	addr   string
	client *http.Client
}

func (s *service) GetUser(ctx context.Context, name string) (*User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/api/v1/users/%s", s.addr, name), nil)
	if err != nil {
		return nil, fmt.Errorf("on get user: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("on get user: %w", err)
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("on get user: %w", err)
	}
	return &user, nil
}

func NewService(cfg *config.UserServerConfig) (Service, error) {
	return &service{
		addr: cfg.Addr,
		client: &http.Client{
			Timeout: time.Millisecond * time.Duration(cfg.Timeout),
		},
	}, nil
}

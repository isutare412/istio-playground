package user

import (
	"context"
	"fmt"
)

type Service interface {
	GetUser(ctx context.Context, name string) (*User, error)
}

type service struct {
	rest Rest
}

func (s *service) GetUser(ctx context.Context, name string) (*User, error) {
	usr, err := s.rest.GetUser(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("service.GetUser: %w", err)
	}
	return usr, nil
}

func NewService(rest Rest) Service {
	return &service{
		rest: rest,
	}
}

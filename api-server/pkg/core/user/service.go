package user

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
)

type Service interface {
	GetUser(ctx context.Context, name string) (*User, error)
}

type service struct {
	rest Rest
}

func (s *service) GetUser(ctx context.Context, name string) (*User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.GetUser")
	defer span.Finish()

	usr, err := s.rest.GetUser(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("service.GetUser: %w", err)
	}
	return usr, nil
}

func NewService(rest Rest) (Service, error) {
	return &service{
		rest: rest,
	}, nil
}

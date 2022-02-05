package health

import "fmt"

type Service interface {
	Liveness() error
	Readiness() error
}

type service struct {
	deps []Dependency
}

func (s *service) Liveness() error {
	return nil
}

func (s *service) Readiness() error {
	for _, d := range s.deps {
		if err := d.IsHealthy(); err != nil {
			return fmt.Errorf("dependency[%s] not healthy: %w", d.Name(), err)
		}
	}
	return nil
}

func NewService(deps ...Dependency) Service {
	return &service{
		deps: deps,
	}
}

package health

type Dependency interface {
	Name() string
	IsHealthy() error
}

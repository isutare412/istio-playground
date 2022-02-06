package ticker

import (
	"context"
	"time"

	"github.com/isutare412/istio-playground/consumer/pkg/config"
	"github.com/isutare412/istio-playground/consumer/pkg/core/user"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
)

type ticker struct {
	interval time.Duration
	uSvc     user.Service
	done     chan struct{}
}

func (t *ticker) Start(ctx context.Context) {
	tkr := time.NewTicker(t.interval)
	go func() {
		defer tkr.Stop()
		defer close(t.done)

	loop:
		for {
			t.handleUser(context.Background(), "Suhyuk")

			select {
			case <-tkr.C:
				continue
			case <-ctx.Done():
				break loop
			}
		}
	}()
}

func (t *ticker) Done() <-chan struct{} {
	return t.done
}

func (t *ticker) handleUser(ctx context.Context, name string) {
	span := opentracing.GlobalTracer().StartSpan("ticker.handleUser")
	ctx = opentracing.ContextWithSpan(ctx, span)
	defer span.Finish()

	usr, err := t.uSvc.GetUser(ctx, name)
	if err != nil {
		log.Errorf("failed to get user: %v", err)
		return
	}
	log.Infof("got user[%+v]", usr)
}

func NewTicker(cfg *config.TickerConfig, uSvc user.Service) *ticker {
	return &ticker{
		interval: time.Millisecond * time.Duration(cfg.Interval),
		uSvc:     uSvc,
		done:     make(chan struct{}),
	}
}

package tracer

import log "github.com/sirupsen/logrus"

type jaegerLogger struct {
}

func (l *jaegerLogger) Error(msg string) {
	log.Error(msg)
}

func (l *jaegerLogger) Infof(msg string, args ...interface{}) {
	log.Infof(msg, args...)
}

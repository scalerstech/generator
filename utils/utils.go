package utils
import (
	log "github.com/sirupsen/logrus"
)

type Utils struct {
	debug bool
	l *log.Logger
}

func New() *Utils{
	return &Utils{
		debug: false,
		l: log.New(),
	}
}
package lib

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

func CreateLogger(name string) *log.Entry {
	log.SetHandler(text.Default)
	return log.WithField("from", name)
}

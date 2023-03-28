package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

func InitLog() {
	if gin.IsDebugging() {
		log.SetLevel(log.DebugLevel)
		log.SetFormatter(&log.TextFormatter{
			TimestampFormat: time.DateTime,
			FullTimestamp:   true,
		})
	} else {
		log.SetLevel(log.InfoLevel)
		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat: time.DateTime,
		})
	}
}

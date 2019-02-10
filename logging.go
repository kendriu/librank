package librank

import (
	log "github.com/Sirupsen/logrus"
)

//SetupLogger sets up default logger
func SetupLogger() {
	log.SetFormatter(&log.TextFormatter{})
	if DEBUG == 1 {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

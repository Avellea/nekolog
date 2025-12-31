package main

import (
	"fmt"
	"github.com/avellea/nekolog"
)

func main() {

	log := nekolog.NewLogger()

	log.Info("Application starting...")

	user := "User"
	port := 1337
	log.Infof("User %s connected on port %d", user, port)

	log.Warn("Cache is almost full")
	log.Warnf("Retry attempt %d failed", 2)

	err := fmt.Errorf("disk not found")
	log.Error(err.Error())
	log.Errorf("Fatal error: %v", err)
}

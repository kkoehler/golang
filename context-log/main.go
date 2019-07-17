package main

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func main() {

	type keyval string

	ctx := context.Background()
	vctx := context.WithValue(ctx, keyval("request-id"), keyval("123"))

	//Für JSON Ausgabe
	//log.SetFormatter(&log.JSONFormatter{})

	log.WithField("request-id", vctx.Value(keyval("request-id"))).Info("Output with Field")

	requestLogger := log.WithFields(
		log.Fields{
			"request-id": vctx.Value(keyval("request-id")),
		})
	requestLogger.Info("Some important message")
	requestLogger.Info("Next important message")

}

package main

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func main() {

	type keyval string

	//Erzeugung Parent-Context
	ctx := context.Background()
	//hier wird der Context mit Werten erstellt
	vctx := context.WithValue(ctx, keyval("request-id"), keyval("123"))

	//Für JSON Ausgabe nächste Zeile einkommentieren
	//log.SetFormatter(&log.JSONFormatter{})

	//Logrus Logger vorkonfigurieren
	requestLogger := log.WithFields(
		log.Fields{
			"request-id": vctx.Value(keyval("request-id")),
		})

	//Loggen ohne angabe weiterer Werte
	//diese sind im Logger vorkonfiguriert
	requestLogger.Info("Some important message")
	requestLogger.Info("Next important message")

}

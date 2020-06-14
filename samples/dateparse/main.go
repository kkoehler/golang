package main

import (
	"log"
	"time"
)

func main() {
	expDate, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if err != nil {
		log.Fatalf("Could not parse date because of %v", err)
	}
	log.Printf("dat is %v", expDate)
}

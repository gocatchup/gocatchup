// Package main the the entry point for the gocatcup command.
package main

import (
	"log"

	"github.com/gocatchup/gocatchup/internal/version"
)

func main() {
	log.Println(version.Version)

	log.Println("Go catchup!")
}

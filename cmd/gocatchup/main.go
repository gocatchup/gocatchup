// Package main the the entry point for the gocatcup command.
package main

import (
	"github.com/gocatchup/gocatchup/internal/version"
)

func main() {
	println(version.Version)

	println("Go catchup!")
}

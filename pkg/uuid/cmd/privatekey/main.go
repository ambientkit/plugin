package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ambientkit/plugin/pkg/uuid"
)

func init() {
	// Verbose logging with file name and line number.
	log.SetFlags(log.Lshortfile)

	// Set the time zone.
	tz := os.Getenv("AMB_TIMEZONE")
	if len(tz) > 0 {
		os.Setenv("TZ", tz)
	}
}

func main() {
	// Generate a new private key for AES-256.
	key := uuid.EncodedString(32)
	fmt.Printf("AMB_SESSION_KEY=%v\n", key)
}

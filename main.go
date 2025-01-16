package main

import (
	"log"

	"github.com/ktruedat/recoAssignment/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("failed to run the application: %v", err)
	}
}

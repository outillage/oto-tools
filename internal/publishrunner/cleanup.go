package publishrunner

import (
	"log"
	"os"
)

// cleanup removes authentication files after publishing.
func cleanup() {
	err := os.Remove("./.npmrc")

	if err != nil {
		log.Printf("Failed to clean up authentication files %v", err)
	}
}

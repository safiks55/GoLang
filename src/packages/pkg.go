package pkg

import (
	"log"
	"time"
)

func execTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

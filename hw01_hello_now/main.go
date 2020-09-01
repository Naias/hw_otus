package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	localTime := time.Now()
	fmt.Printf("current time: %v\n", localTime.Round(0))

	remoteTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("Error getting remote time: %v", err)
	}

	fmt.Printf("exact time: %v\n", remoteTime.Round(0))
}

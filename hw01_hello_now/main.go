package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

// Const ...
const (
	CustomFormat = "2006-01-02 15:04:05 -0700 MST"
	AppleNTPAddr = "time.apple.com"
)

func main() {
	localTime := time.Now().Format(CustomFormat)
	ntpTime, err := ntp.Time(AppleNTPAddr)
	if err != nil {
		log.Fatalf("Error while getting time from %s: %s", AppleNTPAddr, err)
	}
	roundedNtpTime := ntpTime.Round(ntpTime.Sub(ntpTime)).Format(CustomFormat)
	fmt.Printf("current time: %v\nexact time: %v\n", localTime, roundedNtpTime)
}

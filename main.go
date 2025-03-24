package main

import (
	"fmt"
	"time"
)

func main() {
	duration := 15 * time.Minute
	fmt.Println("15 minute pomodoro gabut started.")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for remaining := duration; remaining > 0; remaining -= time.Second {
		<-ticker.C
		fmt.Printf("\rTime left: %02d:%02d", int(remaining.Minutes()), int(remaining.Seconds())%60)
	}

	fmt.Println("\nWaktu gabut telah habis, KERJA KERJA KERJAAAA!!!")
}

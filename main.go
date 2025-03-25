package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	userInput, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	duration := time.Duration(userInput) * time.Minute

	startText := fmt.Sprintf("%v minute pomodoro gabut started", duration)
	fmt.Println(startText)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for remaining := duration; remaining > 0; remaining -= time.Second {
		<-ticker.C
		fmt.Printf("\rTime left: %02d:%02d", int(remaining.Minutes()), int(remaining.Seconds())%60)
	}

	fmt.Println("\nWaktu gabut telah habis, KERJA KERJA KERJAAAA!!!")
}

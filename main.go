package main

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("@every 1s", func() { log.Println("every 1s") })
	c.AddFunc("@every 5s", SampleFunction)

	c.AddFunc("* * * * *", func() { log.Println("every minute") })
	c.Start()
	time.Sleep(2 * time.Minute)
}

func SampleFunction() {
	fmt.Println("Sample Function")
}

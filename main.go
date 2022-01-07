package main

import (
	"fmt"
	"log"

	"example.com/m/v2/Bye"
	"github.com/robfig/cron"
)

func main() {
	ByeString := Bye.GetByeString()
	fmt.Printf("%s\n\n", ByeString)

	log.Print("Init Cron... ")
	c := cron.New()
	c.AddFunc("	0 0 19 * * ?", func() {
		Bye.ShutDown()
		fmt.Println("Bye Bye~")
	})
	log.Print("OK")
	log.Print("Shutdown at 7:00 pm")

	NotStop()
}

func NotStop() {
	for {
		fmt.Scanln()
	}
}

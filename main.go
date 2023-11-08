package main

import (
	"SongAlgo/selenium"
	"log"
)

func main() {
	err := run()
	if err != nil {
		log.Printf("Error: %+v", err)
	}
	selenium.KillChromeDriver()
}

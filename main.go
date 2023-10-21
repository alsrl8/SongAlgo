package main

import (
	"log"
)

func main() {
	err := run()
	if err != nil {
		log.Printf("Error: %+v", err)
		return
	}
}

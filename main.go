package main

import (
	"SongAlgo/selenium"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func initializeWebDriver() *selenium.WebDriverInstance {

	return selenium.GetWebDriverInstance()
}

func main() {
	driver := initializeWebDriver()
	defer func() {
		log.Println("Close the Selenium")
		driver.Close()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Caught Ctrl+C or SIGTERM. Closing Selenium...")
		driver.Close()
		os.Exit(0)
	}()

	err := run()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

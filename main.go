package main

import (
	"SongAlgo/selenium"
	"fmt"
)

func initializeWebDriver() *selenium.WebDriverInstance {

	return selenium.GetWebDriverInstance()
}

func main() {
	driver := initializeWebDriver()
	defer func() {
		driver.Close()
	}()

	err := run()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

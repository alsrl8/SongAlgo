package selenium

import (
	"github.com/tebeka/selenium"
	"log"
	"time"
)

// getChromeDriverPath Get Chrome web driver path.
// It is only tested on Windows environment.
func getChromeDriverPath() (chromeDriverPath string) {
	chromeDriverPath = "./selenium/driver/chromedriver"
	return
}

// getChromeDriverService Get Chrome web driver service from chrome selenium.
func getChromeDriverService() (*selenium.Service, error) {
	var opts []selenium.ServiceOption
	chromeDriverPath := getChromeDriverPath()
	service, err := selenium.NewChromeDriverService(chromeDriverPath, 4444, opts...)
	if err != nil {
		log.Println(err)
	}
	return service, nil
}

// OpenPageWithWebDriver Open web page with given url and web driver
func OpenPageWithWebDriver(wd *selenium.WebDriver, url string) error {
	err := (*wd).Get(url)
	if err != nil {
		return err
	}
	return nil
}

// waitUntilUserCloseBrowser Keep selenium browser awake until user closes the browser.
func waitUntilUserCloseBrowser(wd *selenium.WebDriver) {
	c := make(chan bool)
	go monitorBrowserClose(wd, c)
	<-c
}

// monitorBrowserClose Keep looking at if browser is alive.
func monitorBrowserClose(wd *selenium.WebDriver, c chan bool) {
	for {
		time.Sleep(1 * time.Second) // Poll every second
		_, err := (*wd).Title()
		if err != nil {
			c <- true
			break
		}
	}
	log.Printf("Browser was closed by the user.")
}

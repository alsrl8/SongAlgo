package selenium

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"time"
)

// resourceManager Manges resources for chrome selenium.
type resourceManager struct {
	service *selenium.Service
	caps    selenium.Capabilities
	wd      *selenium.WebDriver
}

// newResourceManager Generate new resource manager.
// It refers to user data directory of Chrome browser.
// It requires to call CleanUp function after using it.
func newResourceManager() (*resourceManager, error) {
	rm := &resourceManager{}

	service, err := getChromeDriverService()
	if err != nil {
		return nil, err
	}
	rm.service = service

	rm.caps = selenium.Capabilities{"browserName": "chrome"}
	userDataDir := "C:/Users/alsrl/AppData/Local/Google/Chrome/User Data"
	chromeCaps := chrome.Capabilities{
		Prefs: map[string]interface{}{
			"profile.default_content_settings.popups": 0,
		},
		Args: []string{
			"--headless",
			"--user-data-dir=" + userDataDir,
		},
		ExcludeSwitches: []string{
			"enable-logging",
		},
	}
	rm.caps.AddChrome(chromeCaps)

	webDriver, err := selenium.NewRemote(rm.caps, "")
	if err != nil {
		return nil, err
	}
	rm.wd = &webDriver

	return rm, nil
}

// cleanupResourceManager Clean resource manager and print error if it exists.
func cleanupResourceManager(rm *resourceManager) {
	if rm == nil {
		return
	}
	if err := rm.Cleanup(); err != nil {
		log.Printf("Failed to clean up resource manager: %v", err)
	}
}

// Cleanup Clear resource manager.
// It must be called after using it.
func (rm *resourceManager) Cleanup() error {
	err := (*rm.wd).Quit()
	if err != nil {
		return err
	}

	err = rm.service.Stop()
	if err != nil {
		return err
	}

	return nil
}

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
		fmt.Println(err)
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

// NavigateToPageWithCookieAndWait Navigate to page with given url and chrome cookie data.
func NavigateToPageWithCookieAndWait(url string) {

	rm, err := newResourceManager()
	if err != nil {
		log.Printf("Failed to create new resource manager: %v", err)
		return
	}
	defer cleanupResourceManager(rm)

	err = OpenPageWithWebDriver(rm.wd, url)
	if err != nil {
		log.Printf("Failed to access to url(%s): %v", url, err)
		return
	}

	waitUntilUserCloseBrowser(rm.wd)
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

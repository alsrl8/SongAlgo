package selenium

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
)

// resourceManager Manges resources for chrome selenium
type resourceManager struct {
	service *selenium.Service
	caps    selenium.Capabilities
	wd      *selenium.WebDriver
}

// newResourceManager Generate new resource manager
// It requires to call CleanUp function after using it
func newResourceManager() (*resourceManager, error) {
	rm := &resourceManager{}

	service, err := getChromeDriverService()
	if err != nil {
		return nil, err
	}
	rm.service = service

	rm.caps = selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Prefs: map[string]interface{}{
			"profile.default_content_settings.popups": 0,
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

func cleanupResourceManager(rm *resourceManager) {
	if rm == nil {
		return
	}
	if err := rm.Cleanup(); err != nil {
		log.Printf("Failed to clean up resource manager: %v", err)
	}
}

// Cleanup Clear resource manager.
// It must be called after using it
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

// getChromeDriverPath Get Chrome web driver path
// It is only tested on Windows environment
func getChromeDriverPath() (chromeDriverPath string) {
	chromeDriverPath = "./selenium/driver/chromedriver"
	return
}

// getCookieDataPath Get Cookie data json file path.
func getCookieDataPath() string {
	return "./selenium/cookie/cookies.json"
}

// getChromeDriverService Get Chrome web driver service from chrome selenium
func getChromeDriverService() (*selenium.Service, error) {
	var opts []selenium.ServiceOption
	chromeDriverPath := getChromeDriverPath()
	service, err := selenium.NewChromeDriverService(chromeDriverPath, 4444, opts...)
	if err != nil {
		fmt.Println(err)
	}
	return service, nil
}

// navigateToPage Open the page with given url and web driver
func navigateToPage(wd *selenium.WebDriver, url string) error {
	err := (*wd).Get(url)
	if err != nil {
		return err
	}
	return nil
}

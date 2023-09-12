package selenium

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"sync"
)

type WebDriverInstance struct {
	service *selenium.Service
	driver  *selenium.WebDriver
	mu      sync.Mutex
}

var instance *WebDriverInstance
var once sync.Once

func GetWebDriverInstance() *WebDriverInstance {
	once.Do(func() {
		service, err := getChromeDriverService()
		if err != nil {
			log.Fatalf("Failed to activate web driver service: %v", err)
		}

		caps := selenium.Capabilities{"browserName": "chrome"}
		userDataDir := "C:/Users/alsrl/AppData/Local/Google/Chrome/User Data"
		chromeCaps := chrome.Capabilities{
			Prefs: map[string]interface{}{
				"profile.default_content_settings.popups": 0,
			},
			Args: []string{
				//"--headless",
				"--user-data-dir=" + userDataDir,
			},
			ExcludeSwitches: []string{
				"enable-logging",
			},
		}
		caps.AddChrome(chromeCaps)

		webDriver, err := selenium.NewRemote(caps, "")
		if err != nil {
			log.Fatalf("Failed to open web driver session: %v", err)
		}

		instance = &WebDriverInstance{service: service, driver: &webDriver}
	})

	return instance
}

func (wd *WebDriverInstance) Close() {
	wd.mu.Lock()
	defer wd.mu.Unlock()

	if wd.driver != nil {
		err := (*wd.driver).Quit()
		if err != nil {
			log.Fatalf("Failed to quit web drvier")
		}
		wd.driver = nil
	}

	if wd.service != nil {
		err := (*wd.service).Stop()
		if err != nil {
			log.Fatalf("Failed to stop web driver service")
		}
		wd.service = nil
	}
}

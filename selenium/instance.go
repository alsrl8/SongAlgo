package selenium

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"sync"
)

type DriverManager struct {
	service *selenium.Service
	driver  *selenium.WebDriver
	mu      sync.Mutex
}

var instance *DriverManager
var once sync.Once

func GetWebDriverManager() *DriverManager {
	once.Do(func() {
		service, err := getChromeDriverService()
		if err != nil {
			log.Fatalf("Failed to activate web driver service: %v", err)
		}

		caps := selenium.Capabilities{"browserName": "chrome"}
		//userDataDir, err := getChromeUserDataDir()
		userDataDir, err := createChromeUserDataDir()
		if err != nil {
			log.Fatalf("Failed to get chrome user data dir: %v", err)
		}
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

		instance = &DriverManager{service: service, driver: &webDriver}
	})

	return instance
}

func (wd *DriverManager) Close() {
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

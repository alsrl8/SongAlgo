package selenium

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
)

type DriverManager struct {
	service *selenium.Service
	driver  *selenium.WebDriver
}

var instance *DriverManager

func IsDriverManagerRunning() bool {
	return instance != nil
}

func GetWebDriverManager(headless bool) *DriverManager {
	if IsDriverManagerRunning() {
		return instance
	}

	service, err := getChromeDriverService()
	if err != nil {
		log.Fatalf("Failed to activate web driver service: %v", err)
	}

	caps := selenium.Capabilities{"browserName": "chrome"}
	userDataDir, err := createChromeUserDataDir()
	if err != nil {
		log.Fatalf("Failed to get chrome user data dir: %v", err)
	}

	args := []string{"--user-data-dir=" + userDataDir}
	if headless {
		args = append(args, "--headless")
	}

	chromeCaps := chrome.Capabilities{
		Prefs: map[string]interface{}{
			"profile.default_content_settings.popups": 0,
		},
		Args: args,
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
	return instance
}

func (wd *DriverManager) Close() {
	if wd == nil {
		return
	}

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
	instance = nil
}

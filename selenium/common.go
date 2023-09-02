package selenium

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type resourceManager struct {
	service *selenium.Service
	caps    selenium.Capabilities
	wd      *selenium.WebDriver
}

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

func (rm *resourceManager) Cleanup() error {
	err := rm.service.Stop()
	if err != nil {
		return err
	}

	err = (*rm.wd).Quit()
	if err != nil {
		return err
	}

	return nil
}

func getChromeDriverPath() (chromeDriverPath string) {
	chromeDriverPath = "./selenium/driver/chromedriver"
	return
}

func getChromeDriverService() (*selenium.Service, error) {
	var opts []selenium.ServiceOption
	chromeDriverPath := getChromeDriverPath()
	service, err := selenium.NewChromeDriverService(chromeDriverPath, 4444, opts...)
	if err != nil {
		fmt.Println(err)
	}
	return service, nil
}

func getChromeCapabilities() *selenium.Capabilities {
	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Prefs: map[string]interface{}{
			"profile.default_content_settings.popups": 0,
		},
	}
	caps.AddChrome(chromeCaps)
	return &caps
}

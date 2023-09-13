package selenium

import (
	"fmt"
	"github.com/tebeka/selenium"
	"log"
	"time"
)

func findSubmitButton(wd *selenium.WebDriver) (selenium.WebElement, error) {
	button, err := (*wd).FindElement(selenium.ByID, "submit-code")
	if err != nil {
		return button, err
	}
	return button, err
}

func clickSubmitButton(button selenium.WebElement) error {
	err := button.Click()
	if err != nil {
		return err
	}
	return nil
}

func waitForSubmitResult(wd *selenium.WebDriver) error {
	err := (*wd).WaitWithTimeoutAndInterval(func(wd selenium.WebDriver) (bool, error) {
		element, err := wd.FindElement(selenium.ByClassName, "modal-title")
		if err != nil {
			return false, nil
		}

		text, err := element.Text()
		if err != nil {
			return false, nil
		}

		return text != "", nil
	}, 30*time.Second, 500*time.Millisecond)
	return err
}

func IsSubmittedCodeCorrect(url string) bool {
	webDriverInstance := GetWebDriverInstance()

	err := OpenPageWithWebDriver(webDriverInstance.driver, url)
	if err != nil {
		log.Printf("Failed to access to url(%s): %+v", url, err)
		return false
	}

	button, err := findSubmitButton(webDriverInstance.driver)
	if err != nil {
		log.Printf("Failed to find submit button: %+v", err)
		return false
	}

	err = clickSubmitButton(button)
	if err != nil {
		log.Printf("Failed to click submit button: %+v", err)
		return false
	}

	err = waitForSubmitResult(webDriverInstance.driver)
	if err != nil {
		log.Printf("Not Found or other error: %+v", err)
		return false
	}

	modalTitle, err := (*webDriverInstance.driver).FindElement(selenium.ByClassName, "modal-title")
	if err != nil {
		log.Printf("Failed to find modal title: %+v", err)
		return false
	}

	titleText, err := modalTitle.Text()
	if err != nil {
		log.Printf("Failed to extract text from modal title: %+v", err)
		return false
	}

	return titleText == "정답입니다!"
}

func UploadPgSourceToGithub(url string, problemTitle string, problemDate string) {
	webDriverInstance := GetWebDriverInstance()

	err := OpenPageWithWebDriver(webDriverInstance.driver, url)
	if err != nil {
		log.Printf("Failed to access to url(%s): %+v", url, err)
		return
	}

	codeElements := findPgSubmitCodeElements(webDriverInstance.driver)
	codes := extractCodeFromCodeElements(codeElements)
	fmt.Println(codes)
}

func findPgSubmitCodeElements(wd *selenium.WebDriver) []selenium.WebElement {
	codeMirror, err := (*wd).FindElement(selenium.ByClassName, "CodeMirror-code")
	if err != nil {
		log.Fatalf("Failed to find element: %v", err)
		return []selenium.WebElement{}
	}

	codeElements, err := codeMirror.FindElements(selenium.ByClassName, "CodeMirror-line")
	if err != nil {
		log.Printf("Error finding elements: %v", err)
		return []selenium.WebElement{}
	}

	return codeElements
}

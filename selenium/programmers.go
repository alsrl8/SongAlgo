package selenium

import (
	"fmt"
	"github.com/tebeka/selenium"
	"log"
)

func NavigateToPgProblemWithCookie(url string) {
	webDriverInstance := GetWebDriverInstance()

	err := OpenPageWithWebDriver(webDriverInstance.driver, url)
	if err != nil {
		log.Printf("Failed to access to url(%s): %v", url, err)
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

package selenium

import (
	"SongAlgo/github"
	"fmt"
	"github.com/tebeka/selenium"
	"log"
	"os"
	"time"
)

type PgSourceData struct {
	Code      string `json:"code"`
	Extension string `json:"extension"`
}

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

func UploadPgSourceToGithub(problemTitle string, problemDate string, githubId string, code string, extension string) {
	dateString := convertDateString(problemDate)
	date := time.Now().Format("060102")
	params := github.UploadParams{
		Token:   os.Getenv("GITHUB_TOKEN"),
		Owner:   github.GetRepositoryOwner(),
		Repo:    github.GetRepositoryName(),
		Path:    fmt.Sprintf("%s/%s.%s", dateString, problemTitle, extension),
		Branch:  githubId,
		Message: date,
		Content: code,
		//Sha:     sha,  // TODO 프로그래머스 코드 제출 시 덮어쓰기 기능을 위해 SHA를 같이 전달하도록 수정
	}
	err := github.UploadFileToGithub(params)
	if err != nil {
		log.Printf("Error occured during uploading file to github: %+v", err)
		return
	}
}

func findPgSubmitCodeElements(wd *selenium.WebDriver) []selenium.WebElement {
	codeMirror, err := (*wd).FindElement(selenium.ByClassName, "CodeMirror-code")
	if err != nil {
		log.Printf("Failed to find element: %v", err)
		return []selenium.WebElement{}
	}

	codeElements, err := codeMirror.FindElements(selenium.ByClassName, "CodeMirror-line")
	if err != nil {
		log.Printf("Error finding elements: %v", err)
		return []selenium.WebElement{}
	}

	return codeElements
}

func findPgLanguageElement(wd *selenium.WebDriver) selenium.WebElement {
	element, err := (*wd).FindElement(selenium.ByCSSSelector, `.btn-dark.dropdown-toggle`)
	if err != nil {
		log.Printf("Failed to find element: %+v", err)
		return nil
	}

	return element
}

func extractLanguageFromLanguageElement(languageElement selenium.WebElement) string {
	language, err := languageElement.Text()
	if err != nil {
		log.Printf("Failed to extract language from html element: %+v", err)
		return ""
	}
	return language
}

func GetPgSourceData(url string) PgSourceData {
	webDriverInstance := GetWebDriverInstance()
	err := OpenPageWithWebDriver(webDriverInstance.driver, url)
	if err != nil {
		log.Printf("Failed to access to url(%s): %+v", url, err)
		return PgSourceData{}
	}

	codeElements := findPgSubmitCodeElements(webDriverInstance.driver)
	code := extractCodeFromCodeElements(codeElements)
	languageElement := findPgLanguageElement(webDriverInstance.driver)
	language := extractLanguageFromLanguageElement(languageElement)
	extension := convertCodeLanguageToFileExtension(language)

	return PgSourceData{
		Code:      code,
		Extension: extension,
	}
}

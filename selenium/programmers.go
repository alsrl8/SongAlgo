package selenium

import (
	"SongAlgo/github"
	"fmt"
	"github.com/tebeka/selenium"
	"log"
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

func IsPgLoggedIn(url string) bool {
	dm := GetWebDriverManager(true)

	err := OpenPageWithWebDriver(dm.driver, url)
	if err != nil {
		log.Printf("Failed to access to url(%s): %+v", url, err)
		return false
	}

	_, err = findSubmitButton(dm.driver)
	if err != nil {
		return false
	}
	return true
}

func IsSubmittedCodeCorrect(url string) bool {
	dm := GetWebDriverManager(true)

	err := OpenPageWithWebDriver(dm.driver, url)
	if err != nil {
		log.Printf("Failed to access to url(%s): %+v", url, err)
		return false
	}

	button, err := findSubmitButton(dm.driver)
	if err != nil {
		log.Printf("Failed to find submit button: %+v", err)
		return false
	}

	err = clickSubmitButton(button)
	if err != nil {
		log.Printf("Failed to click submit button: %+v", err)
		return false
	}

	err = waitForSubmitResult(dm.driver)
	if err != nil {
		log.Printf("Not Found or other error: %+v", err)
		return false
	}

	modalTitle, err := (*dm.driver).FindElement(selenium.ByClassName, "modal-title")
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

func UploadPgSourceToGithub(problemTitle string, problemDate string, githubId string, code string, extension string, sha string) {
	dateString := convertDateString(problemDate)
	date := time.Now().Format("060102")
	params := github.UploadParams{
		Token:   github.GetRepositoryToken(),
		Owner:   github.GetRepositoryOwner(),
		Repo:    github.GetRepositoryName(),
		Path:    fmt.Sprintf("%s/%s.%s", dateString, problemTitle, extension),
		Branch:  githubId,
		Message: date,
		Content: code,
		Sha:     sha,
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
	dm := GetWebDriverManager(true)
	err := OpenPageWithWebDriver(dm.driver, url)
	if err != nil {
		log.Printf("Failed to access to url(%s): %+v", url, err)
		return PgSourceData{}
	}

	codeElements := findPgSubmitCodeElements(dm.driver)
	code := extractCodeFromCodeElements(codeElements)
	languageElement := findPgLanguageElement(dm.driver)
	language := extractLanguageFromLanguageElement(languageElement)
	extension := convertCodeLanguageToFileExtension(language)

	return PgSourceData{
		Code:      code,
		Extension: extension,
	}
}

func GetGithubRepositoryPgSource(problemTitle string, problemDate string, githubId string, extension string) (github.File, error) {
	branchName := githubId
	dateString := convertDateString(problemDate)
	path := fmt.Sprintf("%s/%s.%s", dateString, problemTitle, extension)
	log.Printf("Getting github source from path(%+v)", path)
	return github.GetGithubRepositorySource(branchName, path)
}

func NavigateToPgLoginPage() {
	dm := GetWebDriverManager(false)
	loginPage := "https://programmers.co.kr/account/sign_in?referer=https%3A%2F%2Fprogrammers.co.kr%2F"
	err := (*dm.driver).Get(loginPage)
	if err != nil {
		log.Printf("Failed to access to url(%s): %+v", loginPage, err)
	}
}

func crawlPgProblem(url string) *github.Problem {
	return nil
}

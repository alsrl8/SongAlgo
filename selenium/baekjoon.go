package selenium

import (
	"SongAlgo/github"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tebeka/selenium"
	"log"
	"os"
	"strings"
	"time"
)

type SubmitHistory struct {
	SubmissionNumber string `json:"SubmissionNumber"`
	ID               string `json:"ID"`
	Problem          string `json:"Problem"`
	Result           string `json:"Result"`
	Memory           string `json:"Memory"`
	Time             string `json:"Time"`
	Language         string `json:"Language"`
	CodeLength       string `json:"CodeLength"`
	SubmissionTime   string `json:"SubmissionTime"`
}

func navigateToBjSourcePage(wd *selenium.WebDriver, submissionNumber string) {
	url := fmt.Sprintf("https://www.acmicpc.net/source/%s", submissionNumber)
	err := OpenPageWithWebDriver(wd, url)
	if err != nil {
		log.Fatalf("Failed to navigate to submission source page")
		return
	}
}

func navigateToBjSubmitHistoryPage(wd *selenium.WebDriver) {
	elem, err := (*wd).FindElement(selenium.ByLinkText, "내 제출")
	if err != nil {
		log.Fatalf("Failed to find element: %+v", err)
	}

	href, err := elem.GetAttribute("href")
	if err != nil {
		log.Fatalf("Failed to get href attribute: %v", err)
	}

	err = OpenPageWithWebDriver(wd, href)
	if err != nil {
		log.Printf("Failed to access to url(%s): %v", href, err)
		return
	}
}

func findBjSubmitHistoryTrElements(wd *selenium.WebDriver) []selenium.WebElement {
	table, err := (*wd).FindElement(selenium.ByID, "status-table")
	if err != nil {
		log.Fatalf("Failed to find table: %v", err)
	}

	trElements, err := table.FindElements(selenium.ByCSSSelector, "tbody tr")
	if err != nil {
		log.Fatalf("Failed to find th elements: %v", err)
	}

	return trElements
}

func bindSubmitHistories(trElements []selenium.WebElement) []SubmitHistory {
	var histories []SubmitHistory
	for _, tr := range trElements {
		tdElements, err := tr.FindElements(selenium.ByCSSSelector, "td")
		if err != nil {
			continue
		}
		submitHistory := SubmitHistory{}
		for i, tdElement := range tdElements {
			text, err := tdElement.Text()
			if err != nil {
				err = errors.Wrap(err, "failed to extract submit history from html tags")
				fmt.Printf("%+v", err)
				continue
			}

			switch i {
			case 0:
				submitHistory.SubmissionNumber = text
			case 1:
				submitHistory.ID = strings.Trim(text, " ")
			case 2:
				submitHistory.Problem = text
			case 3:
				submitHistory.Result = text
			case 4:
				submitHistory.Memory = text
			case 5:
				submitHistory.Time = text
			case 6:
				elementLanguage := strings.Replace(text, " / 수정", "", 1)
				submitHistory.Language = elementLanguage
			case 7:
				submitHistory.CodeLength = text
			case 8:
				submitHistory.SubmissionTime = text
			}
		}
		histories = append(histories, submitHistory)
	}
	return histories
}

func findBjSubmitHistories(wd *selenium.WebDriver) []SubmitHistory {
	navigateToBjSubmitHistoryPage(wd)
	trElements := findBjSubmitHistoryTrElements(wd)
	submitHistories := bindSubmitHistories(trElements)
	return submitHistories
}

func NavigateToBjProblemWithCookie(url string) []SubmitHistory {
	wd := GetWebDriverInstance()

	err := OpenPageWithWebDriver(wd.driver, url)
	if err != nil {
		log.Printf("Failed to access to url(%s): %v", url, err)
		return []SubmitHistory{}
	}

	submitHistories := findBjSubmitHistories(wd.driver)
	return submitHistories
}

func UploadBjSourceToGithub(problemTitle string, problemDate string, submission SubmitHistory, sha string) {
	webDriverInstance := GetWebDriverInstance()
	navigateToBjSourcePage(webDriverInstance.driver, submission.SubmissionNumber)
	codeElements := findBjSubmitCodeElements(webDriverInstance.driver)

	dateString := convertDateString(problemDate)
	extension := convertCodeLanguageToFileExtension(submission.Language)
	githubId := convertBjIdToGithubId(submission.ID)
	date := time.Now().Format("060102")
	code := extractCodeFromCodeElements(codeElements)

	params := github.UploadParams{
		Token:   os.Getenv("GITHUB_TOKEN"),
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

func findBjSubmitCodeElements(wd *selenium.WebDriver) []selenium.WebElement {
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

func convertBjIdToGithubId(bjId string) (githubId string) {
	switch bjId {
	case "alsrl9":
		githubId = "alsrl8"
	}
	return
}

func GetGithubRepositoryBjSource(problemTitle string, problemDate string, bjId string, language string) (github.File, error) {
	dateString := convertDateString(problemDate)
	branchName := convertBjIdToGithubId(bjId)
	extension := convertCodeLanguageToFileExtension(language)
	path := fmt.Sprintf("%s/%s.%s", dateString, problemTitle, extension)
	return github.GetGithubRepositorySource(branchName, path)
}

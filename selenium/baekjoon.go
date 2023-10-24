package selenium

import (
	"SongAlgo/github"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tebeka/selenium"
	"log"
	"path"
	"strconv"
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

func IsBjLoggedIn(url string) bool {
	dm := GetWebDriverManager(true)

	err := OpenPageWithWebDriver(dm.driver, url)
	if err != nil {
		log.Printf("Failed to access to url(%s): %v", url, err)
		return false
	}

	_, err = (*dm.driver).FindElement(selenium.ByLinkText, "내 제출")
	if err != nil {
		return false
	}
	return true
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
	dm := GetWebDriverManager(true)

	err := OpenPageWithWebDriver(dm.driver, url)
	if err != nil {
		log.Printf("Failed to access to url(%s): %v", url, err)
		return []SubmitHistory{}
	}

	submitHistories := findBjSubmitHistories(dm.driver)
	return submitHistories
}

func UploadBjSourceToGithub(problemTitle string, problemDate string, submission SubmitHistory, sha string) {
	dm := GetWebDriverManager(true)
	navigateToBjSourcePage(dm.driver, submission.SubmissionNumber)
	codeElements := findBjSubmitCodeElements(dm.driver)

	dateString := convertDateString(problemDate)
	extension := convertCodeLanguageToFileExtension(submission.Language)
	githubId := convertBjIdToGithubId(submission.ID)
	date := time.Now().Format("060102")
	code := extractCodeFromCodeElements(codeElements)

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
	pathStr := fmt.Sprintf("%s/%s.%s", dateString, problemTitle, extension)
	return github.GetGithubRepositorySource(branchName, pathStr)
}

func NavigateToBjLoginPage() {
	dm := GetWebDriverManager(false)
	loginPage := "https://www.acmicpc.net/login?next=%2Fproblem%2F1000"
	err := (*dm.driver).Get(loginPage)
	if err != nil {
		log.Printf("Failed to access to url(%s): %+v", loginPage, err)
	}
}

func crawlBjProblem(url string) *github.Problem {
	dm := GetWebDriverManager(true)
	err := (*dm.driver).Get(url)
	if err != nil {
		log.Printf("Failed to access to url(%s): %+v", url, err)
		return nil
	}

	problemName := getBjProblemNameFromHtmlSource(dm.driver)
	tier := getBjTierFromHtmlSource(dm.driver)

	return &github.Problem{
		Name:          problemName,
		AlgorithmType: "",
		Difficulty:    tier,
		Platform:      "baekjoon",
		Url:           url,
	}
}

func getBjProblemNameFromHtmlSource(wd *selenium.WebDriver) string {
	element, err := (*wd).FindElement(selenium.ByID, "problem_title")
	if err != nil {
		log.Printf("Failed to find problem title element: %v", err)
		return ""
	}

	problemName, err := element.Text()
	if err != nil {
		log.Printf("Failed to get problem title from element: %v", err)
		return ""
	}
	return problemName
}

func getBjTierFromHtmlSource(wd *selenium.WebDriver) string {
	infoElement, err := (*wd).FindElement(selenium.ByCSSSelector, "ul.nav.nav-pills.no-print.problem-menu")
	if err != nil {
		log.Printf("Failed to find info element: %v", err)
		return ""
	}

	imgElement, err := infoElement.FindElement(selenium.ByTagName, "img")
	if err != nil {
		log.Printf("Failed to find img element from info element: %v", err)
		return ""
	}

	src, err := imgElement.GetAttribute("src")
	if err != nil {
		log.Printf("Failed to get src from element: %v", err)
		return ""
	}

	tierImage := strings.TrimSuffix(path.Base(src), ".svg")
	tierImageNum, _ := strconv.Atoi(tierImage)
	tier := getTierByTierImageNum(tierImageNum)
	return tier
}

func getTierByTierImageNum(tierImageNum int) (tier string) {
	if tierImageNum <= 0 || tierImageNum > 31 {
		return ""
	}

	switch (tierImageNum - 1) / 5 {
	case 0: // Bronze
		tier += "B"
	case 1: // Silver
		tier += "S"
	case 2: // Gold
		tier += "G"
	case 3: // Platinum
		tier += "P"
	case 4: // Diamond
		tier += "D"
	case 5: // Ruby
		tier += "R"
	}

	tierNum := 5 - (tierImageNum-1)%5
	tier += strconv.Itoa(tierNum)
	return
}

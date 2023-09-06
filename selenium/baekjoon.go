package selenium

import (
	"github.com/tebeka/selenium"
	"log"
)

type SubmitHistory struct {
	SubmissionNumber string `json:"제출 번호"`
	ID               string `json:"아이디"`
	Problem          string `json:"문제"`
	Result           string `json:"결과"`
	Memory           string `json:"메모리"`
	Time             string `json:"시간"`
	Language         string `json:"언어"`
	CodeLength       string `json:"코드 길이"`
	SubmissionTime   string `json:"제출한 시간"`
}

func navigateToBjSubmitHistoryPage(wd *selenium.WebDriver) {
	elem, err := (*wd).FindElement(selenium.ByLinkText, "내 제출")
	if err != nil {
		log.Fatalf("Failed to find element: %v", err)
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
			switch i {
			case 0:
				submitHistory.SubmissionNumber, err = tdElement.Text()
				if err != nil {
					continue
				}
			case 1:
				submitHistory.ID, err = tdElement.Text()
				if err != nil {
					continue
				}
			case 2:
				submitHistory.Problem, err = tdElement.Text()
				if err != nil {
					continue
				}
			case 3:
				submitHistory.Result, err = tdElement.Text()
				if err != nil {
					continue
				}
			case 4:
				submitHistory.Memory, err = tdElement.Text()
				if err != nil {
					continue
				}
			case 5:
				submitHistory.Time, err = tdElement.Text()
				if err != nil {
					continue
				}
			case 6:
				submitHistory.Language, err = tdElement.Text()
				if err != nil {
					continue
				}
			case 7:
				submitHistory.CodeLength, err = tdElement.Text()
				if err != nil {
					continue
				}
			case 8:
				submitHistory.SubmissionTime, err = tdElement.Text()
				if err != nil {
					continue
				}
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
	rm, err := newResourceManager()
	if err != nil {
		log.Printf("Failed to create new resource manager: %v", err)
		return []SubmitHistory{}
	}
	defer cleanupResourceManager(rm)

	err = OpenPageWithWebDriver(rm.wd, url)
	if err != nil {
		log.Printf("Failed to access to url(%s): %v", url, err)
		return []SubmitHistory{}
	}

	submitHistories := findBjSubmitHistories(rm.wd)
	return submitHistories
}

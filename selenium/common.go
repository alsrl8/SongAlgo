package selenium

import (
	"SongAlgo/github"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tebeka/selenium"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// getChromeDriverPath Get Chrome web driver path.
// It is only tested on Windows environment.
func getChromeDriverPath() (chromeDriverPath string) {
	chromeDriverPath = "./selenium/driver/chromedriver"
	return
}

// getChromeDriverService Get Chrome web driver service from chrome selenium.
func getChromeDriverService() (*selenium.Service, error) {
	var opts []selenium.ServiceOption
	chromeDriverPath := getChromeDriverPath()
	service, err := selenium.NewChromeDriverService(chromeDriverPath, 4444, opts...)
	if err != nil {
		log.Println(err)
	}
	return service, nil
}

// OpenPageWithWebDriver Open web page with given url and web driver
func OpenPageWithWebDriver(wd *selenium.WebDriver, url string) error {
	err := (*wd).Get(url)
	if err != nil {
		return err
	}
	return nil
}

// waitUntilUserCloseBrowser Keep selenium browser awake until user closes the browser.
func waitUntilUserCloseBrowser(dm *DriverManager) {
	c := make(chan bool)
	go monitorBrowserClose(dm.driver, c)
	<-c
	dm.Close()
}

// monitorBrowserClose Keep looking at if browser is alive.
func monitorBrowserClose(wd *selenium.WebDriver, c chan bool) {
	for {
		time.Sleep(100 * time.Millisecond) // Poll every second
		_, err := (*wd).Title()
		if err != nil {
			c <- true
			break
		}
	}
	log.Printf("Browser was closed by the user.")
}

// extractCodeFromCodeElements CodeMirror HTML Tag에서 Code를 한 줄씩 읽어 string으로 반환한다.
func extractCodeFromCodeElements(codeElements []selenium.WebElement) string {
	var codes []string
	for _, ce := range codeElements {
		text, err := ce.Text()
		if err != nil {
			log.Printf("Erorr getting text: %v", err)
			continue
		}
		codes = append(codes, text)
	}
	return strings.Join(codes, "\n")
}

// convertDateString "YYYY-MM-DD" 형식의 date string을 "YYMMDD"로 반환한다.
func convertDateString(dateStr string) string {
	parts := strings.Split(dateStr, "-")
	return parts[0][2:] + parts[1] + parts[2]
}

// convertCodeLanguageToFileExtension 언어 이름을 매개 변수로 받고 확장자를 반환한다.
func convertCodeLanguageToFileExtension(language string) (extension string) {
	language = strings.Trim(language, " ")
	switch language {
	case "Python3", "PyPy3", "Python 3", "Python2":
		return "py"
	case "C90", "C99", "C11":
		return "c"
	case "Java", "Java 8", "Java 8 (OpenJDK)", "Java 11", "Java 15":
		return "java"
	case "C++", "C++98", "C++11", "C++14", "C++17":
		return "cpp"
	case "Go":
		return "go"
	default:
		return language
	}
}

func getChromeUserDataDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Construct the path to the Chrome User Data directory based on the OS.
	var userDataDir string
	switch runtime.GOOS {
	case "windows":
		userDataDir = filepath.Join(home, "AppData", "Local", "Google", "Chrome", "User Data")
	case "darwin":
		userDataDir = filepath.Join(home, "Library", "Application Support", "Google", "Chrome")
	case "linux":
		userDataDir = filepath.Join(home, ".config", "google-chrome")
	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	return userDataDir, nil
}

func createChromeUserDataDir() (string, error) {
	// Get the home directory for the current user.
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Define the path for the custom profile directory.
	var customProfileDir string
	switch runtime.GOOS {
	case "windows":
		customProfileDir = filepath.Join(home, "AppData", "Local", "SongAlgo", "User Data")
	case "darwin":
		customProfileDir = filepath.Join(home, "Library", "Application Support", "SongAlgo")
	case "linux":
		customProfileDir = filepath.Join(home, ".config", "SongAlgo")
	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	// Create the custom profile directory if it doesn't exist.
	err = os.MkdirAll(customProfileDir, 0755)
	if err != nil {
		return "", err
	}

	return customProfileDir, nil
}

func AddProblem(username string, date string, problemUrl1 string, problemUrl2 string, problemUrl3 string) {
	if username != github.GetRepositoryOwner() {
		log.Printf("User name(%s) is invalid", username)
		return
	}

	problemUrlList := []string{problemUrl1, problemUrl2, problemUrl3}
	var problemList []github.Problem
	for _, problemUrl := range problemUrlList {
		parsedUrl, err := url.Parse(problemUrl)
		if err != nil {
			log.Printf("Failed to parsing url(%s): %+v", problemUrl, err)
			return
		}

		var problem *github.Problem
		switch parsedUrl.Hostname() {
		case "www.acmicpc.net":
			problem = crawlBjProblem(parsedUrl.String())
		case "school.programmers.co.kr":
			problem = crawlPgProblem(parsedUrl.String())
		}

		if problem == nil {
			log.Printf("Failed to crawl problem from given url(%s)", problemUrl)
			return
		}
		problemList = append(problemList, *problem)
	}

	params := github.GetParams{
		Token:  github.GetRepositoryToken(),
		Owner:  github.GetRepositoryOwner(),
		Repo:   github.GetRepositoryName(),
		Branch: github.GetScheduleBranchName(),
		Path:   github.GetScheduleFileName(),
	}
	file, _ := github.GetGithubRepositoryContent(params)

	scheduleList := github.ScheduleList{}
	decodedContent, err := base64.StdEncoding.DecodeString(file.Content)
	if err != nil {
		log.Printf("Failed to decode file content from github")
		return
	}
	err = json.Unmarshal(decodedContent, &scheduleList)
	if err != nil {
		log.Printf("Failed to fetch schedule list from github")
		return
	}
	scheduleList.List = append(scheduleList.List, github.Schedule{
		Date:     date,
		Problems: problemList,
	})
	jsonData, _ := json.MarshalIndent(scheduleList, "", "    ")

	uploadParams := github.UploadParams{
		Token:     github.GetRepositoryToken(),
		Owner:     github.GetRepositoryOwner(),
		Committer: github.GetRepositoryOwner(),
		Repo:      github.GetRepositoryName(),
		Path:      "Schedule.json",
		Branch:    "schedule",
		Message:   date,
		Content:   string(jsonData),
		Sha:       file.Sha,
	}

	err = github.UploadFileToGithub(uploadParams)
	if err != nil {
		log.Printf("Error occured during uploading file to github: %+v", err)
		return
	}
}

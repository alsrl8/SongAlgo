package main

import (
	"SongAlgo/github"
	"SongAlgo/selenium"
	"SongAlgo/util"
	"context"
	"os"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts.
// The context is saved so we can call the runtime methods
func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
}

func (app *App) GetMenu() [3]string {
	return [3]string{"1. 문제 리스트", "2. 문제 추가하기", "3. 종료"}
}

func (app *App) GetSchedule() *github.ScheduleList {
	scheduleList, err := github.FetchScheduleListFromGitHub()
	if err != nil {
		return nil
	}
	return scheduleList
}

func (app *App) IsChromeRunning() bool {
	return util.IsChromeRunning()
}

func (app *App) NavigateToBjProblemWithCookie(url string) []selenium.SubmitHistory {
	return selenium.NavigateToBjProblemWithCookie(url)
}

func (app *App) UploadBjSourceToGithub(problemTitle string, problemDate string, submission selenium.SubmitHistory, sha string, userName string) {
	selenium.UploadBjSourceToGithub(problemTitle, problemDate, submission, sha, userName)
}

func (app *App) GetGithubRepositoryBjSource(problemTitle string, problemDate string, userName string, language string) github.FileResponse {
	file, err := selenium.GetGithubRepositoryBjSource(problemTitle, problemDate, userName, language)
	fileResponse := github.ConvertGithubRepositoryFileToFileResponse(file, err)
	return fileResponse
}

func (app *App) IsSubmittedCodeCorrect(url string) bool {
	return selenium.IsSubmittedCodeCorrect(url)
}

func (app *App) UploadPgSourceToGithub(problemTitle string, problemDate string, githubId string, code string, extension string, sha string) {
	selenium.UploadPgSourceToGithub(problemTitle, problemDate, githubId, code, extension, sha)
}

func (app *App) IsBjLoggedIn(url string) bool {
	return selenium.IsBjLoggedIn(url)
}

func (app *App) IsPgLoggedIn(url string) bool {
	return selenium.IsPgLoggedIn(url)
}

func (app *App) GetPgSourceData(url string) selenium.PgSourceData {
	return selenium.GetPgSourceData(url)
}

func (app *App) GetGithubRepositoryPgSource(problemTitle string, problemDate string, githubId string, extension string) github.FileResponse {
	file, err := selenium.GetGithubRepositoryPgSource(problemTitle, problemDate, githubId, extension)
	fileResponse := github.ConvertGithubRepositoryFileToFileResponse(file, err)
	return fileResponse
}

func (app *App) CloseSeleniumBrowser() {
	if selenium.IsDriverManagerRunning() == false {
		return
	}
	manager := selenium.GetWebDriverManager(false)
	manager.Close()
}

func (app *App) NavigateToPgLoginPage() {
	selenium.NavigateToPgLoginPage()
}

func (app *App) NavigateToBjLoginPage() {
	selenium.NavigateToBjLoginPage()
}

func (app *App) CloseProgram() {
	os.Exit(0)
}

func (app *App) AddProblem(username string, date string, problemUrl1 string, problemUrl2 string, problemUrl3 string) {
	selenium.AddProblem(username, date, problemUrl1, problemUrl2, problemUrl3)
}

package main

import (
	"SongAlgo/github"
	"SongAlgo/selenium"
	"SongAlgo/util"
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
}

func (app *App) GetMenu() [2]string {
	return [2]string{"1. 오늘의 문제 확인", "2. 종료"}
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

func (app *App) NavigateToPageWithCookie(url string) {
	selenium.NavigateToPageWithCookieAndWait(url)
}

func (app *App) NavigateToBjProblemWithCookie(url string) []selenium.SubmitHistory {
	return selenium.NavigateToBjProblemWithCookie(url)
}

func (app *App) UploadBjSourceToGithub(SubmissionNumber string) {
	selenium.UploadBjSourceToGithub(SubmissionNumber)
}

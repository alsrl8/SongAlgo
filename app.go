package main

import (
	"SongAlgo/github"
	"SongAlgo/selenium"
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

func (app *App) GenerateCookieForBJ() selenium.Cookie {
	return selenium.GetCookieForBJ()
}

func (app *App) ReadCookieForBJ() selenium.Cookie {
	return selenium.ReadCookieForBJ()
}

func (app *App) OpenBjWithCookie(cookie string) {
	// TODO 저장된 쿠키로 백준 사이트를 연다.
}

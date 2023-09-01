package main

import (
	"SongAlgo/github"
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

func (app *App) GetSchedule() *github.Schedule {
	schedule, err := github.FetchScheduleFromGitHub()
	if err != nil {
		return nil
	}
	return schedule
}

package github

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Problem struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Url      string `json:"url"`
}

type Schedule struct {
	Date     string    `json:"date"`
	Problems []Problem `json:"problems"`
}

func FetchScheduleFromGitHub() (*Schedule, error) {
	// URL to the raw version of the file in the GitHub repository
	url := "https://raw.githubusercontent.com/alsrl8/SongAlgo/schedule/Schedule.json"

	// Fetch the file from GitHub
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON into the Schedule struct
	var schedule Schedule
	err = json.Unmarshal(body, &schedule)
	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

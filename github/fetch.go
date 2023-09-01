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

type ScheduleList struct {
	List []Schedule `json:"list"`
}

func FetchScheduleListFromGitHub() (*ScheduleList, error) {
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
	var scheduleList ScheduleList
	err = json.Unmarshal(body, &scheduleList)
	if err != nil {
		return nil, err
	}

	return &scheduleList, nil
}

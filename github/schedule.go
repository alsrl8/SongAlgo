package github

import (
	"encoding/json"
)

type Problem struct {
	Name          string `json:"name"`
	AlgorithmType string `json:"algorithmType"`
	Difficulty    string `json:"difficulty"`
	Platform      string `json:"platform"`
	Url           string `json:"url"`
}

type Schedule struct {
	Date     string    `json:"date"`
	Problems []Problem `json:"problems"`
}

type ScheduleList struct {
	List []Schedule `json:"list"`
}

func FetchScheduleListFromGitHub() (*ScheduleList, error) {
	fetchParams := FetchParams{
		Owner:  "alsrl8",
		Repo:   "SongAlgo",
		Branch: "schedule",
		Path:   "Schedule.json",
	}
	fetchData, err := FetchFromGithub(fetchParams)
	if err != nil {
		return nil, err
	}

	var scheduleList ScheduleList
	err = json.Unmarshal(fetchData, &scheduleList)
	if err != nil {
		return nil, err
	}

	return &scheduleList, nil
}

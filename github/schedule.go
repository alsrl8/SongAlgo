package github

import (
	"encoding/json"
	"log"
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
	log.Printf("Fetching schedule list from github repository...")
	fetchParams := FetchParams{
		Owner:  GetRepositoryOwner(),
		Repo:   GetRepositoryName(),
		Branch: GetScheduleBranchName(),
		Path:   GetScheduleFileName(),
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

	// Reverse schedule list
	n := len(scheduleList.List)
	for i := 0; i < n/2; i++ {
		scheduleList.List[i], scheduleList.List[n-1-i] = scheduleList.List[n-1-i], scheduleList.List[i]
	}

	return &scheduleList, nil
}

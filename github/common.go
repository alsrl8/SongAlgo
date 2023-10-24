package github

import (
	"gopkg.in/ini.v1"
	"log"
)

func GetRepositoryToken() string {
	cfg, err := ini.Load("./github/config.ini")
	if err != nil {
		log.Printf("Failed to read github config ini file. %+v", err)
	}
	token := cfg.Section("token").Key("fine-grained-access-token").String()
	return token
}

func GetRepositoryOwner() string {
	return "alsrl8"
}

func GetRepositoryName() string {
	return "SongAlgo"
}

func GetScheduleBranchName() string {
	return "schedule"
}

func GetScheduleFileName() string {
	return "Schedule.json"
}

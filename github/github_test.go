package github

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestUploadFileToGithub(t *testing.T) {
	params := UploadParams{
		Token:     os.Getenv("GITHUB_TOKEN"),
		Owner:     GetRepositoryOwner(),
		Committer: os.Getenv("GITHUB_NAME"),
		Email:     os.Getenv("GITHUB_EMAIL"),
		Repo:      GetRepositoryName(),
		Path:      "Boo Upload Test File.txt",
		Branch:    "main",
		Message:   "This is another test",
		Content:   "This is content",
	}
	err := UploadFileToGithub(params)
	assert.NoError(t, err)
}

func TestGetGithubRepositoryContent(t *testing.T) {
	params := GetParams{
		Token:  os.Getenv("GITHUB_TOKEN"),
		Owner:  GetRepositoryOwner(),
		Repo:   GetRepositoryName(),
		Path:   "code2323.py",
		Branch: "alsrl8",
	}
	content, err := GetGithubRepositoryContent(params)
	assert.NoError(t, err)
	log.Printf("content: %+v\n", content)
}

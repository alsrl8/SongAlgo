package github

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestUploadFileToGithub(t *testing.T) {
	params := UploadParams{
		Token:     os.Getenv("GITHUB_TOKEN"),
		Owner:     "alsrl8",
		Committer: os.Getenv("GITHUB_NAME"),
		Email:     os.Getenv("GITHUB_EMAIL"),
		Repo:      "SongAlgo",
		Path:      "Boo Upload Test File.txt",
		Branch:    "main",
		Message:   "This is another test",
		Content:   "This is content",
	}
	err := UploadFileToGithub(params)
	assert.NoError(t, err)
}

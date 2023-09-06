package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type UploadParams struct {
	Token     string
	Owner     string
	Committer string
	Email     string
	Repo      string
	Path      string
	Branch    string
	Message   string
	Content   string
}

type FetchParams struct {
	Owner  string
	Repo   string
	Branch string
	Path   string
}

func UploadFileToGithub(params UploadParams) error {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", params.Owner, params.Repo, params.Path)
	data := map[string]interface{}{
		"message":   params.Message,
		"content":   base64.StdEncoding.EncodeToString([]byte(params.Content)),
		"branch":    params.Branch,
		"committer": map[string]string{"name": params.Committer, "email": params.Email},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+params.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
	}(resp.Body)

	switch resp.StatusCode {
	case 404:
		return errors.New("UploadFileToGithub error: resource not found")
	case 409:
		return errors.New("UploadFileToGithub: conflict")
	case 422:
		return errors.New("UploadFileToGithub: Validation failed, or the endpoint has been spammed")
	}

	return nil
}

func FetchFromGithub(params FetchParams) ([]byte, error) {
	url := fmt.Sprintf(
		"https://raw.githubusercontent.com/%s/%s/%s/%s",
		params.Owner,
		params.Repo,
		params.Branch,
		params.Path,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

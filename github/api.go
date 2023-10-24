package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
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
	Sha       string
}

type GetParams struct {
	Token  string
	Owner  string
	Repo   string
	Branch string
	Path   string
}

type FetchParams struct {
	Owner  string
	Repo   string
	Branch string
	Path   string
}

type FileResponse struct {
	File       `json:"file"`
	StatusCode string `json:"statusCode"`
}

type File struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HtmlURL     string `json:"html_url"`
	GitURL      string `json:"git_url"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Encoding    string `json:"encoding"`
	Links       Links  `json:"_links"`
}

type Links struct {
	Self string `json:"self"`
	Git  string `json:"git"`
	HTML string `json:"html"`
}

func GetGithubRepositoryContent(params GetParams) (File, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s?ref=%s", params.Owner, params.Repo, params.Path, params.Branch)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return File{}, errors.Wrap(err, "failed to create new http get request")
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+params.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return File{}, errors.Wrap(err, "failed to execute http get request")
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			wrappedErr := errors.Wrap(closeErr, "failed to close http response body")
			log.Printf("%+v", wrappedErr)
		}
	}()

	switch resp.StatusCode {
	case 403:
		return File{}, ErrForbidden
	case 404:
		return File{}, ErrResourceNotFound
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error during reading readcloser body: %+v", err)
		return File{}, err
	}

	var file File
	if err := json.Unmarshal(body, &file); err != nil {
		return File{}, err
	}

	return file, nil
}

func GetGithubRepositorySource(branchName string, path string) (File, error) {
	params := GetParams{
		Token:  GetRepositoryToken(),
		Owner:  GetRepositoryOwner(),
		Repo:   GetRepositoryName(),
		Branch: branchName,
		Path:   path,
	}
	return GetGithubRepositoryContent(params)
}

func ConvertGithubRepositoryFileToFileResponse(file File, err error) FileResponse {
	if errors.Is(err, ErrResourceNotFound) {
		return FileResponse{File: file, StatusCode: "404"}
	}
	return FileResponse{File: file, StatusCode: "302"}
}

func UploadFileToGithub(params UploadParams) error {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", params.Owner, params.Repo, params.Path)
	data := map[string]interface{}{
		"message": params.Message,
		"content": base64.StdEncoding.EncodeToString([]byte(params.Content)),
		"branch":  params.Branch,
		"sha":     params.Sha,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "failed to convert byte array to json data")
	}

	req, err := http.NewRequest("PUT", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return errors.Wrap(err, "failed to create put request to github")
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+params.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to execute put request")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			err = errors.Wrap(err, "failed to close response body")
			log.Printf("%+v", err)
			return
		}
	}(resp.Body)

	log.Printf("got response from github api. Status Code:%+v, Body:%+v", resp.StatusCode, resp.Body)

	switch resp.StatusCode {
	case 401: // Unauthorized
		return errors.New("UploadFileToGithub: Unauthorized - Check your access credentials or token")
	case 404: // Not Found
		return errors.New("UploadFileToGithub: Resource not found - The endpoint or file may not exist")
	case 409: // Conflict
		return errors.New("UploadFileToGithub: Conflict - The file may already exist or there's a conflict with the request")
	case 422: // Unprocessable Entity
		return errors.New("UploadFileToGithub: Validation failed, or the endpoint has been spammed - Ensure request parameters are correct")
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error during closing response body: %+v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

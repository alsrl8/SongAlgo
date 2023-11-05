package selenium

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type APIResponse struct {
	Timestamp string             `json:"timestamp"`
	Channels  map[string]Channel `json:"channels"`
}

type Channel struct {
	Channel   string                   `json:"channel"`
	Version   string                   `json:"version"`
	Revision  string                   `json:"revision"`
	Downloads map[string][]DownloadUrl `json:"downloads"`
}

type DownloadUrl struct {
	Platform string `json:"platform"`
	Url      string `json:"url"`
}

func GetAPIResponse() *APIResponse {
	url := "https://googlechromelabs.github.io/chrome-for-testing/last-known-good-versions-with-downloads.json"
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to make GET request: %+v", err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Failed to close response: %+v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %+v", err)
		return nil
	}

	var apiResponse APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Printf("Failed to unmarshal json: %+v", err)
		return nil
	}

	return &apiResponse
}

func GetLatestStableDriverDownloadUrl(response *APIResponse, targetPlatform string) (string, error) {
	channelName := "Stable"
	if _, has := response.Channels[channelName]; !has {
		errMsg := fmt.Sprintf("Expected api response to have `%s` channel, but it is missing in the reposne", channelName)
		log.Printf(errMsg)
		return "", errors.New(errMsg)
	}

	stableChannel := response.Channels[channelName]
	chromeDriverLabel := "chromedriver"
	if _, has := stableChannel.Downloads[chromeDriverLabel]; !has {
		errMsg := fmt.Sprintf("Expected api response to have `%s` url, but it is missing in the reposne", chromeDriverLabel)
		log.Printf(errMsg)
		return "", errors.New(errMsg)
	}

	downloadUrls := stableChannel.Downloads[chromeDriverLabel]
	for _, downloadUrl := range downloadUrls {
		if downloadUrl.Platform != targetPlatform {
			continue
		}
		return downloadUrl.Url, nil
	}

	errMsg := fmt.Sprintf("No download url found for target platform(%s)", targetPlatform)
	log.Printf(errMsg)
	return "", errors.New(errMsg)
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Failed to close response: %+v", err)
		}
	}(resp.Body)

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Printf("Failed to close file: %+v", err)
		}
	}(out)

	_, err = io.Copy(out, resp.Body)
	return err
}

func Unzip(src string, dest string) ([]string, error) {
	var filenames []string

	reader, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer func(reader *zip.ReadCloser) {
		err := reader.Close()
		if err != nil {
			log.Printf("Failed to close zip reader: %+v", err)
		}
	}(reader)

	for _, f := range reader.File {
		fPath := filepath.Join(dest, f.Name)
		if !strings.HasPrefix(fPath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fPath)
		}

		filenames = append(filenames, fPath)

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(fPath, os.ModePerm)
			if err != nil {
				log.Printf("Failed to make directory of file(%s)", fPath)
				return filenames, err
			}
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fPath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return filenames, err
		}

		err = outFile.Close()
		if err != nil {
			fmt.Printf("Failed to close outfile: %+v", err)
			return nil, err
		}
		err = rc.Close()
		if err != nil {
			fmt.Printf("Failed to close rc: %+v", err)
			return nil, err
		}
	}
	return filenames, nil
}

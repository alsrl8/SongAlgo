package selenium

import (
	"testing"
)

func TestGetAPIResponse(t *testing.T) {
	apiResponse := GetAPIResponse()

	if apiResponse == nil {
		t.Fatalf("Expected non-nil response, got nil")
	}

	expectedChannels := []string{"Stable", "Beta", "Dev", "Canary"}
	for _, channelName := range expectedChannels {
		if _, has := apiResponse.Channels[channelName]; !has {
			t.Errorf("Expected response to have `%s` channel, but it is missing in the response", channelName)
		}
	}
}

func TestGetLatestStableDriverVersion(t *testing.T) {
	version, err := GetLatestStableDriverVersion()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("Latest stable driver version: %s", version)
}

func TestGetLatestStableDriverDownloadUrl(t *testing.T) {
	apiResponse := GetAPIResponse()
	targetPlatform := "win64"
	url, err := GetLatestStableDriverDownloadUrl(apiResponse, targetPlatform)
	if err != nil {
		t.Fatalf("Expected to get stable driver download url, but failed")
	}
	t.Logf("Latest stable driver download url: %s", url)
}

func TestDownloadFile(t *testing.T) {
	apiResponse := GetAPIResponse()
	targetPlatform := "win64"
	filepath := "./driver/chromedriver.zip"

	downloadUrl, err := GetLatestStableDriverDownloadUrl(apiResponse, targetPlatform)
	if err != nil {
		t.Fatalf("Expected to get stable driver download url, but failed")
	}

	err = DownloadFile(filepath, downloadUrl)
	if err != nil {
		t.Fatalf("Failed to download driver to filepath(%s) from given url(%s)", filepath, downloadUrl)
	}
}

func TestUnzip(t *testing.T) {
	srcPath := "./driver/chromedriver.zip"
	destPath := "./driver/"
	filenames, err := Unzip(srcPath, destPath)
	if err != nil {
		t.Fatalf("Failed to unzip zip file in path(%s)", srcPath)
	}
	t.Logf("Unzipped file names: %+v", filenames)
}

func TestMoveChromeDriver(t *testing.T) {
	driverPath := "./driver/chromedriver-win64/chromedriver.exe"
	destPath := "./driver/chromedriver.exe"
	MoveChromeDriver(driverPath, destPath)
}

func TestGetLocalDriverVersion(t *testing.T) {
	driverPath := "./driver/chromedriver.exe"
	version, err := GetLocalDriverVersion(driverPath)
	if err != nil {
		t.Fatalf("Failed to get local driver version: %+v", err)
	}
	t.Logf("Local driver version: %s", version)
}

func TestIsNeedUpdate(t *testing.T) {
	driverPath := "./driver/chromedriver.exe"
	ret, err := IsNeedUpdate(driverPath)
	if err != nil {
		t.Fatalf("Failed to compare chrome drivers' version: %+v", err)
	}

	if ret == true {
		t.Logf("Driver needs to be updated")
	} else {
		t.Logf("Driver doesn't need to be updated")
	}
}

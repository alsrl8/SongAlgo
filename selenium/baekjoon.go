package selenium

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

func getBjCookieDataPath() string {
	return getCookieDataPath()
}

func isFilePathValid(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

func readBjLoginCookiesJson() []selenium.Cookie {
	cookieDataPath := getBjCookieDataPath()
	if !isFilePathValid(cookieDataPath) {
		log.Println("Invalid cookie data path.")
		return []selenium.Cookie{}
	}

	file, err := os.Open(cookieDataPath)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return []selenium.Cookie{}
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}()

	// Read file contents
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return []selenium.Cookie{}
	}

	var jsonCookies []selenium.Cookie
	// Unmarshal JSON to our struct
	if err := json.Unmarshal(bytes, &jsonCookies); err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return []selenium.Cookie{}
	}

	return jsonCookies
}

func ReadCookieForBj() []selenium.Cookie {
	return readBjLoginCookiesJson()
}

func waitUntilUserCloseBrowser(wd *selenium.WebDriver) {
	c := make(chan bool)
	go monitorBrowserClose(wd, c)
	<-c
}

func monitorBrowserClose(wd *selenium.WebDriver, c chan bool) {
	for {
		time.Sleep(1 * time.Second) // Poll every second
		_, err := (*wd).Title()
		if err != nil {
			c <- true
			break
		}
	}
	log.Printf("Browser was closed by the user.")
}

func OpenBjWithCookie(url string) {
	bjLoginCookies := ReadCookieForBj()
	if len(bjLoginCookies) == 0 {
		return
	}

	rm, err := newResourceManager()
	if err != nil {
		log.Printf("Failed to create new resource manager: %v", err)
		return
	}
	defer cleanupResourceManager(rm)

	err = navigateToPage(rm.wd, url)
	if err != nil {
		log.Printf("Failed to access to url(%s): %v", url, err)
		return
	}

	for _, c := range bjLoginCookies {
		err := (*rm.wd).AddCookie(&c)
		if err != nil {
			log.Printf("Failed to add cookie for bj: %v", c)
			continue
		}
	}

	err = (*rm.wd).Refresh()
	if err != nil {
		log.Printf("Failed to refreshing page")
	}

	waitUntilUserCloseBrowser(rm.wd)
}

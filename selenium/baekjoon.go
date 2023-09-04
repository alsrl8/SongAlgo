package selenium

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

func getBjCookieDataPath() string {
	return "./selenium/cookie/cookies.json"
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

// TODO Manul Login 부분과 Json 저장 부분을 분리할 것
func performManualLogin(wd *selenium.WebDriver) error {
	err := navigateToLoginPage(wd)
	if err != nil {
		return err
	}

	monitorLoginStatus(wd)

	err = saveCurrentCookiesAsJson(wd)
	if err != nil {
		return err
	}

	return nil
}

func navigateToLoginPage(wd *selenium.WebDriver) error {
	loginPageUrl := "https://www.acmicpc.net/login?next=%2F"
	err := navigateToPage(wd, loginPageUrl)
	if err != nil {
		return err
	}
	return nil
}

func monitorLoginStatus(wd *selenium.WebDriver) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			el, err := (*wd).FindElement(selenium.ByCSSSelector, "ul.loginbar.pull-right")
			if err == nil {
				text, err := el.Text()
				if err != nil {
					continue
				}
				if strings.HasSuffix(text, "로그아웃") {
					return
				}
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func getCookies(wd *selenium.WebDriver) ([]selenium.Cookie, error) {
	cookies, err := (*wd).GetCookies()
	if err != nil {
		return []selenium.Cookie{}, err
	}
	return cookies, nil
}

func saveCurrentCookiesAsJson(wd *selenium.WebDriver) error {
	cookies, err := getCookies(wd)
	if err != nil {
		return err
	}

	cookieData, err := json.MarshalIndent(cookies, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(getBjCookieDataPath(), cookieData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadCookieForBj() []selenium.Cookie {
	return readBjLoginCookiesJson()
}

func GetCookieForBj() []selenium.Cookie {
	bjLoginCookies := ReadCookieForBj()
	if len(bjLoginCookies) > 0 {
		return bjLoginCookies
	}

	rm, err := newResourceManager()
	if err != nil {
		log.Printf("Failed to create new resource manager: %v", err)
		return []selenium.Cookie{}
	}
	defer cleanupResourceManager(rm)

	err = performManualLogin(rm.wd)
	if err != nil {
		log.Printf("Manual login failed: %v", err)
		return []selenium.Cookie{}
	}

	// TODO performManualLogin 이후 Cookie를 바로 쓸 수 있도록 수정할 것
	bjLoginCookies = ReadCookieForBj()
	if len(bjLoginCookies) > 0 {
		return bjLoginCookies
	}
	return []selenium.Cookie{}
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

	// Wait until the browser is closed by the user
	c := make(chan bool)
	go monitorBrowserClose(rm.wd, c)
	<-c
}

package selenium

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

type Cookie struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Path   string `json:"path"`
	Domain string `json:"domain"`
	Secure bool   `json:"secure"`
	Expiry int64  `json:"expiry"`
}

func getCookieDataPath() string {
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

func readLoginCookieJson() (loginCookie Cookie, ok bool) {
	cookieDataPath := getCookieDataPath()
	if !isFilePathValid(cookieDataPath) {
		return
	}

	file, err := os.Open(cookieDataPath)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Read file contents
	bytes, err := io.ReadAll(file)
	if err != nil {
		return
	}

	var cookie []Cookie

	// Unmarshal JSON to our struct
	err = json.Unmarshal(bytes, &cookie)
	if err != nil {
		return
	}

	for _, c := range cookie {
		if c.Name != "bojautologin" {
			continue
		}
		expiryTime := time.Unix(c.Expiry, 0)
		currentTime := time.Now()
		durationUntilExpiry := expiryTime.Sub(currentTime)
		if durationUntilExpiry > 0 {
			loginCookie = c
			ok = true
			break
		}
	}

	return
}

func manualLogin(wd *selenium.WebDriver) error {
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

	err := (*wd).Get(loginPageUrl)
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

	err = os.WriteFile(getCookieDataPath(), cookieData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadCookieForBJ() Cookie {
	if cookie, result := readLoginCookieJson(); result {
		return cookie
	}
	return Cookie{}
}

func GetCookieForBJ() Cookie {
	if cookie, result := readLoginCookieJson(); result {
		return cookie
	}

	rm, err := newResourceManager()
	if err != nil {
		fmt.Println(err)
	}
	defer func(rm *resourceManager) {
		err := rm.Cleanup()
		if err != nil {
			log.Fatalf("Failed to clean up resources in selenium: %s", err)
		}
	}(rm)

	err = manualLogin(rm.wd)
	if err != nil {
		log.Fatalf("Failed to manual login process: %s", err)
	}

	if cookie, result := readLoginCookieJson(); result {
		return cookie
	}
	return Cookie{}
}

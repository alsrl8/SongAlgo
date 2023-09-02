package selenium

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

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

	err = os.WriteFile("./selenium/cookie/cookies.json", cookieData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func CrawlBJ() {
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

	err = navigateToLoginPage(rm.wd)
	if err != nil {
		log.Fatalf("Failed to navigate to login page: %s", err)
	}

	monitorLoginStatus(rm.wd)

	err = saveCurrentCookiesAsJson(rm.wd)
	if err != nil {
		log.Fatalf("Failed to save cookies: %s", err)
	}
}

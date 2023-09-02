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

	wd := *rm.wd

	// Navigate to the login page
	err = wd.Get("https://www.acmicpc.net/login?next=%2F")
	if err != nil {
		log.Fatalf("Failed to navigate: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			// Check login status
			el, err := wd.FindElement(selenium.ByCSSSelector, "ul.loginbar.pull-right")
			if err == nil {
				text, err := el.Text()
				if err != nil {
					fmt.Println("Error getting text:", err)
					continue
				}
				fmt.Println(text)
				if strings.HasSuffix(text, "로그아웃") {
					goto Done
				}
			} else {
				fmt.Println("Couldn't find ul...")
			}
			time.Sleep(1 * time.Second)
		}
	}
Done:
	// Get cookies
	cookies, err := wd.GetCookies()
	if err != nil {
		log.Fatalf("Failed to get cookies: %s", err)
	}

	// Save cookies to a JSON file
	cookieData, err := json.Marshal(cookies)
	if err != nil {
		log.Fatalf("Failed to marshal cookies: %s", err)
	}
	err = os.WriteFile("./selenium/cookie/cookies.json", cookieData, 0644)
	if err != nil {
		log.Fatalf("Failed to write to file: %s", err)
	}

	log.Println("Cookies saved successfully.")
}

package selenium

import (
	"encoding/json"
	"github.com/tebeka/selenium"
	"io"
	"log"
	"os"
)

func getPgCookieDataPath() string {
	return getCookieDataPath()
}

func readPgLoginCookiesJson() []selenium.Cookie {
	cookieDataPath := getPgCookieDataPath()
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

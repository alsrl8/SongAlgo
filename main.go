package main

import "fmt"

func main() {
	//chrome.GetCookiesAndSaveAsJson()
	err := run()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

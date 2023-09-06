package main

import "fmt"

func main() {
	err := run()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

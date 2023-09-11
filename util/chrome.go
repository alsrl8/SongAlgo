package util

import (
	"bytes"
	"log"
	"os/exec"
	"syscall"
)

func IsChromeRunning() bool {
	cmd := exec.Command("tasklist", "/fi", "imagename eq chrome.exe")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	err := cmd.Run()
	if err != nil {
		log.Println("Failed to execute command:", err)
		return false
	}

	if len(out.String()) > 0 && !bytes.Contains(out.Bytes(), []byte("INFO: No tasks are running which match the specified criteria.")) {
		return true
	}

	return false
}

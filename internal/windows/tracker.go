// windows/tracker.go
package windows

import (
	"bytes"
	"os/exec"
	"strings"
	"time"
)

type WindowInfo struct {
	Title string
	Time  time.Time
}

func GetActiveWindow() (*WindowInfo, error) {
	cmd := exec.Command("xdotool", "getactivewindow", "getwindowname")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return &WindowInfo{
		Title: strings.TrimSpace(out.String()),
		Time:  time.Now(),
	}, nil
}

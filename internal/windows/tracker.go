// windows/tracker.go
package windows

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type WindowInfo struct {
	Title string
	Time  time.Time
}

func LogActiveWindow(outputDir string) (err error) {

	var wi WindowInfo
	if wi, err = getActiveWindow(); err != nil {
		return err
	}

	// logging to file
	filename := filepath.Join(outputDir, "", "winlog.log")

	var fs *os.File
	if fs, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		log.Fatal(err)
	}
	defer fs.Close()

	b := fmt.Appendf([]byte{}, "%s | %s\n", wi.Time.Format("2006-01-02_15-04-05"), wi.Title)
	if _, err = fs.Write(b); err != nil {
		return err
	}

	return err
}

func getActiveWindow() (wi WindowInfo, err error) {
	cmd := exec.Command("xdotool", "getactivewindow", "getwindowname")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err = cmd.Run(); err != nil {
		return wi, err
	}

	wi = WindowInfo{
		Title: strings.TrimSpace(out.String()),
		Time:  time.Now(),
	}

	return wi, err
}

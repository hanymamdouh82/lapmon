package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/hanymamdouh82/lapmon/internal/screen"
	"github.com/hanymamdouh82/lapmon/internal/windows"
)

func main() {
	// cli flags
	outputDir := flag.String("o", "/var/log", "Logs output directory. Default: /var/log")
	ttSS := flag.Int("ss", 30, "Screenshot time interval in seconds")
	ttWI := flag.Int("wi", 30, "Window information time interval in seconds")
	flag.Parse()

	// Paths
	date := time.Now().Format("2006-01-02")
	screenshotPath := filepath.Join(*outputDir, "screenshots", date)
	winlogPath := filepath.Join(*outputDir, "winlog")
	os.MkdirAll(*outputDir, 0755)
	os.MkdirAll(screenshotPath, 0755)
	os.MkdirAll(winlogPath, 0755)

	// Tickers
	tickerScreenshot := time.NewTicker(time.Duration(*ttSS) * time.Second)
	tickerWindow := time.NewTicker(time.Duration(*ttWI) * time.Second)
	// tickerHistory := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-tickerScreenshot.C:
			err := screen.TakeScreenshot(screenshotPath)
			if err != nil {
				log.Println("Screenshot error: ", err)
			}

		case <-tickerWindow.C:
			if err := windows.LogActiveWindow(winlogPath); err != nil {
				log.Println("Window Information error: ", err)
			}
		}
	}
}

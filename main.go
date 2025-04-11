package main

import (
	"log"
	"os"
	"time"

	"github.com/hanymamdouh82/lapmon/internal/screen"
	"github.com/hanymamdouh82/lapmon/internal/windows"
)

func main() {
	outputDir := "/home/hany/monitor" // adjust
	os.MkdirAll(outputDir, 0755)

	tickerScreenshot := time.NewTicker(5 * time.Second)
	tickerWindow := time.NewTicker(2 * time.Second)
	// tickerHistory := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-tickerScreenshot.C:
			err := screen.TakeScreenshot(outputDir)
			if err != nil {
				log.Println("Screenshot error:", err)
			}

		case <-tickerWindow.C:
			win, err := windows.GetActiveWindow()
			_ = win
			if err != nil {
				log.Fatal(err)
			}
			// fmt.Printf("[WINDOW] %s @ %s\n", win.Title, win.Time.Format(time.RFC3339))

			// case <-tickerHistory.C:
			// 	// entries, err := history.ReadChromeHistory("/home/hany/.config/google-chrome/Default/History")
			// 	// entries, err := history.ReadChromeHistory("/home/hany/.mozilla/firefox/ogeekpml.default-release/places.sqlite")
			// 	entries, err := history.ReadChromeHistory()
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}
			// 	for _, entry := range entries {
			// 		_ = entry
			// 		// fmt.Printf("[HISTORY] %s | %s @ %s\n", entry.Title, entry.URL, entry.VisitAt)
			// 		fmt.Println(entry)
			// 	}
		}
	}
}

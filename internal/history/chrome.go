// history/chrome.go
package history

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type HistoryEntry struct {
	URL     string
	Title   string
	VisitAt time.Time
}

func ReadChromeHistory() (history []string, err error) {

	profileDir := filepath.Join(os.Getenv("HOME"), ".mozilla", "firefox")
	matches, err := filepath.Glob(filepath.Join(profileDir, "*.default-release"))
	if err != nil || len(matches) == 0 {
		return history, fmt.Errorf("could not locate Firefox profile directory")
	}

	dbPath := filepath.Join(matches[0], "places.sqlite")
	tmpPath := filepath.Join(os.TempDir(), fmt.Sprintf("places_backup_%d.sqlite", time.Now().UnixNano()))

	// Run SQLite CLI backup
	cmd := exec.Command("sqlite3", dbPath, fmt.Sprintf(".backup '%s'", tmpPath))
	if out, err := cmd.CombinedOutput(); err != nil {
		return history, fmt.Errorf("sqlite3 backup failed: %s - %w", out, err)
	}
	defer os.Remove(tmpPath)

	// Now read from the clean backup
	query := `
		SELECT url
		FROM moz_places
		WHERE last_visit_date IS NOT NULL
		ORDER BY last_visit_date DESC
		LIMIT 10;
	`

	cmd = exec.Command("sqlite3", tmpPath, query)
	output, err := cmd.Output()
	if err != nil {
		return history, fmt.Errorf("sqlite3 query failed: %w", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		url := scanner.Text()
		history = append(history, url)
	}

	return history, nil
}

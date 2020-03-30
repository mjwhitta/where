package where

import (
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/mjwhitta/pathname"
)

var cache = map[string]string{}

// Version is the package version.
const Version = "1.0.2"

// Is will return the full path to the specified cmd if it exists in
// the defined PATH env var.
func Is(cmd string) string {
	if len(cmd) == 0 {
		return ""
	}

	// Return cached value if it exists
	var hasKey bool
	if _, hasKey = cache[cmd]; hasKey {
		return cache[cmd]
	}

	var dirs []string
	var exts []string
	var fullpath string

	// Get all PATH directories
	dirs = strings.Split(
		os.Getenv("PATH"),
		string(os.PathListSeparator),
	)

	// Get all valid extensions
	exts = strings.Split(os.Getenv("PATHEXT"), ";")
	if len(exts) == 0 {
		exts = append(exts, "")
	}

	// Find, cache, and return full path
	for _, dir := range dirs {
		for _, ext := range exts {
			fullpath = filepath.Join(dir, cmd+ext)
			if pathname.DoesExist(fullpath) {
				cache[cmd] = pathname.ExpandPath(fullpath)
				return cache[cmd]
			}
		}
	}

	// Otherwise return empty string
	return ""
}

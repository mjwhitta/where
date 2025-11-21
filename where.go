package where

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/mjwhitta/pathname"
)

// Is will return the full path to the specified cmd if it exists in
// the defined PATH env var.
func Is(cmd string) string {
	var cached any
	var dirs []string
	var exts []string
	var fullpath string
	var hasKey bool

	if cmd == "" {
		return ""
	}

	// Return cached value if it exists
	if cached, hasKey = cache.Get(cmd); hasKey {
		//nolint:errcheck,forcetypeassert // Was cached as string
		return cached.(string)
	}

	// Get all PATH directories
	dirs = filepath.SplitList(os.Getenv("PATH"))
	if len(dirs) == 0 {
		if path, ok := defPaths[runtime.GOOS]; ok {
			dirs = path
		}
	}

	// Get all valid extensions
	exts = filepath.SplitList(os.Getenv("PATHEXT"))
	if len(exts) == 0 {
		exts = append(exts, "")
	}

	// Find, cache, and return full path
	for _, dir := range dirs {
		for _, ext := range exts {
			fullpath = filepath.Join(dir, cmd+ext)

			if ok, e := pathname.DoesExist(fullpath); e != nil {
				continue // Not accessible
			} else if ok {
				fullpath = pathname.ExpandPath(fullpath)
				cache.Put(cmd, fullpath)

				return fullpath
			}
		}
	}

	// Otherwise return empty string
	return ""
}

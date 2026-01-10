package engine

import (
	"os"
	"runtime"
	"strings"
)

type OSInfo struct {
	PackageManager string
}

func DetectOS() OSInfo {
	// macOS
	if runtime.GOOS == "darwin" {
		return OSInfo{PackageManager: "brew"}
	}

	// Windows
	if runtime.GOOS == "windows" {
		return OSInfo{PackageManager: "winget"}
	}

	// Linux
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return OSInfo{PackageManager: "unknown"}
	}

	text := string(data)

	switch {
	case strings.Contains(text, "ID=fedora"):
		return OSInfo{PackageManager: "dnf"}

	case strings.Contains(text, "ID=ubuntu") || strings.Contains(text, "ID_LIKE=debian"):
		return OSInfo{PackageManager: "apt"}

	case strings.Contains(text, "ID=arch"):
		return OSInfo{PackageManager: "pacman"}

	default:
		return OSInfo{PackageManager: "unknown"}
	}
}

package utils

import (
	"os/exec"
	"runtime"
	"strings"
)

func CheckJavaInstallation() bool {

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("java", "-version")
	case "linux", "darwin":
		cmd = exec.Command("sh", "-c", "command -v java")
	}

	output, err := cmd.Output()
	if err != nil {
		return false
	}

	// On Unix-like systems, if java is not installed, the output will be empty.
	if runtime.GOOS != "windows" && len(output) == 0 {
		return false
	}

	if runtime.GOOS == "windows" {
		javaPath := string(output)
		if strings.Contains(strings.ToLower(javaPath), "program files") ||
			strings.Contains(strings.ToLower(javaPath), "program files (x86)") {
			return true
		} else {
			return false
		}
	}

	return true
}

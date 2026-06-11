package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetAppVersion() string {
	return "v0.4.2"
}

func getAppDataDir() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	appDir := filepath.Join(dir, "ya/data")
	err = os.MkdirAll(appDir, 0755)
	return appDir, err
}

func IsInvalidString(s string) bool {
	return len(s) == 0
}

// ResolveTemplates prompts the user interactively to fill in {placeholder} tokens.
// Used by the CLI path (non-TUI). The TUI handles its own template resolution.
func ResolveTemplates(command string) (string, error) {
	re := regexp.MustCompile(`\{([^}]+)\}`)
	matches := re.FindAllStringSubmatch(command, -1)
	if len(matches) == 0 {
		return command, nil
	}

	seen := map[string]bool{}
	var placeholders []string
	for _, m := range matches {
		if !seen[m[1]] {
			seen[m[1]] = true
			placeholders = append(placeholders, m[1])
		}
	}

	cyan := CAccent.Render
	yellow := CAccent.Render
	green := CSuccess.Render

	fmt.Printf("%s %s\n", cyan("→"), yellow("This command has template values. Fill them in below:"))
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	values := map[string]string{}

	for i, placeholder := range placeholders {
		fmt.Printf("  %s %s: ", green(fmt.Sprintf("[%d/%d]", i+1, len(placeholders))), cyan(placeholder))
		if !scanner.Scan() {
			fmt.Println()
			fmt.Println("Input cancelled.")
			os.Exit(0)
		}
		values[placeholder] = scanner.Text()
	}

	fmt.Println()

	resolved := re.ReplaceAllStringFunc(command, func(match string) string {
		key := strings.TrimSuffix(strings.TrimPrefix(match, "{"), "}")
		return values[key]
	})

	return resolved, nil
}


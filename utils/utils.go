package utils

import (
	"fmt"
	"os"
	"strings"
)

func formatColumn(text string, width int) string {
	runes := []rune(text)
	if len(runes) > width {
		return string(runes[:width-3]) + "..."
	}
	return fmt.Sprintf(fmt.Sprintf("%%-%ds", width), text)
}

func (s Struct_Server) Col(text string, width int) string {
	return formatColumn(text, width)
}

func ValidateStrictFlags() {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") && len(arg) > 2 {
			fmt.Printf("Error: Use the prefix '--' for flag. '%s' (ex: --%s)\n", arg, strings.TrimPrefix(arg, "-"))
			os.Exit(1)
		}
	}
}

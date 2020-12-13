package color

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
	"github.com/mgutz/ansi"
)

// ref : https://github.com/cli/cli/blob/959b1aae67a0b2cc050ecbd6614745c4f6f18a22/pkg/iostreams/color.go
var (
	magenta = ansi.ColorFunc("magenta")
	cyan    = ansi.ColorFunc("cyan")
	red     = ansi.ColorFunc("red")
	yellow  = ansi.ColorFunc("yellow")
	blue    = ansi.ColorFunc("blue")
	green   = ansi.ColorFunc("green")
	gray    = ansi.ColorFunc("black+h")
	bold    = ansi.ColorFunc("default+b")

	gray256 = func(t string) string {
		return fmt.Sprintf("\x1b[%d;5;%dm%s\x1b[m", 38, 242, t)
	}

	enabled      = true
	is256enabled = true
)

func init() {
	enabled = EnvColorForced() || (!EnvColorDisabled() && isTerminal(os.Stdout))
	is256enabled = Is256ColorSupported()
}

// EnvColorDisabled returns color supported boolean
func EnvColorDisabled() bool {
	return os.Getenv("NO_COLOR") != "" || os.Getenv("CLICOLOR") == "0"
}

// EnvColorForced returns color supported boolean
func EnvColorForced() bool {
	return os.Getenv("CLICOLOR_FORCE") != "" && os.Getenv("CLICOLOR_FORCE") != "0"
}

// Is256ColorSupported returns color supported boolean
func Is256ColorSupported() bool {
	term := os.Getenv("TERM")
	colorterm := os.Getenv("COLORTERM")

	return strings.Contains(term, "256") ||
		strings.Contains(term, "24bit") ||
		strings.Contains(term, "truecolor") ||
		strings.Contains(colorterm, "256") ||
		strings.Contains(colorterm, "24bit") ||
		strings.Contains(colorterm, "truecolor")
}

// Bold returns text with bold text when enabled is true
func Bold(t string) string {
	if !enabled {
		return t
	}
	return bold(t)
}

// Red returns text with red color when enabled is true
func Red(t string) string {
	if !enabled {
		return t
	}
	return red(t)
}

// Yellow returns text with yellow color when enabled is true
func Yellow(t string) string {
	if !enabled {
		return t
	}
	return yellow(t)
}

// Green returns text with green color when enabled is true
func Green(t string) string {
	if !enabled {
		return t
	}
	return green(t)
}

// Gray returns text with gray color when enabled is true
func Gray(t string) string {
	if !enabled {
		return t
	}
	if is256enabled {
		return gray256(t)
	}
	return gray(t)
}

// Magenta returns text with magenta color when enabled is true
func Magenta(t string) string {
	if !enabled {
		return t
	}
	return magenta(t)
}

// Cyan returns text with cyan color when enabled is true
func Cyan(t string) string {
	if !enabled {
		return t
	}
	return cyan(t)
}

// Blue returns text with blue color when enabled is true
func Blue(t string) string {
	if !enabled {
		return t
	}
	return blue(t)
}

// SuccessIcon returns a special character with green color
func SuccessIcon() string {
	return Green("✓")
}

// WarningIcon returns a special character with yello color
func WarningIcon() string {
	return Yellow("!")
}

// FailureIcon returns a special character with red color
func FailureIcon() string {
	return Red("X")
}

func isTerminal(f *os.File) bool {
	return isatty.IsTerminal(f.Fd()) || isatty.IsCygwinTerminal(f.Fd())
}

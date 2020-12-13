package logger

import (
	"fmt"
	"os"

	"github.com/cjaewon/deploysarus/utils/color"
)

// Print is fmt Print wrapper
func Print(v ...interface{}) {
	fmt.Print(v...)
}

// Printf is fmt Printf wrapper
func Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

// Println is fmt Println wrapper
func Println(v ...interface{}) {
	fmt.Println(v...)
}

// Printlnf is Println and Sprintf wrapper
func Printlnf(format string, v ...interface{}) {
	Println(fmt.Sprintf(format, v...))
}

// Warn is fmt Print wrapper with prefix
func Warn(v ...interface{}) {
	fmt.Print(color.Yellow("Warn") + ": ")
	fmt.Print(v...)
}

// Warnf is fmt Printf wrapper with prefix
func Warnf(format string, v ...interface{}) {
	fmt.Printf(color.Yellow("Warn") + ": ")
	fmt.Printf(format, v...)
}

// Warnln is Error wrapper with new line
func Warnln(v ...interface{}) {
	Warn(v...)
	fmt.Println()
}

// Error is fmt Print wrapper with prefix
func Error(v ...interface{}) {
	fmt.Print(color.Red("Error") + ": ")
	fmt.Print(v...)
}

// Errorf is fmt Printf wrapper with prefix
func Errorf(format string, v ...interface{}) {
	fmt.Printf(color.Red("Error") + ": ")
	fmt.Printf(format, v...)
}

// Errorln is Error wrapper with new line
func Errorln(v ...interface{}) {
	Error(v...)
	fmt.Println()
}

// Errorlnf is Errorf wrapper with new line
func Errorlnf(format string, v ...interface{}) {
	Errorf(format, v...)
	fmt.Println()
}

// ErrorFatal is Error wrapper with os exit
func ErrorFatal(v ...interface{}) {
	Error(v...)
	os.Exit(1)
}

// ErrorfFatal is Errorf wrapper with os exit
func ErrorfFatal(format string, v ...interface{}) {
	Errorf(format, v...)
	os.Exit(1)
}

// ErrorlnFatal is Errorln wrapper with os exit
func ErrorlnFatal(v ...interface{}) {
	Errorln(v...)
	os.Exit(1)
}

// ErrorlnfFatal is Errorlnf wrapper with os exit
func ErrorlnfFatal(format string, v ...interface{}) {
	Errorlnf(format, v...)
	os.Exit(1)
}

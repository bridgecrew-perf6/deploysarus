package commandline

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// ParseCommandline parses name, args from line
// ref https://github.com/laurent22/massren/blob/ae4c57da1e09a95d9383f7eb645a9f69790dec6c/main.go#L172
func ParseCommandline(line string) (string, []string, error) {
	var args []string
	state := "start"
	current := ""
	quote := "\""

	for i := 0; i < len(line); i++ {
		c := line[i]

		if state == "quotes" {
			if string(c) != quote {
				current += string(c)
			} else {
				args = append(args, current)
				current = ""
				state = "start"
			}
			continue
		}

		if c == '"' || c == '\'' {
			state = "quotes"
			quote = string(c)
			continue
		}

		if state == "arg" {
			if c == ' ' || c == '\t' {
				args = append(args, current)
				current = ""
				state = "start"
			} else {
				current += string(c)
			}
			continue
		}

		if c != ' ' && c != '\t' {
			state = "arg"
			current += string(c)
		}
	}

	if state == "quotes" {
		return "", []string{}, fmt.Errorf("Unclosed quote in command line: %s", line)
	}

	if current != "" {
		args = append(args, current)
	}

	if len(args) <= 0 {
		return "", []string{}, errors.New("Empty command line")
	}

	if len(args) == 1 {
		return args[0], []string{}, nil
	}

	return args[0], args[1:], nil
}

// Execute executes commandline
func Execute(name string, args ...string) error {
	cmd := exec.Command(name, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

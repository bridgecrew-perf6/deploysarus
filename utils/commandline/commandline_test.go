package commandline

import (
	"reflect"
	"testing"
)

func TestParseCommandline(t *testing.T) {
	cases := map[string][]string{
		"echo Hello World":          {"echo", "Hello", "World"},
		"echo \"Hello World\"":      {"echo", "Hello World"},
		"ls":                        {"ls"},
		"ls -l":                     {"ls", "-l"},
		"ls -l -f":                  {"ls", "-l", "-f"},
		"ls *.go":                   {"ls", "*.go"},
		"ls > test.go":              {"ls", ">", "test.go"},
		"git commit -m \"Testing\"": {"git", "commit", "-m", "Testing"},
		"git add *":                 {"git", "add", "*"},
	}

	for line, expected := range cases {
		name, args, err := ParseCommandline(line)
		if err != nil {
			t.Errorf("Failed to parse command line: %s", err)
		}

		if name != expected[0] || !reflect.DeepEqual(args, expected[1:]) {
			t.Errorf("Unexpected parsed command: expected-> %v, real-> %v %v", expected, name, args)
		}
	}
}

// func TestExecute(t *testing.T) {
// 	cases := map[string][]string{
// 		"echo Hello World":     {"Hello World"},
// 		"echo \"Hello World\"": {"Hello World"},
// 	}

// 	for line, expected := range cases {
// 		name, args, err := ParseCommandline(line)
// 		if err != nil {
// 			t.Errorf("Failed to parse command line: %s", err)
// 		}

// 		Execute(name, args...)
// 	}
// }

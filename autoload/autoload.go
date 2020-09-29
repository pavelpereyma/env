package autoload

import (
	"bufio"
	"io"
	"os"
	"strings"
)

const envFile = ".env"

func init() {
	if err := autoload(); err != nil {
		panic(err)
	}
}

// autoload parse .env file and set environment variables
func autoload() error {
	envMap, err := readFile(envFile)
	if err != nil {
		return err
	}

	currentEnv := map[string]bool{}
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentEnv[key] = true
	}

	for key, value := range envMap {
		err := os.Setenv(key, value)
		return err
	}
	return nil
}

// readFile opened file and return map of env key/value
func readFile(filename string) (envMap map[string]string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	return parse(file)
}

// parse reads an env file from io.Reader, returning a map of keys and values.
func parse(r io.Reader) (envMap map[string]string, err error) {
	envMap = make(map[string]string)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		if validLine(line) {
			setLineKeyValue(line, envMap)
		}
	}

	if err = scanner.Err(); err != nil {
		return
	}
	return envMap, nil
}

// validLine checks if line valid for env parse
func validLine(line string) bool {
	if len(line) == 0 {
		return false
	}
	if strings.HasPrefix(line, "#") {
		return false
	}
	if !strings.Contains(line, "=") {
		return false
	}
	return true
}

// setLineKeyValue split line and set env key\value to map
func setLineKeyValue(line string, envMap map[string]string) {
	lineArgs := strings.Split(line, "=")
	if len(lineArgs) != 2 {
		return
	}
	key := lineArgs[0]
	if strings.Contains(key, "export") {
		key = strings.TrimPrefix(key, "export")
	}
	key = strings.TrimSpace(key)
	value := strings.TrimSpace(lineArgs[1])

	envMap[key] = value
}

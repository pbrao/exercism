package grep

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Search acts like grep
func Search(pattern string, flags []string, files []string) []string {

	var result []string
	var lineNumber, insensitive, fileNames, invert, wholeLine bool

	for _, flag := range flags {
		switch flag {
		case "-n": //  Print the line numbers of each matching line.
			lineNumber = true
		case "-i": //  Match line using a case-insensitive comparison.
			insensitive = true
		case "-l": //  Print only the names of files that contain at least one matching line.
			fileNames = true
		case "-v": //  Invert the program -- collect all lines that fail to match the pattern.
			invert = true
		case "-x": //  Only match entire lines, instead of lines that contain a match.
			wholeLine = true
		}
	}

	for _, file := range files {
		fmt.Println(file)
		f, err := os.Open(file)
		if err != nil {
			fmt.Println("Error reading the file")
		}
		defer f.Close()

		lineCount := 1
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if (insensitive && strings.Contains(strings.ToLower(line), strings.ToLower(pattern))) || strings.Contains(line, pattern) {
				fmt.Println(line)
				if fileNames {
					result = append(result, file)
				} else if lineNumber {
					result = append(result, strconv.Itoa(lineCount)+":"+line)
				} else {
					result = append(result, line)
				}
			}
			lineCount++
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("error ")
		}
	}
	return result
}

func insensitiveContains(line, pattern string) bool {
	return strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
}

func sensitiveContains(line, pattern string) bool {
	return strings.Contains(line, pattern)
}
func invertContains(line, pattern string) bool {
	return true
}

func wholeContains(line, pattern string) bool {
	return true
}

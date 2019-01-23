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
	var lineNumber, insensitive bool

	for _, flag := range flags {
		switch flag {
		case "-n":
			lineNumber = true
		case "-i":
			insensitve = true
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
			if strings.Contains(line, pattern) {
				fmt.Println(line)
				if lineNumber {
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

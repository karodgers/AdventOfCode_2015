package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isNice(str string) bool {
	vowels := "aeiou"
	disallowed := []string{"ab", "cd", "pq", "xy"}
	vowelCount := 0
	hasDouble := false

	// Rule 1: Count vowels and check double letters
	for i := 0; i < len(str); i++ {
		// Check for vowels
		if strings.ContainsRune(vowels, rune(str[i])) {
			vowelCount++
		}
		// Check for double letters
		if i > 0 && str[i] == str[i-1] {
			hasDouble = true
		}
	}

	// Rule 2: Check disallowed substrings
	for _, substr := range disallowed {
		if strings.Contains(str, substr) {
			return false
		}
	}

	return vowelCount >= 3 && hasDouble
}

func OpenFile() {
	file, err := os.Open("nice_day5.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	niceCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if isNice(line) {
			niceCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Number of nice strings:", niceCount)
}

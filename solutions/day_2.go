package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func WrappingPaper(l, w, h int) int {
	values := []int{l, w, h}
	surfaceArea := 2*l*w + 2*w*h + 2*h*l

	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	newValues := values[:len(values)-1]

	var result int
	slack := 1
	for _, num := range newValues {
		slack = slack * num
	}
	result = slack + surfaceArea

	return result
}

func OpenFileDay2() {
	file, err := os.Open("wrapper_day2.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	totalWrappingPaper := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		dimensions := strings.Split(line, "x")
		if len(dimensions) != 3 {
			fmt.Println("Invalid dimensions:", line)
			continue
		}

		l, err1 := strconv.Atoi(dimensions[0])
		w, err2 := strconv.Atoi(dimensions[1])
		h, err3 := strconv.Atoi(dimensions[2])

		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("Error parsing dimensions:", line)
			continue
		}

		totalWrappingPaper += WrappingPaper(l, w, h)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}

	fmt.Println(totalWrappingPaper)
}

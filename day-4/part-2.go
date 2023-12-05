package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getScratchcards() []int {
	points := make([]int, 208)

	for i := range points {
		points[i] = 1
	}

	return points
}

func Scratchcards2() {
	f, _ := os.Open("./day-4/cards.txt")
	scanner := bufio.NewScanner(f)
	scratchcards := getScratchcards()
	lineIndex := 0
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Split(line, ":")[1]
		points := 0

		winningMap := make(map[int]struct{})
		winningNumbers := strings.Fields(strings.Split(line, "|")[0])

		for _, winNum := range winningNumbers {
			if value, err := strconv.Atoi(winNum); err == nil {
				winningMap[value] = struct{}{}
			}
		}

		currentNumbers := strings.Fields(strings.Split(line, "|")[1])

		for _, currentNum := range currentNumbers {
			if value, err := strconv.Atoi(currentNum); err == nil {
				if _, ok := winningMap[value]; ok {
					points++
				}
			}
		}

		for i := 1; i <= points; i++ {
			scratchcards[lineIndex+i] = scratchcards[lineIndex+i] + scratchcards[lineIndex]
		}

		lineIndex++
	}

	fmt.Println(scratchcards)
	for _, value := range scratchcards {
		total += value
	}
	fmt.Println(total)
}

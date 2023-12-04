package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Scratchcards() {
	f, _ := os.Open("./day-4/cards.txt")
	scanner := bufio.NewScanner(f)
	totalPoints := 0

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

		totalPoints += getPoints()[points]
	}

	fmt.Println(totalPoints)
}

func getPoints() []int {
	points := make([]int, 11)

	points[0] = 0
	points[1] = 1
	points[2] = 2
	points[3] = 4
	points[4] = 8
	points[5] = 16
	points[6] = 32
	points[7] = 64
	points[8] = 128
	points[9] = 256
	points[10] = 512

	return points
}

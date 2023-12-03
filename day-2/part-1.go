package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CubeCondrum() {
	f, _ := os.Open("./day-2/records-input.txt")

	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		sum += parseIt(scanner.Text())
	}

	fmt.Println(sum)
}

func getMaxCubeValues() map[string]int {
	return map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
}

func parseIt(game string) int {
	id := strings.Split(strings.Split(game, ":")[0], " ")[1]
	sets := strings.Split(strings.Split(game, ":")[1], ";")
	for _, set := range sets {
		fmt.Println(set)
		cubes := strings.Split(set, ",")

		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			color := strings.Split(cube, " ")[1]
			count, _ := strconv.Atoi(strings.Split(cube, " ")[0])

			fmt.Println(color)
			fmt.Println(count)
			if getMaxCubeValues()[color] < count {
				return 0
			}
		}
	}

	value, _ := strconv.Atoi(id)

	return value
}

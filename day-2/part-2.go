package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CubeCondrum2() {
	f, _ := os.Open("./day-2/records-input.txt")

	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		sum += parseIt2(scanner.Text())
	}

	fmt.Println(sum)
}

func parseIt2(game string) int {
	sets := strings.Split(strings.Split(game, ":")[1], ";")

	maxCubeValues := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, set := range sets {
		fmt.Println(set)
		cubes := strings.Split(set, ",")

		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			color := strings.Split(cube, " ")[1]
			count, _ := strconv.Atoi(strings.Split(cube, " ")[0])

			fmt.Println(color)
			fmt.Println(count)

			if maxCubeValues[color] < count {
				maxCubeValues[color] = count
			}
		}
	}

	return maxCubeValues["red"] * maxCubeValues["green"] * maxCubeValues["blue"]
}

package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Trebuchet() {
	sum := 0

	f, _ := os.Open("./calibration-document.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		sum += timeToCalibrate(line)
	}

	fmt.Println(sum)
}

func timeToCalibrate(calibrationValue string) int {
	left, right := 0, len(calibrationValue)-1
	var leftNumber, rightNumber int

	for leftNumber == 0 || rightNumber == 0 {
		if v, err := strconv.Atoi(string(calibrationValue[left])); err == nil {
			leftNumber = v
		} else {
			left++
		}

		if v, err := strconv.Atoi(string(calibrationValue[right])); err == nil {
			rightNumber = v
		} else {
			right--
		}
	}

	return leftNumber*10 + rightNumber
}

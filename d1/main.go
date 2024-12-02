package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		return
	}

	var leftSlice, rightSlice []int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		currentLine := strings.Split(scanner.Text(), "   ")

		leftInt, err := strconv.Atoi(currentLine[0])
		if err != nil {
			return
		}
		rightInt, err := strconv.Atoi(currentLine[1])
		if err != nil {
			return
		}

		leftSlice = append(leftSlice, leftInt)
		rightSlice = append(rightSlice, rightInt)
	}
	slices.Sort(leftSlice)
	slices.Sort(rightSlice)

	var distance int
	var freqMap = make(map[int]int)
	for index := range len(leftSlice) {
		distance += int(math.Abs(float64(leftSlice[index] - rightSlice[index])))
		freqMap[leftSlice[index]] = 0
	}
	fmt.Println("Distance: ", distance)

	for _, value := range rightSlice {
		if _, ok := freqMap[value]; ok {
			freqMap[value] += 1
		}
	}
	var simScore int
	for key, value := range freqMap {
		simScore += key * value
	}
	fmt.Println("Similarity Score: ", simScore)
}

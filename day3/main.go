package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rectangle := [1000][1000]int{}
	input := []string{}
	for scanner.Scan() {
		a := scanner.Text()
		input = append(input, a)
		rectangle = process(a, rectangle)
	}
	count := 0
	for _, row := range rectangle {
		for _, v := range row {
			if v >= 2 {
				count++
			}
		}
	}
	fmt.Println(count)
	fmt.Println(find_no_overlap(rectangle, input))
}

func find_no_overlap(rectangle [1000][1000]int, input []string) string {
	for _, v := range input {
		value, ok := find_no_overlap_in_each_row(rectangle, v)
		if ok {
			return value
		}
	}
	return ""
}

func find_no_overlap_in_each_row(rectangle [1000][1000]int, input string) (string, bool) {
	matrix := strings.Split(input, "@ ")[1]
	temp := strings.Split(matrix, ": ")
	indexes := strings.Split(temp[0], ",")
	ranges := strings.Split(temp[1], "x")
	x, _ := strconv.Atoi(indexes[0])
	y, _ := strconv.Atoi(indexes[1])
	x_count, _ := strconv.Atoi(ranges[0])
	y_count, _ := strconv.Atoi(ranges[1])
	for i := 0; i < x_count; i++ {
		for j := 0; j < y_count; j++ {
			if rectangle[x+i][y+j] != 1 {
				return "", false
			}
		}
	}
	return strings.Split(input, "@ ")[0], true
}

func process(input string, rectangle [1000][1000]int) [1000][1000]int {
	matrix := strings.Split(input, "@ ")[1]
	temp := strings.Split(matrix, ": ")
	indexes := strings.Split(temp[0], ",")
	ranges := strings.Split(temp[1], "x")
	x, _ := strconv.Atoi(indexes[0])
	y, _ := strconv.Atoi(indexes[1])
	x_count, _ := strconv.Atoi(ranges[0])
	y_count, _ := strconv.Atoi(ranges[1])
	for i := 0; i < x_count; i++ {
		for j := 0; j < y_count; j++ {
			rectangle[x+i][y+j] += 1
		}
	}
	return rectangle
}

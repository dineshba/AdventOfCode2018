package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := []string{}
	for scanner.Scan() {
		a := scanner.Text()
		input = append(input, a)
	}
	fmt.Println(process(input))
}

func process(input []string) string {
	for index, each := range input {
		ok, value := findMatch(each, input[(index+1):])
		if ok {
			return value
		}
	}
	return "not found"
}
func findMatch(each string, input []string) (bool, string) {
	for _, value := range input {
		isClose, index := compare(each, value)
		if isClose {
			result := value[:index] + value[(index+1):]
			return true, result
		}
	}
	return false, ""
}

func compare(a string, b string) (bool, int) {
	diff := 0
	changingIndex := -1
	for index, character := range a {
		characterToByte := []byte(string(character))[0]
		if b[index] != characterToByte {
			diff = diff + 1
			changingIndex = index
			if diff > 1 {
				return false, -1
			}
		}
	}
	return true, changingIndex
}

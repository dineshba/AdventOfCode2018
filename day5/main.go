package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a := scanner.Text()
		process(a)
	}
}

func process(input string) {
	fmt.Println("Problem one")
	fmt.Println(len(computeAfterReaction(input)))
	fmt.Println("Problem two")
	fmt.Println(shortestPolymer(input))
}

func shortestPolymer(input string) int {
	charMap := make(map[rune]rune)
	for _, char := range input {
		if char > 96 {
			charMap[char] = char - 32
		}
	}
	currentMin := len(computeAfterReaction(input))
	for k, v := range charMap {
		unitsRemovedString := strings.Replace(input, string(k), "", -1)
		unitsRemovedString = strings.Replace(unitsRemovedString, string(v), "", -1)
		lenghtOfPolymer := len(computeAfterReaction(unitsRemovedString))
		if lenghtOfPolymer < currentMin {
			currentMin = lenghtOfPolymer
		}
	}
	return currentMin
}

func computeAfterReaction(input string) string {
	iterationString := input
	for {
		shouldContinue := false
		previousChar := iterationString[0]
		var newString string = ""
		for _, char := range iterationString {
			charByte := charToByte(char)
			diff := charByte - previousChar
			diff2 := previousChar - charByte
			if (diff == 32) || (diff2 == 32) {
				last := len(newString) - 1
				newString = newString[:last]
				shouldContinue = true
			} else {
				newString += string(char)
			}
			last := len(newString) - 1
			if last > 0 {
				previousChar = newString[last]
			}
		}
		if !shouldContinue || len(newString) == 0 {
			return newString
		}
		iterationString = newString
	}
}

func charToByte(char rune) byte {
	return ([]byte(string(char)))[0]
}

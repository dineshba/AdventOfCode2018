package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Input file not found")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inputs := [][]int{}
	inputsMap := make(map[string][]int)
	for scanner.Scan() {
		input := scanner.Text()
		inputArr := strings.Split(input, ", ")
		value1, _ := strconv.Atoi(inputArr[0])
		value2, _ := strconv.Atoi(inputArr[1])
		inputsMap[input] = []int{value1, value2}
		inputs = append(inputs, []int{value1, value2})
	}
	top, bottom, left, right := borders(inputs)
	grid := buildGrid(top, bottom, left, right, inputsMap)

	finiteKeys := findInfiniteBoxes(grid, left, right, top, bottom, inputsMap)
	finitekeysWithArea := map[string]int{}
	for k, v := range finiteKeys {
		if v {
			area := 0
			for _, row := range grid {
				for _, box := range row {
					if box.nearestCell == k {
						area++
					}
				}
			}
			finitekeysWithArea[k] = area
		}
	}

	max := 0
	for _, v := range finitekeysWithArea {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
	//	for _, box := range safeBoxes {
	//		fmt.Println(box.x, box.y)
	//	}
	count := 0
	for _, row := range grid {
		for _, box := range row {
			distance := sumOfManhattanDistance(inputsMap, box)
			if distance < 10000 {
				count++
			} else {
			}
		}
	}
	fmt.Println(count)

}

func sumOfManhattanDistance(inputsMap map[string][]int, box *Box) int {
	sum := 0
	for _, v := range inputsMap {
		sum += distanceBetween(v, []int{box.x, box.y})
	}
	return sum
}

func buildGrid(top, bottom, left, right int, inputsMap map[string][]int) [][]*Box {
	grid := [][]*Box{}
	for i := top; i <= bottom; i++ {
		row := []*Box{}
		for j := left; j <= right; j++ {
			key := fmt.Sprintf("%d, %d", i, j)
			_, ok := inputsMap[key]
			cellName := ""
			distance := math.MaxInt32
			if ok {
				cellName = key
				distance = 0
			}
			row = append(row, &Box{nearestCell: cellName, distance: distance, x: i, y: j})
		}
		grid = append(grid, row)
	}
	populateDistance(grid, inputsMap)
	return grid
}

func findInfiniteBoxes(grid [][]*Box, left, right, top, bottom int, inputsMap map[string][]int) map[string]bool {
	allKeys := map[string]bool{}
	for k, _ := range inputsMap {
		allKeys[k] = true
	}
	for _, row := range grid {
		for _, box := range row {
			if box.x == left || box.x == right || box.y == top || box.y == bottom {
				if box.nearestCell != "" {
					allKeys[box.nearestCell] = false
				}
			}
		}
	}
	return allKeys
}

func populateDistance(grid [][]*Box, inputsMap map[string][]int) {
	for k, v := range inputsMap {
		for _, row := range grid {
			for _, box := range row {
				distance := distanceBetween(v, []int{box.x, box.y})
				if distance < box.distance {
					box.distance = distance
					box.nearestCell = k
				}
				if distance == box.distance && box.nearestCell != k {
					box.nearestCell = ""
				}
			}
		}
	}
}

func distanceBetween(v1, v2 []int) int {
	diff1 := v1[0] - v2[0]
	diff2 := v1[1] - v2[1]
	if diff1 < 0 {
		diff1 = -diff1
	}
	if diff2 < 0 {
		diff2 = -diff2
	}
	return diff1 + diff2
}

type Box struct {
	nearestCell string
	distance    int
	x           int
	y           int
}

func (b Box) String() string {
	return fmt.Sprintf(b.nearestCell, b.distance, b.x, b.y)
}

func borders(inputs [][]int) (int, int, int, int) {
	top := inputs[0][1]
	bottom := inputs[0][1]
	right := inputs[0][0]
	left := inputs[0][0]
	for _, value := range inputs {
		x := value[0]
		y := value[1]
		if x < left {
			left = x
		}
		if x > right {
			right = x
		}
		if y < top {
			top = y
		}
		if y > bottom {
			bottom = y
		}
	}
	return top, bottom, left, right
}

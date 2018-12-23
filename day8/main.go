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
		fmt.Println("Input file not found")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputArr := strings.Split(scanner.Text(), " ")
		inputs := []int{}
		for _, i := range inputArr {
			value, _ := strconv.Atoi(i)
			inputs = append(inputs, value)
		}

		var root *Node = buildNodes(inputs)
		metadataEntries := getAllMetadataEntries(root)
		sum := 0
		for _, value := range metadataEntries {
			sum += value
		}
		fmt.Println(sum)
		fmt.Println("second problem")
		fmt.Println(findValue(root))
	}
}

type Node struct {
	Metadata         []int
	Children         []*Node
	NumberOfChildren int
	NumberOfMetadata int
}

func findValue(root *Node) int {
	if len(root.Children) == 0 {
		sumOfEntries := 0
		for _, v := range root.Metadata {
			sumOfEntries += v
		}
		return sumOfEntries
	}
	value := 0
	for _, metadata := range root.Metadata {
		if metadata <= len(root.Children) {
			value += findValue(root.Children[metadata-1])
		}
	}
	return value
}

func getAllMetadataEntries(root *Node) []int {
	return getMetaDataEntries(root)
}

func getMetaDataEntries(root *Node) []int {
	temp := []int{}
	for _, c := range root.Children {
		temp = append(temp, getMetaDataEntries(c)...)
	}
	return append(temp, root.Metadata...)
}

func buildNodes(inputs []int) *Node {
	if len(inputs) == 0 {
		return nil
	}
	numberOfChildren := inputs[0]
	numberOfMetadata := inputs[1]
	if numberOfChildren == 0 {
		return &Node{NumberOfChildren: 0, NumberOfMetadata: numberOfMetadata, Metadata: inputs[2:(2 + numberOfMetadata)]}
	}
	var children []*Node
	for i := 0; i < numberOfChildren; i++ {
		skippingCount := 2 + sizeOf(children)
		node := buildNodes(inputs[skippingCount:])
		children = append(children, node)
	}
	startingIndex := 2 + sizeOf(children)
	endingIndex := startingIndex + numberOfMetadata
	return &Node{NumberOfMetadata: numberOfMetadata, NumberOfChildren: numberOfChildren, Children: children, Metadata: inputs[startingIndex:endingIndex]}
}

func sizeOf(nodes []*Node) int {
	size := 0
	for _, node := range nodes {
		size += (2 + node.NumberOfMetadata + sizeOf(node.Children))
	}
	return size
}

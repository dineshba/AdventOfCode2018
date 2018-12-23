package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Node struct {
	Id           string
	Next         []*Node
	Previous     []*Node
	IsProcessing bool
	IsProcessed  bool
	Duration     int
}

func (n Node) IsReady() bool {
	return len(n.Previous) == 0
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nodeMap := make(map[string]*Node)
	for scanner.Scan() {
		a := scanner.Text()
		previousNodeId, nextNodeId := process(a)
		populateMap(nodeMap, previousNodeId, nextNodeId)
	}
	fmt.Println(getSteps(nodeMap, 1))
}

func populateMap(nodeMap map[string]*Node, previousNodeId, nextNodeId string) {
	previousNode, ok := nodeMap[previousNodeId]
	if !ok {
		duration := int(previousNodeId[0]) - 4
		previousNode = &Node{Id: previousNodeId, Next: []*Node{}, Previous: []*Node{}, IsProcessing: false, IsProcessed: false, Duration: duration}
		nodeMap[previousNodeId] = previousNode
	}

	nextNode, ok := nodeMap[nextNodeId]
	if !ok {
		var duration int = int(nextNodeId[0]) - 4
		nextNode = &Node{Id: nextNodeId, Next: []*Node{}, Previous: []*Node{}, IsProcessed: false, Duration: duration}
		nodeMap[nextNodeId] = nextNode
	}
	previousNode.Next = append(previousNode.Next, nextNode)
	nextNode.Previous = append(nextNode.Previous, previousNode)
}

func getSteps(nodeMap map[string]*Node, workersCount int) string {
	outputString := ""
	processingNodes := []*Node{}
	for {
		readyNodes := getReadyNodes(nodeMap)
		nodes := []*Node{}
		if len(readyNodes) != 0 {
			nodes = getNodesToProcess(readyNodes, workersCount - len(processingNodes))
		}
		processingNodes = append(processingNodes, nodes...)
		if len(processingNodes) == 0 {
			return outputString
		}
		for _, node := range processingNodes {
			fmt.Println(node)
			node.Duration = node.Duration - 1
			node.IsProcessing = true
			if node.Duration == 0 {
				for _, nextNode := range node.Next {
					nextNode.Previous = filter(nextNode.Previous, *node)
				}
				node.IsProcessed = true
				outputString += node.Id
			}
		}
		tempProcessingNodes := []*Node{}
		for _, n := range processingNodes {
			if !n.IsProcessed {
				tempProcessingNodes = append(tempProcessingNodes, n)
				//fmt.Println(n)
			} else {
				//fmt.Println("done with ", n)
			}
		}
		processingNodes = tempProcessingNodes
		//fmt.Println("iteration over")
	}
}

func getReadyNodes(nodeMap map[string]*Node) []*Node {
	var readyNodes []*Node
	for _, v := range nodeMap {
		if v.IsReady() && !v.IsProcessing && !v.IsProcessed {
			readyNodes = append(readyNodes, v)
		}
	}
	return readyNodes
}

func filter(nodes []*Node, node Node) []*Node {
	if len(nodes) == 0 {
		return []*Node{}
	}
	temp := nodes[:0]
	for _, n := range nodes {
		if n.Id != node.Id {
			temp = append(temp, n)
		}
	}
	return temp
}

func getNodesToProcess(nodes []*Node, count int) []*Node {
	if len(nodes) == 1 {
		return nodes
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Id < nodes[j].Id
	})
	//fmt.Println(len(nodes))
	//fmt.Println(count)
	if len(nodes) <= count {
		return nodes
	}
	return nodes[:count]
}

func process(input string) (string, string) {
	str1 := strings.Split(input, " ")
	return str1[1], str1[7]
}

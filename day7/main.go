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

func (n Node) String() string {
	return fmt.Sprintf("node:", n.Id, n.IsProcessing, n.Duration)
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
	fmt.Println(getSteps(nodeMap, 5))
}

func populateMap(nodeMap map[string]*Node, previousNodeId, nextNodeId string) {
	previousNode, ok := nodeMap[previousNodeId]
	offset := 4
	if !ok {
		duration := int(previousNodeId[0]) - offset
		previousNode = &Node{Id: previousNodeId, Next: []*Node{}, Previous: []*Node{}, IsProcessing: false, IsProcessed: false, Duration: duration}
		nodeMap[previousNodeId] = previousNode
	}

	nextNode, ok := nodeMap[nextNodeId]
	if !ok {
		var duration = int(nextNodeId[0]) - offset
		nextNode = &Node{Id: nextNodeId, Next: []*Node{}, Previous: []*Node{}, IsProcessed: false, Duration: duration}
		nodeMap[nextNodeId] = nextNode
	}
	previousNode.Next = append(previousNode.Next, nextNode)
	nextNode.Previous = append(nextNode.Previous, previousNode)
}

func getSteps(nodeMap map[string]*Node, workersCount int) string {
	outputString := ""
	counter := 0
	processingNodes := []*Node{}
	for {
		readyNodes := getReadyNodes(nodeMap)
		if len(readyNodes) != 0 {
			if len(processingNodes) < workersCount {
				nodesToBeProcessed := getNodesToProcess(readyNodes, workersCount-len(processingNodes))
				processingNodes = append(processingNodes, nodesToBeProcessed...)
			}
		}
		if len(processingNodes) == 0 {
			fmt.Println(counter)
			return outputString
		}
		counter++
		processingNodesForNextIteration := []*Node{}
		for _, node := range processingNodes {
			node.Duration = node.Duration - 1
			node.IsProcessing = true
			if node.Duration == 0 {
				for _, nextNode := range node.Next {
					nextNode.Previous = filter(nextNode.Previous, *node)
				}
				node.IsProcessed = true
				outputString += node.Id
			} else {
				processingNodesForNextIteration = append(processingNodesForNextIteration, node)
			}
		}
		processingNodes = processingNodesForNextIteration
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
	filteredNodes := nodes[:0]
	for _, n := range nodes {
		if n.Id != node.Id {
			filteredNodes = append(filteredNodes, n)
		}
	}
	return filteredNodes
}

func getNodesToProcess(nodes []*Node, count int) []*Node {
	if count == 0 {
		return []*Node{}
	}
	if len(nodes) == 1 {
		return nodes
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Id < nodes[j].Id
	})
	if len(nodes) <= count {
		return nodes
	}
	return nodes[:count]
}

func process(input string) (string, string) {
	str1 := strings.Split(input, " ")
	return str1[1], str1[7]
}

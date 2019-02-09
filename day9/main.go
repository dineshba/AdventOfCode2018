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
		playersCount := 0
		maxMarble := 0
		inputArr := strings.Split(scanner.Text(), " ")
		playersCount, _ = strconv.Atoi(inputArr[0])
		maxMarble, _ = strconv.Atoi(inputArr[6])
		players := createPlayers(playersCount)
		getPlayerWithHighScore(players, maxMarble)
	}
}

func createPlayers(count int) []*Player {
	players := []*Player{}
	for i := 1; i <= count; i++ {
		players = append(players, &Player{id: i, score: 0})
	}
	return players
}

func getPlayerWithHighScore(players []*Player, maxMarble int) int {
	marbleCount := 0
	var currentNode = &Node{id: marbleCount}
	currentNode.left = currentNode
	currentNode.right = currentNode
	currentPlayerPosition := 0
	marbleCount = 1
	for {
		if marbleCount%23 == 0 {
			currentPlayer := players[currentPlayerPosition]
			seventhNode := get7thNodeCounterClockwiseFrom(currentNode)
			leftNode := seventhNode.left
			rightNode := seventhNode.right
			rightNode.left = leftNode
			leftNode.right = rightNode

			currentPlayer.score += seventhNode.id
			currentPlayer.score += marbleCount

			currentNode = rightNode
		} else {
			newNode := &Node{id: marbleCount}
			firstNode := currentNode.right
			secondNode := firstNode.right
			firstNode.right = newNode
			newNode.right = secondNode
			secondNode.left = newNode
			newNode.left = firstNode
			currentNode = newNode
		}
		marbleCount++
		currentPlayerPosition++
		currentPlayerPosition %= len(players)
		if marbleCount > (maxMarble + 1) {
			maxScore := 0
			for _, player := range players {
				if player.score > maxScore {
					maxScore = player.score
				}
			}
			fmt.Println(maxScore)
			return maxScore
		}
	}
}

func get7thNodeCounterClockwiseFrom(currentNode *Node) *Node {
	seventhNode := currentNode
	for i := 0; i < 7; i++ {
		seventhNode = seventhNode.left
	}
	return seventhNode
}

type Player struct {
	id    int
	score int
}

type Node struct {
	id    int
	left  *Node
	right *Node
}

func (n Node) String() string {
	return fmt.Sprintf("id: %d, left: %v, right: %v", n.id, n.left.id, n.right.id)
}

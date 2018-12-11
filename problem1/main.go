package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  file, err := os.Open("./input.txt")
  if err != nil {
    fmt.Println("Input file not found")
    return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
	sum := 0
  for scanner.Scan() {
    i, _ := strconv.Atoi(scanner.Text())
    sum += i
	}
	fmt.Println(sum)
}

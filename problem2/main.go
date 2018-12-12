package main

import (
  "fmt"
  "os"
  "strconv"
  "bufio"
)

func main() {

  file, err := os.Open("./input.txt")
  if err != nil {
    fmt.Println("Error in reading file")
    return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  input := []int{}
  for scanner.Scan() {
    a, _ := strconv.Atoi(scanner.Text())
    input = append(input, a)
  }

  fmt.Println(find_value(input))
}

func find_value(input []int) int {
  sum := 0
  sum_map := map[int]bool{}
  sum_map[sum] = true
  for {
    for _, a := range input {
      sum += a
      _, ok := sum_map[sum]
      if ok {
        return sum
      } else {
        sum_map[sum] = true
      }
    }
  }
}

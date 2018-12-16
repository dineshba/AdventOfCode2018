package main

import (
  "os"
  "fmt"
  "bufio"
)

func main() {
  file, err := os.Open("./input.txt")
  if err != nil {
    fmt.Println("Error reading file")
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  input := []map[rune]int{}
  for scanner.Scan() {
    a := scanner.Text()
    input = append(input, process(a) )
  }
  three_count := 0
  two_count := 0
  for _, element := range input {
    contains_three := false
    contains_two := false
    for _, value := range element {
      if value == 2 {
        contains_two = true
      }
      if value == 3 {
        contains_three = true
      }
    }
    if contains_three {
      three_count++
    }
    if contains_two {
      two_count++
    }
  }
  fmt.Println(three_count * two_count)
}

func process(input string) map[rune]int {
  temp := map[rune]int{}
  for _, i := range input {
    count, ok := temp[i]
    if ok {
      temp[i] = count + 1
    } else {
      temp[i] = 1
    }
  }
  return temp
}

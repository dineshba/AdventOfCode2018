package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Guard struct {
	Id                 int
	TotalAsleepMinutes int
	StartedSleepingAt  int
	WasSleeping        bool
	SleepingMap        map[int]int
  MinuteAtMaxAsleep  int
  CountOfAsleep      int
  }

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	guardMap := make(map[int]*Guard)
	inputs := []string{}
	for scanner.Scan() {
		a := scanner.Text()
		inputs = append(inputs, a)
	}

	sort.Slice(inputs, func(i, j int) bool {
		layout := "2006-01-02 15:04"
		first := strings.Split(inputs[i], "]")[0]
		second := strings.Split(inputs[j], "]")[0]
		first = strings.Split(first, "[")[1]
		second = strings.Split(second, "[")[1]
		t, _ := time.Parse(layout, first)
		t2, _ := time.Parse(layout, second)
		return t.Before(t2)
	})

	previousGuardId := 0
	for _, a := range inputs {
		guardId, startSleeping, minutes := process(a, previousGuardId)
		previousGuardId = guardId
		guard, ok := guardMap[guardId]
		if ok {
			if guard.WasSleeping && !startSleeping {
				duration := minutes - guard.StartedSleepingAt
				for i := guard.StartedSleepingAt; i < minutes; i++ {
					value, ok := guard.SleepingMap[i]
					if ok {
						guard.SleepingMap[i] = value + 1
					} else {
						guard.SleepingMap[i] = 1
					}
				}
				guard.TotalAsleepMinutes = guard.TotalAsleepMinutes + duration
        guard.WasSleeping = false
			} else {
				guard.StartedSleepingAt = minutes
				guard.WasSleeping = startSleeping
			}
		} else {
			guardMap[guardId] = &Guard{Id: guardId, TotalAsleepMinutes: 0, WasSleeping: false, SleepingMap: make(map[int]int)}
		}
	}
	var guardWithMaxAsleepMinutes *Guard
	for _, v := range guardMap {
		if guardWithMaxAsleepMinutes == nil {
			guardWithMaxAsleepMinutes = v
		}
		if v.TotalAsleepMinutes > guardWithMaxAsleepMinutes.TotalAsleepMinutes {
			guardWithMaxAsleepMinutes = v
		}
	}
  minuteAtMaxAsleep, _ := maxAsleepMinuteAndCount(*guardWithMaxAsleepMinutes)
	fmt.Println(minuteAtMaxAsleep * guardWithMaxAsleepMinutes.Id)
  fmt.Println("Problem two")
  mostFrequentAsleepOnSameMinute(guardMap)
}

func maxAsleepMinuteAndCount(guard Guard) (int, int) {
	minuteAtMaxAsleep := 0
	countOfAsleepAtAMinute := 0
	for k, v := range guard.SleepingMap {
		if v > countOfAsleepAtAMinute {
			countOfAsleepAtAMinute = v
			minuteAtMaxAsleep = k
		}
	}
  return minuteAtMaxAsleep, countOfAsleepAtAMinute
}

func mostFrequentAsleepOnSameMinute(guardMap map[int]*Guard) {
  var guardWithMaxCount *Guard
  for _, guard := range guardMap {
    minuteAtMaxAsleep, countOfAsleep := maxAsleepMinuteAndCount(*guard)
    guard.MinuteAtMaxAsleep = minuteAtMaxAsleep
    guard.CountOfAsleep = countOfAsleep
    if guardWithMaxCount == nil {
      guardWithMaxCount = guard
    }
    if guardWithMaxCount.CountOfAsleep < guard.CountOfAsleep {
      guardWithMaxCount = guard
    }
  }
  for k, v := range guardMap {
    fmt.Println(k, v.CountOfAsleep, v.MinuteAtMaxAsleep)
  }
  fmt.Println(guardWithMaxCount.Id * guardWithMaxCount.MinuteAtMaxAsleep)
}

func process(input string, previousGuardId int) (int, bool, int) {
	timeAndString := strings.Split(input, "] ")
	time := strings.Split(timeAndString[0], " ")[1]
	minutes, _ := strconv.Atoi(strings.Split(time, ":")[1])
	guardString := strings.Split(timeAndString[1], "#")
	var startedSleeping bool
	var guardId = previousGuardId
	if len(guardString) == 1 {
		startedSleeping = strings.Split(guardString[0], " ")[0] == "falls"
	} else {
		guardId, _ = strconv.Atoi(strings.Split(guardString[1], " ")[0])
	}
	return guardId, startedSleeping, minutes
}

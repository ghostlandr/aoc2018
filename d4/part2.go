package main

import (
	"fmt"
	"github.com/gholtslander-va/aoc2018/d4/util"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("/Users/gholtslander/go/src/github.com/gholtslander-va/aoc2018/d4/input.txt")
	stringData := strings.Split(string(data), "\n")

	guardNaps := util.ProcessInput(stringData)

	var longestGuard, highestMinute, highestMinuteAmount int
	for guardID, guardNap := range guardNaps {
		newHigh, value := util.GetHighestMinute(guardNap)
		if value > highestMinuteAmount {
			highestMinuteAmount = value
			highestMinute = newHigh
			longestGuard = guardID
		}
	}

	fmt.Println(longestGuard, highestMinute, highestMinuteAmount)
}

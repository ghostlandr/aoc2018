package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gholtslander-va/aoc2018/d4/util"
)

func main() {
	data, _ := ioutil.ReadFile("/Users/gholtslander/go/src/github.com/gholtslander-va/aoc2018/d4/input.txt")
	stringData := strings.Split(string(data), "\n")

	guardNaps := util.ProcessInput(stringData)

	var longestGuard int
	var longestNap time.Duration
	for guardID, guardNaps := range guardNaps {
		var guardNap time.Duration
		for _, n := range guardNaps {
			guardNap += n.End.Sub(n.Start)
		}
		if guardNap > longestNap {
			longestGuard = guardID
			longestNap = guardNap
		}
	}
	// Attempted 6040 (too low)
	// Attempting 6191 (too low)
	// Ended up being 20859
	minute, _ := util.GetHighestMinute(guardNaps[longestGuard])
	fmt.Println(longestGuard, longestNap, minute)
}

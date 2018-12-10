package main

import (
	"io/ioutil"
	"strings"

	"github.com/gholtslander-va/aoc2018/d7/util"
)

func main() {
	data, _ := ioutil.ReadFile("/Users/gholtslander/go/src/github.com/gholtslander-va/aoc2018/d7/input.txt")
	stringData := strings.Split(string(data), "\n")
	stringData = stringData[:len(stringData)-1]
	steps, stepMap := util.ProcessAllSteps(stringData)
	//input := "OCPUEFIXHRGWDZABTQJYMNKVSL"
	//fmt.Println("Building", input)
	util.ActuallyCompleteSteps(steps, stepMap)

	// Tried 420, too low
	//fmt.Println(time)
}

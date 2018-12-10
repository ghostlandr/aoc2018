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

	// Tried OPCUXEHFIRWZADGBTQYJMNKV (wrong)
	// Tried OCPUEFIXHRGWDZABTQJYMNKVSL (right)
	util.CompleteSteps(steps, stepMap)
}

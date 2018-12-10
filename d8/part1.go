package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//data, _ := ioutil.ReadFile("/Users/gholtslander/go/src/github.com/gholtslander-va/aoc2018/d8/input.txt")
	data := "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"
	stringData := strings.Split(string(data), "\n")
	nodeData := strings.Split(stringData[0], " ")

	nd := getMetaEntries(nodeData)
	fmt.Println(nd)
}

func getMetaEntries(inp []string) []string {
	children, metaEntries := inp[0], inp[1]
	i, _ := strconv.Atoi(children)
	totalMetaEntries := make([]string, 0)
	if i != 0 {
		totalMetaEntries = append(totalMetaEntries, getMetaEntries(inp[2:])...)
	}
	j, _ := strconv.Atoi(metaEntries)
	if j != 0 {
		totalMetaEntries = append(totalMetaEntries, inp[2:2+j]...)
	}
	return totalMetaEntries
}

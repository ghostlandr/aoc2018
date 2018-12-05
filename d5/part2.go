package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("/Users/gholtslander/go/src/github.com/gholtslander-va/aoc2018/d5/input.txt")
	stringData := strings.Split(string(data), "\n")

	for s := range stringData {
		fmt.Println(s)
	}
}

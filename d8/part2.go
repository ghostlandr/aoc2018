package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("/Users/gholtslander/go/src/github.com/gholtslander-va/aoc2018/d8/input.txt")
	stringData := strings.Split(string(data), "\n")

	for _, s := range stringData {
		fmt.Println(s)
	}
}

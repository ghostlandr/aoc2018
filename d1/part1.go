package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	start := 0
	dat, err := ioutil.ReadFile("/Users/gholtslander/go/src/github.com/gholtslander-va/aoc2018/d1/p1input.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	sInput := strings.Split(string(dat), "\n")
	for _, in := range sInput {
		intIn, _ := strconv.Atoi(in)
		start += intIn
	}
	fmt.Print(start)
}

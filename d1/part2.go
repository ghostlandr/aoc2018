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
	freqs := make(map[int]bool)
	for {
		for _, in := range sInput {
			intIn, err := strconv.Atoi(in)
			if err != nil {
				continue
			}
			fmt.Println(intIn)
			if _, ok := freqs[start]; ok {
				fmt.Printf("Second time seeing %d after %d reps\n", start, len(freqs))
				return
			}
			freqs[start] = true
			start += intIn
		}
	}
	//fmt.Print(start)
}

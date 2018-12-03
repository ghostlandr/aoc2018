package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("/Users/gholtslander/go/src/github.com/gholtslander-va/aoc2018/d2/input.txt")
	sInput := strings.Split(string(dat), "\n")

	for _, line := range sInput {
		//fmt.Println(line)
		lineChars := strings.Split(line, "")
		for _, innerLine := range sInput {
			diffs := 0
			innerLineChars := strings.Split(innerLine, "")
			for i, char := range lineChars {
				if len(innerLineChars) == 0 {
					break
				}
				if char != innerLineChars[i] {
					diffs++
				}
			}
			if diffs == 1 {
				fmt.Printf("Possible match: %s %s\n", line, innerLine)
			}
		}
	}
}


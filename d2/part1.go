package main

import (
"fmt"
"io/ioutil"
"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("/Users/gholtslander/go/src/github.com/gholtslander-va/aoc2018/d2/input.txt")
	sInput := strings.Split(string(dat), "\n")

	var twos, threes int
	for _, line := range sInput {
		fmt.Println(line)
		m := make(map[string]int)
		for _, char := range strings.Split(line, "") {
			m[char]++
		}
		var hasTwo, hasThree bool
		for char, count := range m {
			fmt.Printf("%s: %d\n", char, count)
			if count == 2 {
				if hasTwo {
					continue
				}
				twos++
				hasTwo = true
			}
			if count == 3 {
				if hasThree {
					continue
				}
				threes++
				hasThree = true
			}
		}
	}
	// Submitted answers: 18357, 7163 (correct)
	fmt.Printf("Twos: %d, threes: %d. Checksum: %d", twos, threes, twos*threes)
}

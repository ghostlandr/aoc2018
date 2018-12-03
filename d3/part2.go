
package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("/Users/gholtslander/go/src/github.com/gholtslander-va/aoc2018/d3/input.txt")
	stringData := strings.Split(string(data), "\n")

	fab := make([][]int, 1000)
	for i := range fab {
		fab[i] = make([]int, 1000)
	}
	r := regexp.MustCompile(`#\d+ @ (\d+),(\d+): (\d+)x(\d+)`)
	/*
	A claim like #123 @ 3,2: 5x4 means that claim ID 123 specifies a rectangle
	3 inches from the left edge,
	2 inches from the top edge,
	5 inches wide, and
	4 inches tall.
	*/
	for _, c := range stringData {
		if c == "" {
			continue
		}
		details := r.FindStringSubmatch(c)
		fromLeft, fromTop, wide, tall := details[1], details[2], details[3], details[4]
		iFromLeft, _ := strconv.Atoi(fromLeft)
		iFromTop, _ := strconv.Atoi(fromTop)
		iWide, _ := strconv.Atoi(wide)
		iTall, _ := strconv.Atoi(tall)

		//fmt.Println(iFromLeft, iFromTop, iWide, iTall)
		for i := 0; i < iTall; i++ {
			for j := 0; j < iWide; j++ {
				//fmt.Printf("Incrementing %d x %d\n", iFromTop + i, iFromLeft + j)
				fab[iFromTop + i][iFromLeft + j]++
			}
		}
		//fmt.Println(c)
	}

	r2 := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
	for _, c := range stringData {
		if c == "" {
			continue
		}
		details := r2.FindStringSubmatch(c)
		id, fromLeft, fromTop, wide, tall := details[1], details[2], details[3], details[4], details[5]
		iFromLeft, _ := strconv.Atoi(fromLeft)
		iFromTop, _ := strconv.Atoi(fromTop)
		iWide, _ := strconv.Atoi(wide)
		iTall, _ := strconv.Atoi(tall)

		intacts := 0
		totalArea := iWide * iTall
		//fmt.Println(iFromLeft, iFromTop, iWide, iTall)
		for i := 0; i < iTall; i++ {
			for j := 0; j < iWide; j++ {
				//fmt.Printf("Incrementing %d x %d\n", iFromTop + i, iFromLeft + j)
				if fab[iFromTop + i][iFromLeft + j] == 1 {
					intacts++
				}
			}
		}
		if intacts == totalArea {
			fmt.Printf("%s is in tact!", id)
		}
		//fmt.Println(c)
	}
}

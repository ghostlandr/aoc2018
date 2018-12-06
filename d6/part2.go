package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gholtslander-va/aoc2018/d6/util"
)

func main() {
	data, _ := ioutil.ReadFile("/Users/gholtslander/go/src/github.com/gholtslander-va/aoc2018/d6/input.txt")
	stringData := strings.Split(string(data), "\n")
	stringData = stringData[:len(stringData)-1]

	grid := util.CreateEmptyGrid(400)
	for i, s := range stringData {
		i++
		c2, c1 := util.GetCoordsFromString(s)
		grid[c1][c2] = util.NewManHat(strconv.Itoa(i), 0, util.IsCoordInfinite(c1, c2), true)
	}

	var closeOnes int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			dist := util.GetManhattanDistanceToAllPoints(x, y, stringData)
			if dist < 10000 {
				closeOnes++
			}
		}
	}
	// Correct: 44667
	fmt.Println("Answer: ", closeOnes)
}

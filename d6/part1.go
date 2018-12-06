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
	// Trim the extra newline
	stringData = stringData[:len(stringData)-1]

	grid := util.CreateEmptyGrid(400)
	for i, s := range stringData {
		i++
		c2, c1 := util.GetCoordsFromString(s)
		fmt.Println("Placing point at ", c1, c2)
		grid[c1][c2] = util.NewManHat(strconv.Itoa(i), 0, util.IsCoordInfinite(c1, c2), true)
	}

	for i, s := range stringData {
		i++
		c1, c2 := util.GetCoordsFromString(s)
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				dist := util.GetManhattanDistance(c1, x, c2, y)
				point := grid[y][x]
				if point.IsHomeCoord {
					continue
				}
				if point.DistanceToClosest == 0 {
					grid[y][x] = util.NewManHat(strconv.Itoa(i), dist, util.IsCoordInfinite(y, x), false)
					continue
				}
				// If the distance between these two points is less than the current closest, replace it
				if dist < point.DistanceToClosest {
					grid[y][x].DistanceToClosest = dist
					grid[y][x].ClosestCoord = strconv.Itoa(i)
					continue
				}
				// Need to think about equidistant points soon? yes, now.
				if point.DistanceToClosest == dist {
					grid[y][x] = util.NewManHat(".", dist, util.IsCoordInfinite(y, x), false)
				}
			}
		}
	}

	for i, g := range grid {
		fmt.Println(i, g)
	}

	answer := make(map[string]int, 50)
	infinite := make(map[string]bool, 50)
	for i := range grid {
		for _, ig := range grid[i] {
			answer[ig.ClosestCoord]++
			if ig.IsInfinite {
				// If any of its coords is infinite then the whole thing is infinite
				infinite[ig.ClosestCoord] = true
			}
		}
	}
	fmt.Println(answer)
	fmt.Println(infinite)
	var biggestCoord string
	var biggestArea int
	for coord, area := range answer {
		if area > biggestArea {
			if _, ok := infinite[coord]; !ok {
				biggestArea = area
				biggestCoord = coord
			}
		}
	}
	// Attempts:
	// Too high: 5916
	// Trying: 3871 - Correct!
	fmt.Println("Answer: ", biggestCoord, biggestArea)
	//fmt.Printf("%+v", grid[0][0])
	/*
		Loop over each coord and add them to the grid
		For each coordinate,
		Walk the grid and calculate the manhattan distance between that coord and the original
		If there is already a point there
	*/
}

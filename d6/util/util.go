package util

import (
	"math"
	"regexp"
	"strconv"
)

type ManHat struct {
	ClosestCoord      string
	DistanceToClosest int
	IsInfinite        bool
	IsHomeCoord       bool
	// What about equidistant points? Equidistant points will have a . for ClosestCoord
}

func (m ManHat) String() string {
	if m.IsHomeCoord {
		return "H" + m.ClosestCoord
	}
	return m.ClosestCoord
}

func NewManHat(closest string, distance int, isInfinite, isHomeCoord bool) ManHat {
	return ManHat{
		ClosestCoord:      closest,
		DistanceToClosest: distance,
		IsInfinite:        isInfinite,
		IsHomeCoord:       isHomeCoord,
	}
}

func IsCoordInfinite(c1, c2 int) bool {
	if c1 == 0 || c2 == 0 {
		return true
	}
	if c1 == 399 || c2 == 399 {
		return true
	}
	return false
}

func GetManhattanDistanceToAllPoints(p1, p2 int, coordData []string) int {
	var total int
	for _, coord := range coordData {
		c1, c2 := GetCoordsFromString(coord)
		total += GetManhattanDistance(p1, c1, p2, c2)
	}
	return total
}

func GetManhattanDistance(p1, q1, p2, q2 int) int {
	return int(math.Abs(float64(p1)-float64(q1)) + math.Abs(float64(p2)-float64(q2)))
}

var coordReg = regexp.MustCompile(`(\d+), (\d+)`)

func GetCoordsFromString(coords string) (int, int) {
	matches := coordReg.FindStringSubmatch(coords)
	c1, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	c2, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}
	return c1, c2
}

func CreateEmptyGrid(size int) [][]ManHat {
	grid := make([][]ManHat, size)
	for i := range grid {
		grid[i] = make([]ManHat, size)
	}
	return grid
}

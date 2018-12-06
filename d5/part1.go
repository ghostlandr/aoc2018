package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gholtslander-va/aoc2018/d5/util"
)

func main() {
	data, _ := ioutil.ReadFile("/Users/gholtslander/go/src/github.com/gholtslander-va/aoc2018/d5/input.txt")
	stringData := strings.Split(string(data), "\n")
	polymerData := strings.Split(stringData[0], "")

	final := polymerData
	var removed []string
	var prevFinal string
	for {
		final, removed = util.ReactPolymers(final, util.PolaritiesReact)
		fmt.Println("Removed elements: ", strings.Join(removed, ""))
		if prevFinal == strings.Join(final, "") {
			break
		}
		fmt.Println(len(final), strings.Join(final, ""))
		// Keep going until no reacting polymer chains remainnnnnnn
		prevFinal = strings.Join(final, "")
	}

	// Tried 10600, too high!
	// Omg, 10,598 was right
	fmt.Println("Final:", len(strings.Join(final, "")), strings.Join(final, ""))
}

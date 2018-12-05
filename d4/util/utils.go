package util

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var DateReg = regexp.MustCompile(`\[(.+)\]`)

type Shifts []string

func (s Shifts) Len() int {
	return len(s)
}

func (s Shifts) Less(i, j int) bool {
	if s[i] == "" {
		return true
	}
	if s[j] == "" {
		return false
	}
	iTime := DateReg.FindStringSubmatch(s[i])
	iT, _ := time.Parse("2006-01-02 15:04", iTime[1])

	jTime := DateReg.FindStringSubmatch(s[j])
	jT, _ := time.Parse("2006-01-02 15:04", jTime[1])
	return iT.Before(jT)
}

func (s Shifts) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var IdReg = regexp.MustCompile(`\[.+\] Guard #(\d+) begins shift`)

type Nap struct {
	Start time.Time
	End   time.Time
}

func ProcessInput(stringData []string) map[int][]Nap {
	shiftData := Shifts(stringData)
	sort.Sort(shiftData)

	guardID := 0
	guardNaps := make(map[int][]Nap)
	for _, s := range shiftData {
		// Get guard ID if it's there
		if strings.Contains(s, "Guard") {
			matches := IdReg.FindStringSubmatch(s)
			guardID, _ = strconv.Atoi(matches[1])
			if _, ok := guardNaps[guardID]; !ok {
				guardNaps[guardID] = make([]Nap, 0)
			}
			continue
		}
		// If no guard id, log events to the guard id we have
		if strings.Contains(s, "falls") {
			t := GetTimeFromLog(s)
			guardNaps[guardID] = append(guardNaps[guardID], Nap{Start: t})
		}
		if strings.Contains(s, "wakes") {
			t := GetTimeFromLog(s)
			guardNaps[guardID][len(guardNaps[guardID])-1].End = t
		}
	}
	return guardNaps
}

func GetHighestMinute(n []Nap) (int, int) {
	d := make(map[int]int)
	for _, na := range n {
		sMin := na.Start.Minute()
		for i := sMin; i < na.End.Minute(); i++ {
			d[i]++
		}
	}
	var max int
	var maxMin int
	for min, val := range d {
		if val > max {
			max = val
			maxMin = min
		}
	}
	return maxMin, max
}

func GetTimeFromLog(entry string) time.Time {
	iTime := DateReg.FindStringSubmatch(entry)
	iT, _ := time.Parse("2006-01-02 15:04", iTime[1])
	return iT
}

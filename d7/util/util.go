package util

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/gholtslander-va/chat-webhook-responder-go/timer"
)

type Step struct {
	Letter   string
	Complete bool
	Next     []string
}

type Steps []*Step

func (s Steps) Len() int {
	return len(s)
}

func (s Steps) Less(i, j int) bool {
	return s[i].Letter < s[j].Letter
}

func (s Steps) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func CompleteSteps(steps []*Step, stepMap map[string]*Step) {
	completed := make([]*Step, 0, 50)
	for {
		//fmt.Println(len(completed), len(steps))
		//if len(completed) >= len(steps) {
		//	break
		//}
		completable := make(Steps, 0)
		for _, s := range steps {
			if stepMap[s.Letter].Complete {
				// Step already complete
				fmt.Println(s.Letter, "is already complete")
				continue
			}
			// Check for steps that can be completed
			isCompletable := true
			for _, n := range s.Next {
				if !stepMap[n].Complete {
					// If any of them aren't complete, not completable yet.
					isCompletable = false
					break
				}
			}
			if isCompletable {
				completable = append(completable, s)
			}
		}
		sort.Sort(completable)
		for _, c := range completable {
			stepMap[c.Letter].Complete = true
			completed = append(completed, c)
			// Break to try and see if there is anything better that is completable
			break
		}
		for letter, step := range stepMap {
			fmt.Println(letter, fmt.Sprintf("%+v", *step))
		}
		fmt.Println(PrintSteps(completed))
	}
}

func ActuallyCompleteSteps(steps []*Step, stepMap map[string]*Step) {
	//completed := make([]*Step, 0, 50)
	doStep := make(chan string, len(steps))
	//for i := 0; i < len(sSteps); i++ {
	//	doStep <- sSteps[i]
	//}
	//close(doStep)
	tim := timer.New()
	tim.Start("elfs")
	dones := make(chan string)
	go elf(doStep, dones, 1)
	go elf(doStep, dones, 2)
	go elf(doStep, dones, 3)
	go elf(doStep, dones, 4)
	go elf(doStep, dones, 5)

	started := make(map[string]bool)
	go func() {
		select {
		case stepDone := <-dones:
			fmt.Println(stepDone, "is done!")
			stepMap[stepDone].Complete = true
		default:
			//fmt.Println("No activity")
		}
	}()

	count := 0
	for {
		for _, s := range steps {
			if stepMap[s.Letter].Complete {
				// Step already complete
				fmt.Println(s.Letter, "is already complete")
				continue
			}
			// Check for steps that can be completed
			isCompletable := true
			for _, n := range s.Next {
				if !stepMap[n].Complete {
					// If any of them aren't complete, not completable yet.
					isCompletable = false
					break
				}
			}
			if isCompletable {
				if _, ok := started[s.Letter]; !ok {
					doStep <- s.Letter
					started[s.Letter] = true
				}
			}
		}
		count++
		time.Sleep(20 * time.Millisecond)
		if count > 1000 {
			break
		}
	}
	tim.End("elfs")
	fmt.Println(tim.Elapsed("elfs"))
}

func elf(stepC <-chan string, buildDone chan string, id int) {
	//for {
	select {
	case step := <-stepC:
		bTime := buildSeconds(step)
		fmt.Println(id, "is building part", step, "it'll take", bTime)
		time.Sleep(bTime)
		fmt.Println("Emitting")
		buildDone <- step
	default:
	}
	//step, more := <-stepC
	//if !more {
	//	done <- true
	//	return
	//}

	//}
}

func PrintSteps(c []*Step) (o string) {
	for _, s := range c {
		o += s.Letter
	}
	return
}

var sReg = regexp.MustCompile(`Step ([A-Z]) must be finished before step ([A-Z]) can begin.`)

func ProcessAllSteps(stepData []string) ([]*Step, map[string]*Step) {
	steps := make([]*Step, 0, len(stepData))
	stepMap := make(map[string]*Step, 0)
	for _, s := range stepData {
		next, letter := ProcessInput(s)
		if _, ok := stepMap[next]; !ok {
			newStep := &Step{Letter: next, Next: []string{}}
			steps = append(steps, newStep)
			stepMap[next] = newStep
		}
		if _, ok := stepMap[letter]; !ok {
			newStep := &Step{Letter: letter, Next: []string{next}}
			steps = append(steps, newStep)
			stepMap[letter] = newStep
			continue
		}
		step := stepMap[letter]
		if !strings.Contains(strings.Join(step.Next, ""), next) {
			step.Next = append(step.Next, next)
		}
	}
	return steps, stepMap
}

func ProcessInput(step string) (string, string) {
	// Need to get smarter here, as there are multiple "previous" steps required for steps.
	stepParts := sReg.FindStringSubmatch(step)
	// Complete autoinits to false
	return stepParts[1], stepParts[2]
}

func GetBuildTime(steps string) (time.Duration, error) {
	sSteps := strings.Split(steps, "")
	s := make(chan string, len(sSteps))
	for i := 0; i < len(sSteps); i++ {
		s <- sSteps[i]
	}
	close(s)
	done := make(chan bool)
	tim := timer.New()
	tim.Start("elfs")
	//go elf(s, done, 1)
	//go elf(s, done, 2)
	//go elf(s, done, 3)
	//go elf(s, done, 4)
	//go elf(s, done, 5)
	dones := 1
	for {
		select {
		case <-done:
			dones++
		}
		if dones == 5 {
			break
		}
	}
	tim.End("elfs")
	fmt.Println(tim.Elapsed("elfs"))
	return tim.Elapsed("elfs")
}

func buildSeconds(letter string) time.Duration {
	return (60 + letterToSeconds[letter]) * time.Millisecond
}

var letterToSeconds = map[string]time.Duration{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}

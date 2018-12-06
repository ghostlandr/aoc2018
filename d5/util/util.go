package util

import (
	"strings"
)

type ReactionFunction func(string, string) bool

func CleansePolymerStrain(polymerData []string, strain string) (cleansed []string) {
	lowStrain := strings.ToLower(strain)
	for _, s := range polymerData {
		if strings.ToLower(s) != lowStrain {
			cleansed = append(cleansed, s)
		}
	}
	return
}

func PolaritiesReact(l, m string) bool {
	if l == m {
		// Same case, same polarity, no reaction
		return false
	}
	if strings.ToLower(l) == strings.ToLower(m) {
		// Same letter, but different cases... reaction!
		return true
	}
	// Different letters presumably, no reaction
	return false
}

func ReactPolymers(polymerData []string, rFunc ReactionFunction) ([]string, []string) {
	//var current string
	// Check the current letter against the next letter in the chain. If they don't react, add the current letter
	// to the final []string
	var final, removed []string

	for i := 0; i < len(polymerData); i++ {
		if len(final) == 0 {
			final = append(final, polymerData[i])
			continue
		}
		if rFunc(polymerData[i], final[len(final)-1]) {
			// Trim off the last
			final = final[:len(final)-1]
			continue
		}
		// If it isn't the first loop and the polarities didn't react, append it
		final = append(final, polymerData[i])
	}

	return final, removed
}

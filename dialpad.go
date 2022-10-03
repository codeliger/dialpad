package main

var lookup = map[rune][]rune{
	'2': {'A', 'B', 'C'},
	'3': {'D', 'E', 'F'},
	'4': {'G', 'H', 'I'},
	'5': {'J', 'K', 'L'},
	'6': {'M', 'N', 'O'},
	'7': {'P', 'Q', 'R', 'S'},
	'8': {'T', 'U', 'V'},
	'9': {'W', 'X', 'Y', 'Z'},
}

// num2strings takes a number and looksup the corresponding dialpad characters and returns all possible strings
func num2strings(number string, lookup map[rune][]rune) []string {
	numberLetterGroups := [][]rune{}
	for _, r := range number {
		if letterGroup, ok := lookup[r]; ok {
			numberLetterGroups = append(numberLetterGroups, letterGroup)
		}
	}

	numberLetterIndexes := make([]int, len(numberLetterGroups))

	for i := range numberLetterIndexes {
		numberLetterIndexes[i] = 0
	}

	combinations := []string{}

	numberIndex := len(numberLetterIndexes) - 1
	for {
		combo := []rune{}
		for groupIndex, letterIndex := range numberLetterIndexes {
			combo = append(combo, numberLetterGroups[groupIndex][letterIndex])
		}
		combinations = append(combinations, string(combo))

		if numberLetterIndexes[numberIndex] < len(numberLetterGroups[numberIndex])-1 {
			numberLetterIndexes[numberIndex]++
		} else {
			numberLetterIndexes[numberIndex] = 0
			valueSet := false
			for numberIndex > 0 && !valueSet {
				numberIndex--
				if numberLetterIndexes[numberIndex] == len(numberLetterGroups[numberIndex])-1 {
					numberLetterIndexes[numberIndex] = 0
				} else {
					numberLetterIndexes[numberIndex]++
					valueSet = true
				}
			}
			if !valueSet {
				return combinations
			} else {
				numberIndex = len(numberLetterIndexes) - 1
			}
		}
	}
}

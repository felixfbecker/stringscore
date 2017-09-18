// Based on https://github.com/Microsoft/vscode/blob/321cec5618ce19067ebeb187a782329f028957aa/src/vs/base/common/scorer.ts

package stringscore

import (
	"strings"
	"unicode"
)

var wordPathBoundary = [...]byte{'-', '_', ' ', '/', '\\', '.'}

// Score computes a score for the given string and the given query.
//
// Rules:
// Character score: 1
// Same case bonus: 1
// Upper case bonus: 1
// Consecutive match bonus: 5
// Start of word/path bonus: 7
// Start of string bonus: 8
func Score(target string, query string) int {

	if target == "" || query == "" {
		return 0 // return early if target or query are undefined
	}

	queryLen := len(query)
	targetLower := strings.ToLower(target)
	queryLower := strings.ToLower(query)

	index := 0
	startAt := 0
	score := 0

	for index < queryLen {

		indexOf := strings.IndexByte(targetLower[startAt:], queryLower[index]) + startAt

		if indexOf == -1 {
			score = 0 // This makes sure that the query is contained in the target
			break
		}

		// Character match bonus
		score += 1

		// Consecutive match bonus
		if startAt == indexOf {
			score += 5
		}

		// Same case bonus
		if indexOf < queryLen && target[indexOf] == query[indexOf] {
			score += 1
		}

		// Start of word bonus
		if indexOf == 0 {
			score += 8
		} else {
			// After separator bonus
			for _, w := range wordPathBoundary {
				if w == target[indexOf-1] {
					score += 7
					goto next
				}
			}
		}
		if unicode.IsUpper(rune(target[indexOf])) {
			// Inside word upper case bonus
			score += 1
		}

	next:
		startAt = indexOf + 1
		index++
	}

	return score
}

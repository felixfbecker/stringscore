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

	targetLower := strings.ToLower(target)
	queryLower := strings.ToLower(query)

	startAt := 0
	score := 0

	for queryIdx := range query {
		targetIdx := strings.IndexByte(targetLower[startAt:], queryLower[queryIdx]) + startAt

		if targetIdx == -1 {
			score = 0 // This makes sure that the query is contained in the target
			break
		}

		// Character match bonus
		score += 1

		// Consecutive match bonus
		if startAt == targetIdx {
			score += 5
		}

		// Same case bonus
		if target[targetIdx] == query[queryIdx] {
			score += 1
		}

		// Start of word bonus
		if targetIdx == 0 {
			score += 8
		} else {
			// After separator bonus
			for _, w := range wordPathBoundary {
				if w == target[targetIdx-1] {
					score += 7
					goto next
				}
			}
		}
		if unicode.IsUpper(rune(target[targetIdx])) {
			// Inside word upper case bonus
			score += 1
		}

	next:
		startAt = targetIdx + 1
	}

	return score
}

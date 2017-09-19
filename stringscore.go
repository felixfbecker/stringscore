// Based on https://github.com/Microsoft/vscode/blob/321cec5618ce19067ebeb187a782329f028957aa/src/vs/base/common/scorer.ts

package stringscore

import (
	"strings"
	"unicode"
)

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

	if len(query) > len(target) {
		return 0 // impossible for query to be a substring
	}

	targetRunes := []rune(target)

	startAt := 0
	score := 0

	for _, queryC := range query {
		targetIdx := runeIndexLower(targetRunes[startAt:], unicode.ToLower(queryC)) + startAt

		if targetIdx < startAt {
			score = 0 // This makes sure that the query is contained in the target
			break
		}

		// Character match bonus
		score++

		// Consecutive match bonus
		if targetIdx == startAt {
			score += 5
		}

		// Same case bonus
		if targetRunes[targetIdx] == queryC {
			score++
		}

		// Start of word bonus
		if targetIdx == 0 {
			score += 8
		} else if isWordSeparator(targetRunes[targetIdx-1]) {
			// After separator bonus
			score += 7
		} else if unicode.IsUpper(targetRunes[targetIdx]) {
			// Inside word upper case bonus
			score++
		}

		startAt = targetIdx + 1
	}
	return score
}

const wordPathBoundary = "-_ /\\."

func isWordSeparator(r rune) bool {
	return strings.IndexRune(wordPathBoundary, r) >= 0
}

func runeIndexLower(s []rune, r rune) int {
	for i, s := range s {
		if unicode.ToLower(s) == r {
			return i
		}
	}
	return -1
}

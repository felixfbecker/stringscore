package stringscore_test

import (
	"sort"
	"testing"

	"github.com/felixfbecker/stringscore"
)

func TestScore(t *testing.T) {
	target := "HeLlo-World"

	scores := make([]int, 0, 0)
	scores = append(scores, stringscore.Score(target, "HelLo-World")) // direct case match
	scores = append(scores, stringscore.Score(target, "hello-world")) // direct mix-case match
	scores = append(scores, stringscore.Score(target, "HW"))          // direct case prefix (multiple)
	scores = append(scores, stringscore.Score(target, "hw"))          // direct mix-case prefix (multiple)
	scores = append(scores, stringscore.Score(target, "H"))           // direct case prefix
	scores = append(scores, stringscore.Score(target, "h"))           // direct mix-case prefix
	scores = append(scores, stringscore.Score(target, "W"))           // direct case word prefix
	scores = append(scores, stringscore.Score(target, "w"))           // direct mix-case word prefix
	scores = append(scores, stringscore.Score(target, "Ld"))          // in-string case match (multiple)
	scores = append(scores, stringscore.Score(target, "ld"))          // in-string mix-case match
	scores = append(scores, stringscore.Score(target, "L"))           // in-string case match
	scores = append(scores, stringscore.Score(target, "l"))           // in-string mix-case match
	scores = append(scores, stringscore.Score(target, "4"))           // no match

	// Assert scoring order
	sortedScores := make([]int, len(scores))
	copy(sortedScores, scores)
	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})

	for i, score := range scores {
		if score != sortedScores[i] {
			t.Errorf("Scored array was incorrect\ngot:  %v\nwant: %v", sortedScores, scores)
			break
		}
	}
}

func TestZeroScoreOnEmptyTarget(t *testing.T) {
	score := stringscore.Score("", "foo")
	if score != 0 {
		t.Errorf("Expected empty target to result in score of zero, got: %v", score)
	}
}

func TestZeroScoreOnEmptyQuery(t *testing.T) {
	score := stringscore.Score("foo", "")
	if score != 0 {
		t.Errorf("Expected empty query to result in score of zero, got: %v", score)
	}
}

func TestZeroScoreOnQueryLongerThanTarget(t *testing.T) {
	score := stringscore.Score("foo", "foobar")
	if score != 0 {
		t.Errorf("Expected query longer than target to result in score of zero, got: %v", score)
	}
}

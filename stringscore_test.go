package stringscore_test

import (
	"sort"
	"testing"

	"github.com/felixfbecker/stringscore"
)

func TestScore(t *testing.T) {
	target := "H❄Ll❄-World"

	scores := []int{
		stringscore.Score(target, "H❄lL❄-World"), // direct case match
		stringscore.Score(target, "h❄ll❄-world"), // direct mix-case match
		stringscore.Score(target, "HW"),          // direct case prefix (multiple)
		stringscore.Score(target, "hw"),          // direct mix-case prefix (multiple)
		stringscore.Score(target, "H"),           // direct case prefix
		stringscore.Score(target, "h"),           // direct mix-case prefix
		stringscore.Score(target, "W"),           // direct case word prefix
		stringscore.Score(target, "w"),           // direct mix-case word prefix
		stringscore.Score(target, "Ld"),          // in-string case match (multiple)
		stringscore.Score(target, "ld"),          // in-string mix-case match
		stringscore.Score(target, "L"),           // in-string case match
		stringscore.Score(target, "l"),           // in-string mix-case match
		stringscore.Score(target, "4"),           // no match
	}

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

func TestExactMatchIsPrefferedOverFuzzyMatch(t *testing.T) {
	query := "backend"
	fuzzyScore := stringscore.Score("vendor/github.com/gorilla/websocket/conn_read_legacy.go", query)
	exactScore := stringscore.Score("pkg/backend/trace.go", query)
	if fuzzyScore >= exactScore {
		t.Errorf("Expected a fuzzy match to have a lower score than an exact match, fuzzy: %v, exact: %v", fuzzyScore, exactScore)
	}
}

func TestExactMatchIsPrefferedOverSubstring(t *testing.T) {
	query := "react-router"
	substringScore := stringscore.Score("react-router-cdn", query)
	exactScore := stringscore.Score("react-router", query)
	if substringScore >= exactScore {
		t.Errorf("Expected a substring match to have a lower score than an exact match, substring: %v, exact: %v", substringScore, exactScore)
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

func BenchmarkScoreASCII(b *testing.B) {
	query := "backend"
	target := "vendor/github.com/gorilla/websocket/conn_read_legacy.go"
	for n := 0; n <= b.N; n++ {
		score := stringscore.Score(target, query)
		if score <= 0 {
			b.Fatal("Expected a match")
		}
	}
}

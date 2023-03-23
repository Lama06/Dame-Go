package dame_test

import (
	"fmt"
	"testing"

	"github.com/Lama06/Dame-Go.git/dame"
)

func brett(zeilen ...string) dame.Brett {
	brett, err := dame.ParseBrett(zeilen...)
	if err != nil {
		panic(fmt.Errorf("failed to parse brett: %w", err))
	}
	return brett
}

func setAndSliceEqual[T comparable](slice []T, set map[T]struct{}) bool {
	if len(slice) != len(set) {
		return false
	}

	for _, value := range slice {
		if _, ok := set[value]; !ok {
			return false
		}
	}

	return true
}

func testMoves(t *testing.T, start dame.Brett, spieler dame.Spieler, expectedMoves ...dame.Brett) {
	moves := start.PossibleMoves(spieler)

	if !setAndSliceEqual(expectedMoves, moves) {
		t.Log("expected:\n", expectedMoves)
		t.Log("got:\n", moves)
		t.FailNow()
	}
}

func TestSteinBewegen(t *testing.T) {
	testMoves(
		t,
		brett(
			"_ _ a _ ",
			" _ _ A _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" b _ _ _",
		),
		dame.SpielerUnten,
		brett(
			"_ _ a _ ",
			" _ _ A _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ b _ _ ",
			" _ _ _ _",
		),
		brett(
			"_ _ a _ ",
			" _ _ A _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ _ _",
			"b _ _ _ ",
			" _ _ _ _",
		),
	)
}

func TestSteinSchlagen(t *testing.T) {
	testMoves(
		t,
		brett(
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ A a ",
			" _ b a _",
			"_ _ _ a ",
			" _ _ _ _",
		),
		dame.SpielerUnten,
		brett(
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ a _",
			"_ _ _ _ ",
			" _ _ b _",
		),
	)
}

func TestDameBewegen(t *testing.T) {
	testMoves(
		t,
		brett(
			"b _ _ _ ",
			" _ _ _ b",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ A _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" b _ _ b",
		),
		dame.SpielerOben,
		// Nach unten rechts
		brett(
			"b _ _ _ ",
			" _ _ _ b",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ A _",
			"_ _ _ _ ",
			" b _ _ b",
		),
		brett(
			"b _ _ _ ",
			" _ _ _ b",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ A ",
			" b _ _ b",
		),
		// Nach unten links
		brett(
			"b _ _ _ ",
			" _ _ _ b",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ A _ _",
			"_ _ _ _ ",
			" b _ _ b",
		),
		brett(
			"b _ _ _ ",
			" _ _ _ b",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ A _ _ ",
			" b _ _ b",
		),
		// Nach oben rechts
		brett(
			"b _ _ _ ",
			" _ _ _ b",
			"_ _ _ _ ",
			" _ _ A _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" b _ _ b",
		),
		brett(
			"b _ _ _ ",
			" _ _ _ b",
			"_ _ _ A ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" b _ _ b",
		),
		// Nach oben links
		brett(
			"b _ _ _ ",
			" _ _ _ b",
			"_ _ _ _ ",
			" _ A _ _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" b _ _ b",
		),
		brett(
			"b _ _ _ ",
			" _ _ _ b",
			"_ A _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" b _ _ b",
		),
		brett(
			"b _ _ _ ",
			" A _ _ b",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" b _ _ b",
		),
	)
}

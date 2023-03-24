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

func testMoves(t *testing.T, start dame.Brett, spieler dame.Spieler, expected dame.PossibleMoves) {
	got := start.PossibleMovesForSpieler(spieler)

	if !got.Equals(expected) {
		t.Log("expected:\n", expected)
		t.Log("got:\n", got)
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
		dame.PossibleMoves{
			dame.Move{
				Start: dame.Position{Zeile: 7, Spalte: 1},
				Ende:  dame.Position{Zeile: 6, Spalte: 2},
				Steps: dame.MoveSteps{
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
				},
			},
			dame.Move{
				Start: dame.Position{Zeile: 7, Spalte: 1},
				Ende:  dame.Position{Zeile: 6, Spalte: 0},
				Steps: dame.MoveSteps{
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
				},
			},
		},
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
		dame.PossibleMoves{
			dame.Move{
				Start: dame.Position{Zeile: 5, Spalte: 3},
				Ende:  dame.Position{Zeile: 7, Spalte: 5},
				Steps: dame.MoveSteps{
					brett(
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ b _",
						"_ _ _ a ",
						" _ _ a _",
						"_ _ _ a ",
						" _ _ _ _",
					),
					brett(
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ a b",
						"_ _ _ a ",
						" _ _ _ _",
					),
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
				},
			},
		},
	)

	testMoves(
		t,
		brett(
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ a _ ",
			" b b b _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ _ _",
		),
		dame.SpielerOben,
		dame.PossibleMoves{
			dame.Move{
				Start: dame.Position{Zeile: 2, Spalte: 4},
				Ende:  dame.Position{Zeile: 4, Spalte: 6},
				Steps: dame.MoveSteps{
					brett(
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" b b _ _",
						"_ _ _ a ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ _ _",
					),
				},
			},
			dame.Move{
				Start: dame.Position{Zeile: 2, Spalte: 4},
				Ende:  dame.Position{Zeile: 2, Spalte: 0},
				Steps: dame.MoveSteps{
					brett(
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" b _ b _",
						"_ a _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ _ _",
					),
					brett(
						"_ _ _ _ ",
						" _ _ _ _",
						"a _ _ _ ",
						" _ _ b _",
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ _ _",
					),
				},
			},
		},
	)

	testMoves(
		t,
		brett(
			"_ _ _ _ ",
			" _ A a _",
			"_ b _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ _ _",
		),
		dame.SpielerUnten,
		dame.PossibleMoves{
			dame.Move{
				Start: dame.Position{Zeile: 2, Spalte: 2},
				Ende:  dame.Position{Zeile: 0, Spalte: 4},
				Steps: dame.MoveSteps{
					brett(
						"_ _ B _ ",
						" _ _ a _",
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ _ _",
					),
				},
			},
		},
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
		dame.PossibleMoves{
			// Nach unten rechts
			dame.Move{
				Start: dame.Position{Zeile: 4, Spalte: 4},
				Ende:  dame.Position{Zeile: 5, Spalte: 5},
				Steps: dame.MoveSteps{
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
				},
			},
			dame.Move{
				Start: dame.Position{Zeile: 4, Spalte: 4},
				Ende:  dame.Position{Zeile: 6, Spalte: 6},
				Steps: dame.MoveSteps{
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
				},
			},
			// Nach unten links
			dame.Move{
				Start: dame.Position{Zeile: 4, Spalte: 4},
				Ende:  dame.Position{Zeile: 5, Spalte: 3},
				Steps: dame.MoveSteps{
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
				},
			},
			dame.Move{
				Start: dame.Position{Zeile: 4, Spalte: 4},
				Ende:  dame.Position{Zeile: 6, Spalte: 2},
				Steps: dame.MoveSteps{
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
				},
			},
			// Nach oben rechts
			dame.Move{
				Start: dame.Position{Zeile: 4, Spalte: 4},
				Ende:  dame.Position{Zeile: 3, Spalte: 5},
				Steps: dame.MoveSteps{
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
				},
			},
			dame.Move{
				Start: dame.Position{Zeile: 4, Spalte: 4},
				Ende:  dame.Position{Zeile: 2, Spalte: 6},
				Steps: dame.MoveSteps{
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
				},
			},
			// Nach oben links
			dame.Move{
				Start: dame.Position{Zeile: 4, Spalte: 4},
				Ende:  dame.Position{Zeile: 3, Spalte: 3},
				Steps: dame.MoveSteps{
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
				},
			},
			dame.Move{
				Start: dame.Position{Zeile: 4, Spalte: 4},
				Ende:  dame.Position{Zeile: 2, Spalte: 2},
				Steps: dame.MoveSteps{
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
				},
			},
			dame.Move{
				Start: dame.Position{Zeile: 4, Spalte: 4},
				Ende:  dame.Position{Zeile: 1, Spalte: 1},
				Steps: dame.MoveSteps{
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
				},
			},
		},
	)
}

func TestDameSchlagen(t *testing.T) {
	testMoves(
		t,
		brett(
			"_ _ _ _ ",
			" _ _ b _",
			"_ _ _ _ ",
			" _ A _ _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ _ _",
		),
		dame.SpielerOben,
		dame.PossibleMoves{
			dame.Move{
				Start: dame.Position{Zeile: 3, Spalte: 3},
				Ende:  dame.Position{Zeile: 0, Spalte: 6},
				Steps: dame.MoveSteps{
					brett(
						"_ _ _ A ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ _ _",
					),
				},
			},
		},
	)

	testMoves(
		t,
		brett(
			"_ _ _ _ ",
			" _ b _ _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ B _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" A _ _ _",
		),
		dame.SpielerOben,
		dame.PossibleMoves{
			dame.Move{
				Start: dame.Position{Zeile: 7, Spalte: 1},
				Ende:  dame.Position{Zeile: 0, Spalte: 2},
				Steps: dame.MoveSteps{
					brett(
						"_ _ _ _ ",
						" _ b _ _",
						"_ _ _ _ ",
						" _ _ A _",
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ _ _",
					),
					brett(
						"_ A _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ _ _",
					),
				},
			},
		},
	)

	testMoves(
		t,
		brett(
			"_ _ _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" b _ a _",
			"_ _ _ _ ",
			" _ _ _ _",
			"_ b B _ ",
			" A _ _ _",
		),
		dame.SpielerOben,
		dame.PossibleMoves{
			dame.Move{
				Start: dame.Position{Zeile: 7, Spalte: 1},
				Ende:  dame.Position{Zeile: 2, Spalte: 0},
				Steps: dame.MoveSteps{
					brett(
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" b _ a _",
						"_ _ _ _ ",
						" _ A _ _",
						"_ _ B _ ",
						" _ _ _ _",
					),
					brett(
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" b _ a _",
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ A _",
					),
					brett(
						"_ _ _ _ ",
						" _ _ _ _",
						"A _ _ _ ",
						" _ _ a _",
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ _ _",
					),
				},
			},
			dame.Move{
				Start: dame.Position{Zeile: 7, Spalte: 1},
				Ende:  dame.Position{Zeile: 7, Spalte: 5},
				Steps: dame.MoveSteps{
					brett(
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" b _ a _",
						"_ _ _ _ ",
						" _ A _ _",
						"_ _ B _ ",
						" _ _ _ _",
					),
					brett(
						"_ _ _ _ ",
						" _ _ _ _",
						"A _ _ _ ",
						" _ _ a _",
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ B _ ",
						" _ _ _ _",
					),
					brett(
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ a _",
						"_ _ _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ A _",
					),
				},
			},
		},
	)
}

func TestKeineZugMöglich(t *testing.T) {
	testMoves(
		t,
		brett(
			"_ _ _ _ ",
			" _ B _ _",
			"_ _ b _ ",
			" _ _ _ _",
			"_ b _ _ ",
			" _ _ _ _",
			"_ _ _ _ ",
			" _ _ _ _",
		),
		dame.SpielerOben,
		dame.PossibleMoves{
			dame.Move{
				Steps: dame.MoveSteps{
					brett(
						"_ _ _ _ ",
						" _ B _ _",
						"_ _ b _ ",
						" _ _ _ _",
						"_ b _ _ ",
						" _ _ _ _",
						"_ _ _ _ ",
						" _ _ _ _",
					),
				},
			},
		},
	)
}

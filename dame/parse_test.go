package dame_test

import (
	"testing"

	"github.com/Lama06/Dame-Go.git/dame"
)

func TestParseBrett(t *testing.T) {
	brett, err := dame.ParseBrett(
		"a a a a ",
		" a a a a",
		"a a a a ",
		" _ _ _ _",
		"_ _ _ _ ",
		" b b b b",
		"b b b b ",
		" b b b b",
	)
	if err != nil {
		t.Fatalf("failed to parse brett: %v", err)
	}

	expected := dame.Brett{
		0: dame.Zeile{
			dame.SteinSpielerOben, dame.SteinSpielerOben, dame.SteinSpielerOben, dame.SteinSpielerOben,
		},
		1: dame.Zeile{
			dame.SteinSpielerOben, dame.SteinSpielerOben, dame.SteinSpielerOben, dame.SteinSpielerOben,
		},
		2: dame.Zeile{
			dame.SteinSpielerOben, dame.SteinSpielerOben, dame.SteinSpielerOben, dame.SteinSpielerOben,
		},
		dame.BrettSize - 3: dame.Zeile{
			dame.SteinSpielerUnten, dame.SteinSpielerUnten, dame.SteinSpielerUnten, dame.SteinSpielerUnten,
		},
		dame.BrettSize - 2: dame.Zeile{
			dame.SteinSpielerUnten, dame.SteinSpielerUnten, dame.SteinSpielerUnten, dame.SteinSpielerUnten,
		},
		dame.BrettSize - 1: dame.Zeile{
			dame.SteinSpielerUnten, dame.SteinSpielerUnten, dame.SteinSpielerUnten, dame.SteinSpielerUnten,
		},
	}

	if brett != expected {
		t.FailNow()
	}
}

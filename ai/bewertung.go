package ai

import (
	"math"

	"github.com/Lama06/Dame-Go.git/dame"
)

func countFelder(brett dame.Brett, feld dame.Feld) int {
	var count int
	for zeile := 0; zeile < dame.BrettSize; zeile++ {
		for spalte := 0; spalte < dame.BrettSize; spalte++ {
			position := dame.Position{Zeile: zeile, Spalte: spalte}
			if !position.Valid() {
				continue
			}
			if brett.Get(position) == feld {
				count++
			}
		}
	}
	return count
}

type spielerStatistik struct {
	steine, damen int
}

func (s spielerStatistik) bewertung() float32 {
	if s.steine == 0 && s.damen == 0 {
		return float32(math.Inf(-1))
	}
	return float32(s.steine) + float32(s.damen)*3
}

type brettStatistik struct {
	spielerOben, spielerUnten spielerStatistik
}

func statistikFromBrett(brett dame.Brett) brettStatistik {
	return brettStatistik{
		spielerOben: spielerStatistik{
			steine: countFelder(brett, dame.Stein(dame.SpielerOben)),
			damen:  countFelder(brett, dame.Dame(dame.SpielerOben)),
		},
		spielerUnten: spielerStatistik{
			steine: countFelder(brett, dame.Stein(dame.SpielerUnten)),
			damen:  countFelder(brett, dame.Dame(dame.SpielerUnten)),
		},
	}
}

func (s brettStatistik) bewertung(perspektive dame.Spieler) float32 {
	switch perspektive {
	case dame.SpielerOben:
		return s.spielerOben.bewertung() - s.spielerUnten.bewertung()
	case dame.SpielerUnten:
		return s.spielerUnten.bewertung() - s.spielerOben.bewertung()
	default:
		panic("unreachable")
	}
}

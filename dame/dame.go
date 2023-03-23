package dame

import (
	"fmt"
	"strings"
)

const BrettSize = 8

type Position struct {
	Spalte, Zeile int
}

func (p Position) Valid() bool {
	return p.Spalte >= 0 && p.Zeile >= 0 &&
		p.Spalte < BrettSize && p.Zeile < BrettSize &&
		p.Spalte%2 == p.Zeile%2
}

type RichtungHorizontal bool

const (
	Links  RichtungHorizontal = false
	Rechts RichtungHorizontal = true
)

func (r RichtungHorizontal) Offset() int {
	switch r {
	case Links:
		return -1
	case Rechts:
		return 1
	default:
		panic("unreachable")
	}
}

type RichtungVertikal bool

const (
	Oben  RichtungVertikal = false
	Unten RichtungVertikal = true
)

func (r RichtungVertikal) Offset() int {
	switch r {
	case Oben:
		return -1
	case Unten:
		return 1
	default:
		panic("unreachable")
	}
}

type Spieler bool

const (
	SpielerOben  Spieler = false
	SpielerUnten Spieler = true
)

func (s Spieler) MoveDirection() RichtungVertikal {
	switch s {
	case SpielerOben:
		return Unten
	case SpielerUnten:
		return Oben
	default:
		panic("unreachable")
	}
}

func (s Spieler) DameZeile() int {
	switch s {
	case SpielerOben:
		return BrettSize - 1
	case SpielerUnten:
		return 0
	default:
		panic("unreachable")
	}
}

func (s Spieler) String() string {
	switch s {
	case SpielerOben:
		return "Spieler Oben (Blau)"
	case SpielerUnten:
		return "Spieler Unten (Rot)"
	default:
		panic("unreachable")
	}
}

type Feld byte

const (
	Leer              Feld = iota
	SteinSpielerOben       = iota
	SteinSpielerUnten      = iota
	DameSpielerOben        = iota
	DameSpielerUnten       = iota
)

func Stein(spieler Spieler) Feld {
	switch spieler {
	case SpielerOben:
		return SteinSpielerOben
	case SpielerUnten:
		return SteinSpielerUnten
	default:
		panic("unreachable")
	}
}

func Dame(spieler Spieler) Feld {
	switch spieler {
	case SpielerOben:
		return DameSpielerOben
	case SpielerUnten:
		return DameSpielerUnten
	default:
		panic("unreachable")
	}
}

func (f Feld) IsStein() (Spieler, bool) {
	switch f {
	case SteinSpielerOben:
		return SpielerOben, true
	case SteinSpielerUnten:
		return SpielerUnten, true
	default:
		return false, false
	}
}

func (f Feld) IsDame() (Spieler, bool) {
	switch f {
	case DameSpielerOben:
		return SpielerOben, true
	case DameSpielerUnten:
		return SpielerUnten, true
	default:
		return false, false
	}
}

func (f Feld) Rune() rune {
	switch f {
	case Leer:
		return '_'
	case SteinSpielerOben:
		return 'a'
	case DameSpielerOben:
		return 'A'
	case SteinSpielerUnten:
		return 'b'
	case DameSpielerUnten:
		return 'B'
	default:
		panic("unreachable")
	}
}

type Zeile [BrettSize / 2]Feld

func (z Zeile) Get(position Position) Feld {
	if !position.Valid() {
		panic("invalid position")
	}

	return z[position.Spalte/2]
}

func (z *Zeile) Set(position Position, feld Feld) {
	if !position.Valid() {
		panic("invalid position")
	}

	z[position.Spalte/2] = feld
}

func (z Zeile) String(zeile int) string {
	var result strings.Builder
	for spalte := 0; spalte < BrettSize; spalte++ {
		position := Position{spalte, zeile}
		if position.Valid() {
			result.WriteRune(rune(z.Get(position).Rune()))
		} else {
			result.WriteRune(' ')
		}
	}
	return result.String()
}

type Brett [BrettSize]Zeile

var DefaultBrett Brett

func init() {
	defaultBrett, err := ParseBrett(
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
		panic(fmt.Errorf("failed to parse default brett: %w", err))
	}
	DefaultBrett = defaultBrett
}

func (b Brett) Get(position Position) Feld {
	if !position.Valid() {
		panic("invalid position")
	}

	return b[position.Zeile].Get(position)
}

func (b *Brett) Set(position Position, feld Feld) {
	if !position.Valid() {
		panic("invalid position")
	}

	b[position.Zeile].Set(position, feld)
}

func (b Brett) String() string {
	var result strings.Builder
	for zeile := 0; zeile < BrettSize; zeile++ {
		result.WriteString(b[zeile].String(zeile))
		if zeile != BrettSize-1 {
			result.WriteRune('\n')
		}
	}
	return result.String()
}

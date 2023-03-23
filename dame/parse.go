package dame

import (
	"errors"
	"fmt"
)

func ParseFeld(char rune) (Feld, error) {
	switch char {
	case '_':
		return Leer, nil
	case 'a':
		return SteinSpielerOben, nil
	case 'A':
		return DameSpielerOben, nil
	case 'b':
		return SteinSpielerUnten, nil
	case 'B':
		return DameSpielerUnten, nil
	default:
		return 0, fmt.Errorf("invalid character: %c", char)
	}
}

func ParseZeile(zeile int, text string) (Zeile, error) {
	if len(text) != BrettSize {
		return Zeile{}, fmt.Errorf("invalid length: %v", len(text))
	}

	var result Zeile
	for spalte, feldText := range text {
		position := Position{spalte, zeile}

		if !position.Valid() {
			if feldText != ' ' {
				return Zeile{}, fmt.Errorf("expected whitespace at spalte %v", spalte)
			}
			continue
		}

		feld, err := ParseFeld(feldText)
		if err != nil {
			return Zeile{}, fmt.Errorf("failed to parse feld: %w", err)
		}

		result.Set(position, feld)
	}
	return result, nil
}

func ParseBrett(zeilen ...string) (Brett, error) {
	if len(zeilen) != BrettSize {
		return Brett{}, errors.New("invalid amount of rows")
	}

	var result Brett

	for zeileIndex, zeileText := range zeilen {
		zeile, err := ParseZeile(zeileIndex, zeileText)
		if err != nil {
			return Brett{}, fmt.Errorf("failed to parse zeile %v: %w", zeileIndex, err)
		}
		result[zeileIndex] = zeile
	}

	return result, nil
}

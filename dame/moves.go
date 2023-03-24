package dame

type MoveSteps []Brett

func (first MoveSteps) Equals(second MoveSteps) bool {
	if len(first) != len(second) {
		return false
	}

	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func (m MoveSteps) Result() Brett {
	return m[len(m)-1]
}

type Move struct {
	Start Position
	Ende  Position
	Steps MoveSteps
}

func (first Move) Equals(second Move) bool {
	if !first.Steps.Equals(second.Steps) {
		return false
	}

	if first.Start != second.Start {
		return false
	}

	if first.Ende != second.Ende {
		return false
	}

	return true
}

type PossibleMoves []Move

func (p PossibleMoves) Contains(move Move) bool {
	for i := range p {
		if p[i].Equals(move) {
			return true
		}
	}

	return false
}

func (first PossibleMoves) Equals(second PossibleMoves) bool {
	if len(first) != len(second) {
		return false
	}

	for _, firstMove := range first {
		if !second.Contains(firstMove) {
			return false
		}
	}

	return true
}

func (b Brett) getSteinSchlagenMoves(position Position, backwards bool) PossibleMoves {
	if !position.Valid() {
		return nil
	}

	if !b.Get(position).IsStein() {
		return nil
	}
	spieler, _ := b.Get(position).Spieler()

	var result PossibleMoves

	var richtungenVertikal []RichtungVertikal
	if backwards {
		richtungenVertikal = []RichtungVertikal{Oben, Unten}
	} else {
		richtungenVertikal = []RichtungVertikal{spieler.MoveDirection()}
	}

	for _, richtungVertikal := range richtungenVertikal {
		for _, richtungHorizontal := range [2]RichtungHorizontal{Links, Rechts} {
			schlagenPosition := Position{
				Zeile:  position.Zeile + richtungVertikal.Offset(),
				Spalte: position.Spalte + richtungHorizontal.Offset(),
			}
			if !schlagenPosition.Valid() {
				continue
			}
			schlagenFeld := b.Get(schlagenPosition)
			if schlagenFeld == Stein(spieler) || schlagenFeld == Dame(spieler) || schlagenFeld == Leer {
				continue
			}

			neuePosition := Position{
				Zeile:  position.Zeile + richtungVertikal.Offset()*2,
				Spalte: position.Spalte + richtungHorizontal.Offset()*2,
			}
			if !neuePosition.Valid() {
				continue
			}
			if b.Get(neuePosition) != Leer {
				continue
			}

			neuePositionFeld := Stein(spieler)
			if neuePosition.Zeile == spieler.DameZeile() {
				neuePositionFeld = Dame(spieler)
			}

			neuesBrett := b
			neuesBrett.Set(position, Leer)
			neuesBrett.Set(schlagenPosition, Leer)
			neuesBrett.Set(neuePosition, neuePositionFeld)

			followingSchlagenMoves := neuesBrett.getSteinSchlagenMoves(neuePosition, true)
			if len(followingSchlagenMoves) == 0 {
				result = append(result, Move{
					Start: position,
					Ende:  neuePosition,
					Steps: MoveSteps{neuesBrett},
				})
			} else {
				for _, followingSchlagenMove := range followingSchlagenMoves {
					result = append(result, Move{
						Start: position,
						Ende:  followingSchlagenMove.Ende,
						Steps: append(MoveSteps{neuesBrett}, followingSchlagenMove.Steps...),
					})
				}
			}
		}
	}

	return result
}

func (b Brett) getSteinBewegenMoves(position Position) PossibleMoves {
	if !position.Valid() {
		return nil
	}

	if !b.Get(position).IsStein() {
		return nil
	}
	spieler, _ := b.Get(position).Spieler()

	var result PossibleMoves

	for _, richtungHorizontal := range [2]RichtungHorizontal{Links, Rechts} {
		neuePosition := Position{
			Spalte: position.Spalte + richtungHorizontal.Offset(),
			Zeile:  position.Zeile + spieler.MoveDirection().Offset(),
		}
		if !neuePosition.Valid() {
			continue
		}
		if b.Get(neuePosition) != Leer {
			continue
		}

		neuePositionFeld := Stein(spieler)
		if neuePosition.Zeile == spieler.DameZeile() {
			neuePositionFeld = Dame(spieler)
		}

		neuesBrett := b
		neuesBrett.Set(position, Leer)
		neuesBrett.Set(neuePosition, neuePositionFeld)
		result = append(result, Move{
			Start: position,
			Ende:  neuePosition,
			Steps: MoveSteps{neuesBrett},
		})
	}

	return result
}

func (b Brett) getDameBewegenMoves(position Position) PossibleMoves {
	if !position.Valid() {
		return nil
	}

	if !b.Get(position).IsDame() {
		return nil
	}
	spieler, _ := b.Get(position).Spieler()

	var result PossibleMoves

	for _, richtungVertikal := range [2]RichtungVertikal{Oben, Unten} {
	richtungHorizontal:
		for _, richtungHorizontal := range [2]RichtungHorizontal{Links, Rechts} {
			for numberOfFields := 1; numberOfFields < BrettSize; numberOfFields++ {
				neuePosition := Position{
					Zeile:  position.Zeile + richtungVertikal.Offset()*numberOfFields,
					Spalte: position.Spalte + richtungHorizontal.Offset()*numberOfFields,
				}
				if !neuePosition.Valid() {
					continue
				}
				if b.Get(neuePosition) != Leer {
					continue richtungHorizontal
				}

				neuesBrett := b
				neuesBrett.Set(position, Leer)
				neuesBrett.Set(neuePosition, Dame(spieler))
				result = append(result, Move{
					Start: position,
					Ende:  neuePosition,
					Steps: MoveSteps{neuesBrett},
				})
			}
		}
	}

	return result
}

func (b Brett) getDameSchlagenMoves(position Position) PossibleMoves {
	if !position.Valid() {
		return nil
	}

	if !b.Get(position).IsDame() {
		return nil
	}
	spieler, _ := b.Get(position).Spieler()

	var result PossibleMoves

	for _, richtungVertikal := range [2]RichtungVertikal{Oben, Unten} {
	richtungHorizontal:
		for _, richtungHorizontal := range [2]RichtungHorizontal{Links, Rechts} {
			for numberOfFields := 1; numberOfFields < BrettSize; numberOfFields++ {
				schlagenPosition := Position{
					Zeile:  position.Zeile + richtungVertikal.Offset()*numberOfFields,
					Spalte: position.Spalte + richtungHorizontal.Offset()*numberOfFields,
				}
				if !schlagenPosition.Valid() {
					continue
				}
				schlagenFeld := b.Get(schlagenPosition)
				if schlagenFeld == Leer {
					continue
				}
				if schlagenFeld == Stein(spieler) || schlagenFeld == Dame(spieler) {
					continue richtungHorizontal
				}

				neuePosition := Position{
					Zeile:  schlagenPosition.Zeile + richtungVertikal.Offset(),
					Spalte: schlagenPosition.Spalte + richtungHorizontal.Offset(),
				}
				if !neuePosition.Valid() {
					continue
				}
				if b.Get(neuePosition) != Leer {
					continue richtungHorizontal
				}

				neuesBrett := b
				neuesBrett.Set(position, Leer)
				neuesBrett.Set(schlagenPosition, Leer)
				neuesBrett.Set(neuePosition, Dame(spieler))

				followingSchlagenMoves := neuesBrett.getDameSchlagenMoves(neuePosition)
				if len(followingSchlagenMoves) == 0 {
					result = append(result, Move{
						Start: position,
						Ende:  neuePosition,
						Steps: MoveSteps{neuesBrett},
					})
				} else {
					for _, followingSchlagenMove := range followingSchlagenMoves {
						result = append(result, Move{
							Start: position,
							Ende:  followingSchlagenMove.Ende,
							Steps: append(MoveSteps{neuesBrett}, followingSchlagenMove.Steps...),
						})
					}
				}
			}
		}
	}

	return result
}

func (b Brett) PossibleMovesForSpieler(spieler Spieler) PossibleMoves {
	var result PossibleMoves

	for zeile := 0; zeile < BrettSize; zeile++ {
		for spalte := 0; spalte < BrettSize; spalte++ {
			position := Position{Zeile: zeile, Spalte: spalte}
			if !position.Valid() {
				continue
			}

			if feldSpieler, ok := b.Get(position).Spieler(); !ok || feldSpieler != spieler {
				continue
			}

			result = append(result, b.getSteinSchlagenMoves(position, false)...)
			result = append(result, b.getDameSchlagenMoves(position)...)
		}
	}

	if len(result) > 0 {
		return result
	}

	for zeile := 0; zeile < BrettSize; zeile++ {
		for spalte := 0; spalte < BrettSize; spalte++ {
			position := Position{Zeile: zeile, Spalte: spalte}
			if !position.Valid() {
				continue
			}

			if feldSpieler, ok := b.Get(position).Spieler(); !ok || feldSpieler != spieler {
				continue
			}

			result = append(result, b.getSteinBewegenMoves(position)...)
			result = append(result, b.getDameBewegenMoves(position)...)
		}
	}

	if len(result) > 0 {
		return result
	}

	return PossibleMoves{
		Move{
			Steps: MoveSteps{b},
		},
	}
}

func (b Brett) PossibleMovesForPosition(position Position) PossibleMoves {
	spieler, ok := b.Get(position).Spieler()
	if !ok {
		return nil
	}

	movesForSpieler := b.PossibleMovesForSpieler(spieler)

	var result PossibleMoves
	for _, moveForSpieler := range movesForSpieler {
		if moveForSpieler.Start == position {
			result = append(result, moveForSpieler)
		}
	}
	return result
}

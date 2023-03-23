package dame

func (b Brett) addSteinSchlagenMoves(moves map[Brett]struct{}, position Position, backwards bool) {
	if !position.Valid() {
		return
	}

	spieler, stein := b.Get(position).IsStein()
	if !stein {
		return
	}

	var richtungenVertikal []RichtungVertikal
	if backwards {
		richtungenVertikal = []RichtungVertikal{Oben, Unten}
	} else {
		richtungenVertikal = []RichtungVertikal{spieler.MoveDirection()}
	}

	for _, richtungVertikal := range richtungenVertikal {
		for _, richtungHorizontal := range []RichtungHorizontal{Links, Rechts} {
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

			lenMoves := len(moves)
			neuesBrett.addSteinSchlagenMoves(moves, neuePosition, true)
			if lenMoves == len(moves) {
				moves[neuesBrett] = struct{}{}
			}
		}
	}
}

func (b Brett) addAlleSteinSchlagenMoves(moves map[Brett]struct{}, spieler Spieler) {
	for zeile := 0; zeile < BrettSize; zeile++ {
		for spalte := 0; spalte < BrettSize; spalte++ {
			position := Position{
				Spalte: spalte,
				Zeile:  zeile,
			}
			if !position.Valid() {
				continue
			}
			if b.Get(position) != Stein(spieler) {
				continue
			}
			b.addSteinSchlagenMoves(moves, position, false)
		}
	}
}

func (b Brett) addAlleSteinBewegenMoves(moves map[Brett]struct{}, spieler Spieler) {
	for zeile := 0; zeile < BrettSize; zeile++ {
		for spalte := 0; spalte < BrettSize; spalte++ {
			position := Position{
				Spalte: spalte,
				Zeile:  zeile,
			}
			if !position.Valid() {
				continue
			}
			if b.Get(position) != Stein(spieler) {
				continue
			}

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
				moves[neuesBrett] = struct{}{}
			}
		}
	}
}

func (b Brett) addAlleDameBewegenMoves(moves map[Brett]struct{}, spieler Spieler) {
	for zeile := 0; zeile < BrettSize; zeile++ {
		for spalte := 0; spalte < BrettSize; spalte++ {
			position := Position{
				Spalte: spalte,
				Zeile:  zeile,
			}
			if !position.Valid() {
				continue
			}
			if b.Get(position) != Dame(spieler) {
				continue
			}

			for _, richtungVertikal := range [2]RichtungVertikal{Oben, Unten} {
			richtungHorizontal:
				for _, richtungHorizontal := range [2]RichtungHorizontal{Links, Rechts} {
					for numberOfFields := 1; numberOfFields < BrettSize; numberOfFields++ {
						neuesFeld := Position{
							Zeile:  position.Zeile + richtungVertikal.Offset()*numberOfFields,
							Spalte: position.Spalte + richtungHorizontal.Offset()*numberOfFields,
						}
						if !neuesFeld.Valid() {
							continue
						}
						if b.Get(neuesFeld) != Leer {
							continue richtungHorizontal
						}

						neuesBrett := b
						neuesBrett.Set(position, Leer)
						neuesBrett.Set(neuesFeld, Dame(spieler))
						moves[neuesBrett] = struct{}{}
					}
				}
			}
		}
	}
}

func (b Brett) addDameSchlagenMoves(moves map[Brett]struct{}, position Position) {
	if !position.Valid() {
		return
	}
	spieler, isDame := b.Get(position).IsDame()
	if !isDame {
		return
	}

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

				lenBefore := len(moves)
				neuesBrett.addDameSchlagenMoves(moves, neuePosition)
				if lenBefore == len(moves) {
					moves[neuesBrett] = struct{}{}
				}
			}
		}
	}
}

func (b Brett) addAlleDameSchlagenMoves(moves map[Brett]struct{}, spieler Spieler) {
	for zeile := 0; zeile < BrettSize; zeile++ {
		for spalte := 0; spalte < BrettSize; spalte++ {
			position := Position{
				Spalte: spalte,
				Zeile:  zeile,
			}
			if !position.Valid() {
				continue
			}
			if b.Get(position) != Dame(spieler) {
				continue
			}
			b.addDameSchlagenMoves(moves, position)
		}
	}
}

func (b Brett) PossibleMoves(spieler Spieler) map[Brett]struct{} {
	result := make(map[Brett]struct{})
	b.addAlleSteinSchlagenMoves(result, spieler)
	b.addAlleDameSchlagenMoves(result, spieler)
	if len(result) > 0 {
		return result
	}
	b.addAlleSteinBewegenMoves(result, spieler)
	b.addAlleDameBewegenMoves(result, spieler)
	if len(result) == 0 {
		result[b] = struct{}{}
	}
	return result
}

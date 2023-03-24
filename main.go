package main

import (
	"image/color"

	"github.com/Lama06/Dame-Go.git/ai"
	"github.com/Lama06/Dame-Go.git/dame"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

const (
	FeldSize          = 50
	WindowSize        = dame.BrettSize * FeldSize
	ComputerSpieler   = dame.SpielerOben
	MenschSpieler     = dame.SpielerUnten
	AiDepth           = 7
	ComputerStepDelay = 50
)

var (
	SelectedFieldColor     = colornames.Purple
	PossibleMoveFieldColor = colornames.Pink
)

func feldColor(feld dame.Feld) color.RGBA {
	switch feld {
	case dame.Leer:
		return colornames.Black
	case dame.SteinSpielerOben:
		return colornames.Lightblue
	case dame.DameSpielerOben:
		return colornames.Darkblue
	case dame.SteinSpielerUnten:
		return colornames.Orangered
	case dame.DameSpielerUnten:
		return colornames.Darkred
	default:
		panic("unreachable")
	}
}

func drawRect(pixels []byte, xStart, yStart, width, height int, color color.RGBA) {
	for x := xStart; x < xStart+width; x++ {
		for y := yStart; y < yStart+height; y++ {
			pixelIndex := (y*WindowSize + x) * 4
			pixels[pixelIndex] = color.R
			pixels[pixelIndex+1] = color.G
			pixels[pixelIndex+2] = color.B
			pixels[pixelIndex+3] = color.A
		}
	}
}

type Game struct {
	pixels []byte

	brett dame.Brett

	hasSelectedPosition bool
	selectedPosition    dame.Position

	remainingComputerSteps dame.MoveSteps
	nextComputerStepTimer  int
}

func (g *Game) isPossibleMoveEndePosition(position dame.Position) (dame.Move, bool) {
	if !g.hasSelectedPosition {
		return dame.Move{}, false
	}

	spieler, ok := g.brett.Get(g.selectedPosition).Spieler()
	if !ok || spieler != MenschSpieler {
		return dame.Move{}, false
	}

	possibleMoves := g.brett.PossibleMovesForPosition(g.selectedPosition)
	for _, possibleMove := range possibleMoves {
		if possibleMove.Ende == position {
			return possibleMove, true
		}
	}

	return dame.Move{}, false
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.ReadPixels(g.pixels)

	for zeile := 0; zeile < dame.BrettSize; zeile++ {
		for spalte := 0; spalte < dame.BrettSize; spalte++ {
			position := dame.Position{Spalte: spalte, Zeile: zeile}
			if !position.Valid() {
				drawRect(g.pixels, spalte*FeldSize, zeile*FeldSize, FeldSize, FeldSize, colornames.White)
				continue
			}

			if g.hasSelectedPosition && g.selectedPosition == position {
				drawRect(g.pixels, spalte*FeldSize, zeile*FeldSize, FeldSize, FeldSize, SelectedFieldColor)
			} else if _, ok := g.isPossibleMoveEndePosition(position); ok {
				drawRect(g.pixels, spalte*FeldSize, zeile*FeldSize, FeldSize, FeldSize, PossibleMoveFieldColor)
			} else {
				drawRect(g.pixels, spalte*FeldSize, zeile*FeldSize, FeldSize, FeldSize, colornames.Black)
			}

			if g.brett.Get(position) != dame.Leer {
				drawRect(g.pixels, spalte*FeldSize+10, zeile*FeldSize+10, FeldSize-20, FeldSize-20, feldColor(g.brett.Get(position)))
			}
		}
	}

	screen.WritePixels(g.pixels)
}

func (g *Game) Update() error {
	if len(g.remainingComputerSteps) != 0 {
		if g.nextComputerStepTimer > 0 {
			g.nextComputerStepTimer--
		}

		if g.nextComputerStepTimer == 0 {
			g.brett = g.remainingComputerSteps[0]
			g.remainingComputerSteps = g.remainingComputerSteps[1:len(g.remainingComputerSteps)]
			g.nextComputerStepTimer = ComputerStepDelay
		}

		return nil
	}

	if !inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		return nil
	}

	mouseX, mouseY := ebiten.CursorPosition()
	spalte := mouseX / FeldSize
	zeile := mouseY / FeldSize
	position := dame.Position{Spalte: spalte, Zeile: zeile}
	if !position.Valid() {
		g.hasSelectedPosition = false
		return nil
	}

	if move, ok := g.isPossibleMoveEndePosition(position); ok {
		g.brett = move.Steps.Result()
		g.remainingComputerSteps = ai.FindBestMove(g.brett, ComputerSpieler, AiDepth).Steps
		g.nextComputerStepTimer = 0
		g.hasSelectedPosition = false
		return nil
	}

	spieler, ok := g.brett.Get(position).Spieler()
	if ok && spieler == MenschSpieler {
		g.hasSelectedPosition = true
		g.selectedPosition = position
		return nil
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (windowWidth, windowHeight int) {
	return WindowSize, WindowSize
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.RunGame(&Game{
		brett:               dame.DefaultBrett,
		pixels:              make([]byte, WindowSize*WindowSize*4),
		hasSelectedPosition: false,
	})
}

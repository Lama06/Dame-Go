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
	ComputerStepDelay = 30
)

var (
	SelectedFieldColor     = color.RGBA{84, 6, 66, 0}
	PossibleMoveFieldColor = color.RGBA{191, 25, 119, 0}
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

	calculatingComputerMove bool
	computerMove            chan dame.Move

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

func (g *Game) handleClick(x, y int) {
	spalte := x / FeldSize
	zeile := y / FeldSize
	position := dame.Position{Spalte: spalte, Zeile: zeile}
	if !position.Valid() {
		g.hasSelectedPosition = false
		return
	}

	if move, ok := g.isPossibleMoveEndePosition(position); ok {
		g.brett = move.Steps.Result()
		go func(brett dame.Brett) {
			g.computerMove <- ai.FindBestMove(brett, ComputerSpieler, AiDepth)
		}(g.brett)
		g.hasSelectedPosition = false
		g.calculatingComputerMove = true
		ebiten.SetWindowTitle("Dame - Überlegen...")
		return
	}

	spieler, ok := g.brett.Get(position).Spieler()
	if ok && spieler == MenschSpieler {
		g.hasSelectedPosition = true
		g.selectedPosition = position
	}
}

func (g *Game) Update() error {
	if g.calculatingComputerMove {
		select {
		case move := <-g.computerMove:
			g.calculatingComputerMove = false
			g.remainingComputerSteps = move.Steps
			g.nextComputerStepTimer = 10
			ebiten.SetWindowTitle("Dame")
		default:
			return nil
		}
	}

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

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		g.handleClick(ebiten.CursorPosition())
		return nil
	}

	touches := inpututil.AppendJustPressedTouchIDs(nil)
	if len(touches) == 1 {
		g.handleClick(ebiten.TouchPosition(touches[0]))
		return nil
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (windowWidth, windowHeight int) {
	return WindowSize, WindowSize
}

func main() {
	ebiten.SetWindowTitle("Dame")
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.RunGame(&Game{
		brett:               dame.DefaultBrett,
		pixels:              make([]byte, WindowSize*WindowSize*4),
		hasSelectedPosition: false,
		computerMove:        make(chan dame.Move),
	})
}

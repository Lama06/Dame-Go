package main

import (
	"image/color"
	"log"

	"github.com/Lama06/Dame-Go.git/ai"
	"github.com/Lama06/Dame-Go.git/dame"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

const (
	FeldSize   = 50
	WindowSize = dame.BrettSize * FeldSize
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
	brett  dame.Brett
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

			drawRect(g.pixels, spalte*FeldSize, zeile*FeldSize, FeldSize, FeldSize, colornames.Black)
			drawRect(g.pixels, spalte*FeldSize+10, zeile*FeldSize+10, FeldSize-20, FeldSize-20, feldColor(g.brett.Get(position)))
		}
	}

	screen.WritePixels(g.pixels)
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		log.Println("Finde besten Zug...")
		g.brett = ai.FindBestMove(g.brett, dame.SpielerOben, 8)
		log.Println("Fertig")
		return nil
	}

	mouseX, mouseY := ebiten.CursorPosition()
	spalte := mouseX / FeldSize
	zeile := mouseY / FeldSize
	position := dame.Position{Spalte: spalte, Zeile: zeile}
	if !position.Valid() {
		return nil
	}

	var neuesFeld dame.Feld = g.brett.Get(position)
	if inpututil.IsKeyJustReleased(ebiten.KeyM) {
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			neuesFeld = dame.DameSpielerUnten
		} else {
			neuesFeld = dame.SteinSpielerUnten
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyC) {
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			neuesFeld = dame.DameSpielerOben
		} else {
			neuesFeld = dame.SteinSpielerOben
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyBackspace) {
		neuesFeld = dame.Leer
	}

	g.brett.Set(position, neuesFeld)

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (windowWidth, windowHeight int) {
	return WindowSize, WindowSize
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.RunGame(&Game{
		brett:  dame.DefaultBrett,
		pixels: make([]byte, WindowSize*WindowSize*4),
	})
}

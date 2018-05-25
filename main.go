package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

var s *sprite

func init() {
	s = newSprite()
}

func update(screen *ebiten.Image) error {

	// Fill the screen with #FF0000 color
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
	s.handleInput()
	s.update()
	s.draw(screen)

	//ebitenutil.DebugPrint(screen, "Hello world!")
	return nil
}

func main() {
	ebiten.Run(update, 320, 240, 2, "Hello world!")
}

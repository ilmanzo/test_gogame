package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

var s *sprite
var o *obstacle

func init() {
	s = newSprite()
	o = newObstacle()
}

func update(screen *ebiten.Image) error {

	// Fill the screen with #800000 color
	screen.Fill(color.NRGBA{0x80, 0x00, 0x00, 0xff})
	o.update()
	s.handleInput()
	s.update()
	s.draw(screen)
	o.draw(screen)

	//ebitenutil.DebugPrint(screen, "Hello world!")
	return nil
}

func main() {
	ebiten.Run(update, 320, 240, 2, "Hello world!")
}

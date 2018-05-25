package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type sprite struct {
	x, y, xv, yv float64
	img          *ebiten.Image
}

func (s *sprite) update() {
	s.x += s.xv
	s.y += s.yv
	if s.x < 0 {
		s.x = 0
	}
	if s.y < 0 {
		s.y = 0
	}
	s.xv *= 0.95
	s.yv *= 0.95
}

func newSprite() *sprite {
	square, _ := ebiten.NewImage(16, 16, ebiten.FilterDefault)
	square.Fill(color.NRGBA{0x00, 0xff, 0x00, 0xff})
	return &sprite{x: 0, y: 0, xv: 0, yv: 0, img: square}
}

func (s *sprite) handleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		s.yv -= 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		s.yv += 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.xv -= 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.xv += 0.1
	}
}

func (s *sprite) draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(s.x, s.y)
	screen.DrawImage(s.img, opts)
}

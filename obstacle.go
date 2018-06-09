package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

type obstacle struct {
	x, y, xv, yv float64
	img          *ebiten.Image
}

func (o *obstacle) update() {
	o.y += o.yv
	if o.y > 240 {
		o.y = 0
		o.yv = 1
		o.x = rand.Float64() * 200
	}
	o.yv *= 1.05
}

func newObstacle() *obstacle {
	square, _ := ebiten.NewImage(16, 16, ebiten.FilterDefault)
	square.Fill(color.NRGBA{0x00, 0x00, 0xff, 0xff})
	return &obstacle{x: rand.Float64() * 200, y: 0, xv: 0, yv: 1, img: square}
}

func (o *obstacle) draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(o.x, o.y)
	screen.DrawImage(o.img, opts)
}

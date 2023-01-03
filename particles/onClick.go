package particles

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"math"
)

func (s *System) onClick() {
	if inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) > 0 {
		for e := s.Content.Front(); e != nil; e = e.Next() {
			p := e.Value.(*Particle)
			if int(math.Sqrt(float64((p.PositionY - float64(s.CursorY))*float64((p.PositionY - float64(s.CursorY))) + float64((p.PositionX - float64(s.CursorX)))*float64((p.PositionX - float64(s.CursorX)))))) < 100 {
				p.ColorBlue = 0
				p.ColorGreen = 0
                p.ColorRed = 1
			} else {
				p.ColorBlue = 1
				p.ColorGreen = 1
                p.ColorRed = 1
			}
		}
	}

	if inpututil.MouseButtonPressDuration(ebiten.MouseButtonRight) > 0 {
	}
}
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
			// Si particule est à moins de 100px du pointeur
			if int(math.Sqrt(float64((p.PositionY - float64(s.CursorY))*float64((p.PositionY - float64(s.CursorY))) + float64((p.PositionX - float64(s.CursorX)))*float64((p.PositionX - float64(s.CursorX)))))) < 300 {
				var vectorX float64 = float64(s.CursorX) - p.PositionX
				var vectorY float64 = float64(s.CursorY) - p.PositionY
				p.SpeedX = vectorX * 0.05
				p.SpeedY = vectorY * 0.05	
			} 
			if int(math.Sqrt(float64((p.PositionY - float64(s.CursorY))*float64((p.PositionY - float64(s.CursorY))) + float64((p.PositionX - float64(s.CursorX)))*float64((p.PositionX - float64(s.CursorX)))))) < 1 {
				p.respawn()
			}
		}
	}

	if inpututil.MouseButtonPressDuration(ebiten.MouseButtonRight) > 0 {
	}
}
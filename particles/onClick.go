package particles

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (s *System) onClick() {
	if inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) > 0 {
		s.ParticleSpawnX = s.CursorX
		s.ParticleSpawnY = s.CursorY
		s.CreateParticle()
	}

	if inpututil.MouseButtonPressDuration(ebiten.MouseButtonRight) > 0 {

	}
}
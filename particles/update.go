package particles

import (
	"project-particles/config"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	s.Tick++
	if config.General.SpawnRate < 1 && config.General.SpawnRate > 0 {
		if s.Tick % int(1 / config.General.SpawnRate) == 0 {
			s.CreateParticle()
		}
	} else {
		for i := 0; i < int(config.General.SpawnRate); i++ {
			s.CreateParticle()
		}
	}

	l := s.Content
	if s.Tick % 60 == 0 { ebiten.SetWindowTitle("Project particles - Paricules: " + fmt.Sprint(l.Len()) + " - TPS: " + fmt.Sprint(int(ebiten.CurrentTPS()))) }

	for e := l.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
		// Update
		p.updateParticle()

		// KILL PARTICLES
		p.LifeSpan--
		if p.LifeSpan == 0 {
			p.hideParticle()
			p.KillState = 1
			continue
		} else if p.PositionX < -20 || p.PositionY < -20 {
			p.hideParticle()
			p.KillState = 1
			continue
		} else if  int(p.PositionY) > config.General.WindowSizeY + 20 || int(p.PositionX) > config.General.WindowSizeX + 20 {
			p.hideParticle()
			p.KillState = 1
			continue
		}
	}
	for e := l.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
		if p.KillState == 1 {
			l.Remove(e)
		}
	}
}
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
	l := s.Content
	s.Tick++

	s.CursorX, s.CursorY = ebiten.CursorPosition()
	if s.CursorX < 0 { s.CursorX = 0 }
	if s.CursorY < 0 { s.CursorY = 0 }
	if s.CursorX > config.General.WindowSizeX { s.CursorX = config.General.WindowSizeX }
	if s.CursorY > config.General.WindowSizeY { s.CursorY = config.General.WindowSizeY }
	
	if config.General.EventOnClick {
		s.onClick()
	}


	if (l.Len() < config.General.MaxParticles || config.General.MaxParticles == -1) {
		if config.General.SpawnRate < 1 && config.General.SpawnRate > 0 {
			if s.Tick % int(1 / config.General.SpawnRate) == 0 {
				s.CreateParticle()
			}
		} else {
			for i := 0; i < int(config.General.SpawnRate) && (l.Len() < config.General.MaxParticles || config.General.MaxParticles == -1); i++ {
				s.CreateParticle()
			}
		}
	}

	if s.Tick % 60 == 0 { ebiten.SetWindowTitle("Project particles - Particules: " + fmt.Sprint(l.Len()) + " - TPS: " + fmt.Sprint(int(ebiten.CurrentTPS()))) }

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
			if config.General.ToggleRespawn {
				p.respawn()
			} else {
				l.Remove(e)
			}
			
		}
	}
	s.UpdateMode()
}
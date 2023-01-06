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
	l := s.Content // On résucpère la liste de particule
	s.Tick++

	if (l.Len() < config.General.MaxParticles || config.General.MaxParticles == -1) { // Mettre "MaxParticles" à -1 permet de pouvoir générer un nombre infini de particules
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

	// Actualiser l'affichage du nombre de TPS toute les secondes
	if s.Tick % 60 == 0 { ebiten.SetWindowTitle("Project particles - Paricules: " + fmt.Sprint(l.Len()) + " - TPS: " + fmt.Sprint(int(ebiten.CurrentTPS()))) }

	for e := l.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
		// Update
		p.UpdateParticle()

		// Détruir les particules
		// Si les particules sortent de l'écran, on appel la fonction "HideParticle()" et le "KillState" de la particule passe à 1
		if p.PositionX < -10 || p.PositionY < -10 {
			p.HideParticle()
			p.KillState = 1
			continue
		} else if  int(p.PositionY) > config.General.WindowSizeY + 10 || int(p.PositionX) > config.General.WindowSizeX + 10 {
			p.HideParticle()
			p.KillState = 1
			continue
		}
	}

	// On parcour la liste de toutes les particules et si une à un "KillState" de 1, elle se supprime (appel de la fonction "Remove").
	for e := l.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
		if p.KillState == 1 {
			l.Remove(e)
		}
	}
}
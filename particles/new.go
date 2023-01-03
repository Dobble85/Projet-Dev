package particles

import (
	"container/list"
	"project-particles/config"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	var s System
	s.Tick = 0
	s.Content = list.New()
	s.ParticleSpawnX = config.General.SpawnX
	s.ParticleSpawnY = config.General.SpawnY
	s.ParticleColor = config.General.DefaultSpawnColor

	for nbParticles := 0; nbParticles < config.General.InitNumParticles; nbParticles++ {
		s.CreateParticle()
	}
	return s
}

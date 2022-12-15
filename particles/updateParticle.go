package particles

import (
	"project-particles/config"
)

func (p *Particle) updateParticle() {
	p.PositionX += p.SpeedX
	if config.General.ParticleGravity == 0 {
		p.PositionY += p.SpeedY
	} else {
		p.PositionY += p.SpeedY
		p.SpeedY *= config.General.ParticleGravity
	}
}
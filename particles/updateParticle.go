package particles

func (p *Particle) UpdateParticle() {
	// Ajouter une position X et Y à la position X et Y initial de la particule
	p.PositionX += p.SpeedX
	p.PositionY += p.SpeedY
}
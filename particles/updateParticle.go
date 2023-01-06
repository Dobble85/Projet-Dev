package particles

func (p *Particle) UpdateParticle() {
	p.PositionX += p.SpeedX
	p.PositionY += p.SpeedY
}
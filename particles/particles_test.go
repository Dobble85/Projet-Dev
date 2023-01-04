package particles

import (
	"testing"
	"project-particles/config"
)

func TestNewSystem(t *testing.T) {
	system := NewSystem()
    if system.Content.Len() != config.General.InitNumParticles {
		t.Error("Init num particles should be ", config.General.InitNumParticles)
	}

	if system.Tick != 0 {
		t.Error("Tick should be 0")
	}
	
	if system.ParticleSpawnX != config.General.SpawnX || system.ParticleSpawnY != config.General.SpawnY {
		t.Error("Particle spawn should be ", config.General.SpawnX, config.General.SpawnY)
	}

	if system.ParticleColor != config.General.DefaultSpawnColor {
		t.Error("Particle spawn color should be ", config.General.DefaultSpawnColor)
	}
}

func testCreateParticle(t *testing.T) {
	s := NewSystem()
	s.CreateParticle()

	if s.Content.Len()!= config.General.InitNumParticles+1 {
		t.Error("Init num particles should be ", config.General.InitNumParticles+1)
	}

	particle := s.Content.Back().Value.(*Particle)

	if (int(particle.PositionX) != config.General.SpawnX && config.General.RandomSpawnX) || (int(particle.PositionY) != config.General.SpawnY && config.General.RandomSpawnY) {
		t.Error("Particle spawn should be ", config.General.SpawnX, config.General.SpawnY, "when random spawn position is true")
	}

	if particle.Opacity != 1 && !config.General.RandomSpawnOpacity {
		t.Error("Particle spawn opacity should be ", 1, "when random spawn opacity is false")
	}

	if particle.LifeSpan != config.General.LifeSpan {
		t.Error("Particle life span should be ", config.General.LifeSpan)
	}
}

func testUpdateParticle(t *testing.T) {
	s := NewSystem()
    s.CreateParticle()
	p := s.Content.Back().Value.(*Particle)

    p.SpeedX = 1
	p.SpeedY = 1
	p.PositionX = 10
	p.PositionY = 10

	p.updateParticle()
	if p.PositionX!= 11 || p.PositionY!= 11 {
		t.Error("Particle position should be ", 11)
	}

	p.hideParticle()
	if p.Opacity != 0 {
		t.Error("Particle opacity should be ", 0)
	}
}


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

	particle := s.Content.Back()

	if (particle.PositionX!= config.General.SpawnX || particle.PositionY!= config.General.SpawnY) && config.General.Ran
}
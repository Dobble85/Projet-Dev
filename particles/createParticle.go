package particles

import (
	"project-particles/config"
	"math/rand"
	"time"
)

func (s *System) CreateParticle() {
	// POSITION
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	var ParticleX int
	var ParticleY int
	if config.General.RandomSpawnX {
		ParticleX = random.Int() % config.General.WindowSizeX
	} else {
		ParticleX = config.General.SpawnX
	}
	if config.General.RandomSpawnY {
		ParticleY = random.Int() % config.General.WindowSizeY
	} else {
		ParticleY = config.General.SpawnY
	}

	// SPEED
	var ParticleSpeedX float64 = random.Float64() * config.General.SpeedMultiplier
	var ParticleSpeedY float64 = random.Float64() * config.General.SpeedMultiplier

	if random.Int() % 2 == 0 {
		ParticleSpeedX = -ParticleSpeedX
	}
	if random.Int() % 2 == 0 {
		ParticleSpeedY = -ParticleSpeedY
	}

	// COLOR
	var ParticleColorR float64 = 0.4
	var ParticleColorG float64 = 0.6
	var ParticleColorB float64 = 1
	if config.General.RandomSpawnColor {
		ParticleColorR = float64(random.Int() % 100) / 100
		ParticleColorG = float64(random.Int() % 100) / 100
		ParticleColorB = float64(random.Int() % 100) / 100
	}

	// PUSH_FRONT
	s.Content.PushFront(&Particle{
		PositionX: float64(ParticleX),
		PositionY: float64(ParticleY),
		SpeedX: ParticleSpeedX,
		SpeedY: ParticleSpeedY,
		ScaleX:    1, ScaleY: 1,
		ColorRed: ParticleColorR, ColorGreen: ParticleColorG, ColorBlue: ParticleColorB,
		Opacity: float64(random.Int() % 100) / 100, LifeSpan: config.General.LifeSpan,
		KillState: 0,
	})
}
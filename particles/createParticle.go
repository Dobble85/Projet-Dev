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
		ParticleX = s.ParticleSpawnX
	}
	if config.General.RandomSpawnY {
		ParticleY = random.Int() % config.General.WindowSizeY
	} else {
		ParticleY = s.ParticleSpawnY
	}

	var ParticleRotation float64

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
	var ParticleColorR float64 = 0
	var ParticleColorG float64 = 0
	var ParticleColorB float64 = 0
	if config.General.RandomSpawnColor {
		ParticleColorR = float64(random.Int() % 100) / 100
		ParticleColorG = float64(random.Int() % 100) / 100
		ParticleColorB = float64(random.Int() % 100) / 100
	} else {
		switch {
		case config.General.DefaultSpawnColor == "white":
			ParticleColorR = 1
			ParticleColorG = 1
			ParticleColorB = 1
		case config.General.DefaultSpawnColor == "red":
			ParticleColorR = 1
		case config.General.DefaultSpawnColor == "blue":
			ParticleColorB = 1
		case config.General.DefaultSpawnColor == "green":
			ParticleColorG = 1
		case config.General.DefaultSpawnColor == "aqua":
			ParticleColorG = 1
			ParticleColorB = 1
		case config.General.DefaultSpawnColor == "yellow":
			ParticleColorR = 1
			ParticleColorG = 1
		case config.General.DefaultSpawnColor == "magenta":
			ParticleColorR = 1
			ParticleColorB = 1
		case config.General.DefaultSpawnColor == "orange":
			ParticleColorR = 1
			ParticleColorG = 0.65
		}
	}

	var ParticleOpacity float64 = 1
	if config.General.RandomSpawnOpacity {
		ParticleOpacity = float64(random.Int() % 100) / 100
	}

	// PUSH_FRONT
	s.Content.PushFront(&Particle{
		PositionX: float64(ParticleX),
		PositionY: float64(ParticleY),
		SpeedX: ParticleSpeedX,
		SpeedY: ParticleSpeedY,
		Rotation: ParticleRotation,
		ScaleX:    0.3, ScaleY: 0.3,
		ColorRed: ParticleColorR, ColorGreen: ParticleColorG, ColorBlue: ParticleColorB,
		Opacity: ParticleOpacity, LifeSpan: config.General.LifeSpan,
		KillState: 0,
	})
}
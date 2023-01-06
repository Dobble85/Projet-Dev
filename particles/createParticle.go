package particles

import (
	"project-particles/config"
	"math/rand"
	"time"
)

func (s *System) CreateParticle() {
	// POSITION

	// Créer une variable "random" pour définir de façon aléatoire la position de la particule
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	var ParticleX int
	var ParticleY int

	// Si la génération aléatoire en X est activé, alors assigner un nombre aléatoire qui ne dépasse pas la taille de l'écran à la variable "ParticleX", sinon, assigner "ParticleSpawnX" à "ParticleX"
	if config.General.RandomSpawnX {
		ParticleX = random.Int() % config.General.WindowSizeX
	} else {
		// Faire en sorte que quand on rentre une coordonné plus grande que la taille de l'écran, elle soit automatiquement réduite à la limite de la taille de l'écran
		if config.General.SpawnX > config.General.WindowSizeX {
			ParticleX = config.General.WindowSizeX
		} else {
			ParticleX = config.General.SpawnX
		}

		if config.General.SpawnX < 0 {
			ParticleX = 0
		}
	}

	// Même chose pour "ParticleY"
	if config.General.RandomSpawnY {
		ParticleY = random.Int() % config.General.WindowSizeY
	} else {
		if config.General.SpawnY > config.General.WindowSizeY {
			ParticleY = config.General.WindowSizeY
		} else {
			ParticleY = config.General.SpawnY
		}

		if config.General.SpawnY < 0 {
			ParticleY = 0
		}
	}

	var ParticleRotation float64

	// SPEED

	var ParticleSpeedX float64 = random.Float64() // Générer un valeur aléatoire
	var ParticleSpeedY float64 = random.Float64()

	// Si un nombre généré est paire, inverser la vitesse de la particule
	if random.Int() % 2 == 0 {
		ParticleSpeedX = -ParticleSpeedX
	}

	if random.Int() % 2 == 0 {
		ParticleSpeedY = -ParticleSpeedY
	}

	// COLOR

	// La couleur de base de la particule est blanche
	var ParticleColorR float64 = 1
	var ParticleColorG float64 = 1
	var ParticleColorB float64 = 1

	var ParticleOpacity float64 = 1

	// PUSH_FRONT

	s.Content.PushFront(&Particle{

		PositionX: float64(ParticleX),
		PositionY: float64(ParticleY),

		SpeedX: ParticleSpeedX,
		SpeedY: ParticleSpeedY,

		Rotation: ParticleRotation,

		// Taille : 1 est la taille de base
		ScaleX: 1, 
		ScaleY: 1,

		ColorRed: ParticleColorR, ColorGreen: ParticleColorG, ColorBlue: ParticleColorB,
		Opacity: ParticleOpacity,

		// Etat de la particule : 0 -> Vivante | 1 -> Morte
		KillState: 0,
	})
}
package particles

import (
	"container/list"
	"project-particles/config"
	"image"
	_ "image/jpeg"
	_ "image/gif"
	"net/http"
	"github.com/hajimehoshi/ebiten/v2"
	"strings"
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

	if config.General.Preset[:4] == "img-" {
		url := config.General.Preset[4:]
		response, _ := http.Get(url)
		// Décoder l'image
		img, _, _ := image.Decode(response.Body)
		s.img = img
		imgBounds := img.Bounds()
		//strings.Contains(url, ".gif")
		config.General.WindowSizeX = imgBounds.Max.X
		config.General.WindowSizeY = imgBounds.Max.Y
		ebiten.SetWindowSize(imgBounds.Max.X, imgBounds.Max.Y)

		config.General.MaxParticles = (imgBounds.Max.X * imgBounds.Max.X) / 5
		config.General.InitNumParticles = (imgBounds.Max.X * imgBounds.Max.X) / 1000
		config.General.SpawnRate = float64((imgBounds.Max.X * imgBounds.Max.X) / 1000)
	} else {
		s.img = image.NewRGBA(image.Rect(0, 0, 10, 10))
	}

	s.ParticleSpawnX = config.General.SpawnX
	s.ParticleSpawnY = config.General.SpawnY
	s.ParticleColor = config.General.DefaultSpawnColor



	for nbParticles := 0; nbParticles < config.General.InitNumParticles; nbParticles++ {
		s.CreateParticle()
	}
	return s
}
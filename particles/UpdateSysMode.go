package particles

import (
	"project-particles/config"
)

func (s *System) UpdateMode() {
	var mode string = config.General.Preset
	switch {
	case mode == "fr":
		for e := s.Content.Front(); e != nil; e = e.Next() {
			p := e.Value.(*Particle)

			if int(p.PositionX) <= config.General.WindowSizeX / 3 {
				p.ColorRed = 0
				p.ColorGreen = 0
                p.ColorBlue = 1
			} else if int(p.PositionX) <= (config.General.WindowSizeX / 3) * 2 {
				p.ColorRed = 1
				p.ColorGreen = 1
                p.ColorBlue = 1
			} else {
				p.ColorRed = 1
				p.ColorGreen = 0
                p.ColorBlue = 0
			}
		}

	case len(mode) > 4 && mode[:4] == "img-":
		for e := s.Content.Front(); e != nil; e = e.Next() {
			p := e.Value.(*Particle)
			// Récupérer les valeurs RGB de la couleur
			r, g, b, a := s.img.At(int(p.PositionX), int(p.PositionY)).RGBA()
			p.ColorRed = float64(uint8(r)) / 255
			p.ColorGreen = float64(uint8(g)) / 255
			p.ColorBlue = float64(uint8(b)) / 255
			p.Opacity = float64(uint8(a)) / 255
		}
	}
}

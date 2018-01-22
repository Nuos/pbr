package material

import "github.com/hunterloftis/pbr/rgb"

var (
	Default    = Plastic(1, 1, 1, 0.7)
	Gold       = Metal(1.022, 0.782, 0.344, 0.9, 0.9)
	GreenGlass = Glass(0.2, 1, 0.1, 0.95)
	Chrome     = Metal(1, 1, 1, 1, 1)
)

// Light constructs a new light
// r, g, b (0-Inf) specifies the light color
func Light(r, g, b float64) *Material {
	return New(MaterialDesc{
		Light: rgb.Energy{r, g, b},
	})
}

// Plastic constructs a new plastic material
// r, g, b (0-1) controls the color
// gloss (0-1) controls the microfacet roughness (how polished the surface looks)
func Plastic(r, g, b float64, gloss float64) *Material {
	return New(MaterialDesc{
		Color:   rgb.Energy{r, g, b},
		Fresnel: rgb.Energy{0.04, 0.04, 0.04},
		Gloss:   gloss,
	})
}

// Lambert constructs a new lambert material
// r, g, b (0-1) controls the color
func Lambert(r, g, b float64) *Material {
	return New(MaterialDesc{
		Color:   rgb.Energy{r, g, b},
		Fresnel: rgb.Energy{0.02, 0.02, 0.02},
	})
}

// Metal constructs a new metal material
// r, g, b (0-1) controls the fresnel color
// gloss (0-1) controls the microfacet roughness (how polished the surface looks)
func Metal(r, g, b, gloss, metal float64) *Material {
	return New(MaterialDesc{
		Fresnel: rgb.Energy{r, g, b},
		Color:   rgb.Energy{r, g, b},
		Gloss:   gloss,
		Metal:   metal,
	})
}

// Glass constructs a new glass material
// r, g, b (0-1) controls the transmission of the glass (beer-lambert)
// gloss (0-1) controls the microfacet roughness (how polished the surface looks)
func Glass(r, g, b, gloss float64) *Material {
	return New(MaterialDesc{
		Color:    rgb.Energy{r, g, b},
		Fresnel:  rgb.Energy{0.042, 0.042, 0.042},
		Transmit: 1,
		Gloss:    gloss,
	})
}
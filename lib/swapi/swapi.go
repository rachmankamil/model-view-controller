package swapi

import "time"

type Planets struct {
	Climate        string    `json:"climate"`
	Created        time.Time `json:"created"`
	Diameter       string    `json:"diameter"`
	Edited         time.Time `json:"edited"`
	Films          []string  `json:"films"`
	Gravity        string    `json:"gravity"`
	Name           string    `json:"name"`
	OrbitalPeriod  string    `json:"orbital_period"`
	Population     string    `json:"population"`
	Residents      []string  `json:"residents"`
	RotationPeriod string    `json:"rotation_period"`
	SurfaceWater   string    `json:"surface_water"`
	Terrain        string    `json:"terrain"`
	URL            string    `json:"url"`
}

package gins

// Gin attributes of a Gin
type Gin struct {
	Name     string
	Origin   string
	Launched string
}

var cabinet []Gin

func init() {

	cabinet = addGin(Gin{"GunPowder", "Drumshanbo, Ireland", "2016"})
	cabinet = addGin(Gin{"Bombay Sapphire", "Laverstoke Mill, Hampshire", "1987"})
	cabinet = addGin(Gin{"Hendricks", "Girvan, Scotland", "1999"})
}

// All returns an array of Gins
func All() []Gin {
	return cabinet
}

func addGin(g Gin) []Gin {
	return append(cabinet, g)
}

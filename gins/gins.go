package gins

type Gin struct {
	Name     string
	Origin   string
	Launched string
}

var cabnet []Gin

func init() {

	cabnet = addGin(Gin{"GunPowder", "Drumshanbo, Ireland", "2016"})
	cabnet = addGin(Gin{"Bombay Sapphire", "Laverstoke Mill, Hampshire", "1987"})
	cabnet = addGin(Gin{"Hendricks", "Ireland", "1999"})
}

func All() []Gin {
	return cabnet
}

func addGin(g Gin) []Gin {
	return append(cabnet, g)
}

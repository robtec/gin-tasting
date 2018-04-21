package tonics

// Tonic attributes of a Tonic
type Tonic struct {
	Name string
}

var fridge []Tonic

func init() {

	fridge = addTonic(Tonic{"Fever Tree"})
	fridge = addTonic(Tonic{"Q Tonic"})
	fridge = addTonic(Tonic{"Fever Tree Elderflower"})
}

// All returns an array of Tonics
func All() []Tonic {
	return fridge
}

func addTonic(t Tonic) []Tonic {
	return append(fridge, t)
}

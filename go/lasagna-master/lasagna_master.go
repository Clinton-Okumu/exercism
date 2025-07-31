package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layer []string, PreparationTimeinMinutes int) int {
	if PreparationTimeinMinutes == 0 {
		PreparationTimeinMinutes = 2
	}
	return PreparationTimeinMinutes * len(layer)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	gramsOfNoodles := 0
	litresOfSauce := 0.0
	for _, layer := range layers {
		switch layer {
		case "noodles":
			gramsOfNoodles += 50
		case "sauce":
			litresOfSauce += 0.2
		}
	}
	return gramsOfNoodles, litresOfSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	secret := friendList[len(friendList)-1]
	myList[len(myList)-1] = secret
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(portions []float64, numberOfPortions int) []float64 {
	scale := float64(numberOfPortions) / 2.0
	scaled := make([]float64, len(portions))
	for i, amount := range portions {
		scaled[i] = amount * scale
	}
	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

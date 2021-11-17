package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg int) int {
	if avg == 0 {
		avg = 2
	}

	lay := len(layers)

	return lay * avg
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += .2
		}
	}

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friend, own []string) []string {
	secret := friend[len(friend)-1]
	improved := append(own, secret)
	return improved
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) (improved []float64) {

	portion := float64(portions)

	for _, quantity := range quantities {
		qty := (quantity / 2.0) * portion
		improved = append(improved, qty)
	}

	return
}

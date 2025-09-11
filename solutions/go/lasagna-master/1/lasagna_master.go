package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) (totalEstimate int) {
    if avgPrepTime == 0 {avgPrepTime = 2}
    totalEstimate = len(layers) * avgPrepTime
    return
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodleGrams int, sauceLiters float64) {
    for _, l := range layers {
        switch l {
            case "noodles" : noodleGrams += 50
            case "sauce" : sauceLiters += 0.2
        }
    }
    return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendRecipe, ownRecipe []string) {
    ownLength := len(ownRecipe)
    friendLength := len(friendRecipe)
    ownRecipe[ownLength - 1] = friendRecipe[friendLength - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) (amountsNeeded []float64) {
    for _, v := range quantities {
        amountsNeeded = append(amountsNeeded, v / 2.0 * float64(portions))
    }
    return
}

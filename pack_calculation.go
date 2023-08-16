package main

var PackSizes = []int{5000, 2000, 1000, 500, 250}

func CalculatePacks(oQuantity int) map[int]int {
	result := make(map[int]int)

	for _, packSize := range PackSizes {
		if oQuantity >= packSize {
			numPacks := oQuantity / packSize
			result[packSize] = numPacks
			oQuantity -= numPacks * packSize
		}
	}

	return result
}
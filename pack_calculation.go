package main

// recursive is a recursive function that takes a order quantity, packSizes, and accumulatedSizes.
// It returns a slice of accumulated sizes that sum up to or exceed the given order quantity.
func recursive(oQuantity int, packSizes []int, accumulatedSizes []int) []int {
	sum := GetSum(accumulatedSizes) // Calculate the sum of accumulated sizes
	if sum >= oQuantity { // If the sum is greater or equal to the order quantity, return accumulated sizes
		return accumulatedSizes
	}

	for i, size := range packSizes {
		prevNumsSum := GetSum(packSizes[0 : i+1]) // Calculate the sum of pack sizes up to the current index

		// Check various conditions to decide whether to append the size and recurse
		if sum+size >= oQuantity || (prevNumsSum+sum >= oQuantity) || i == len(packSizes)-1 || (GetSum(packSizes[i+1:i+2])-size >= oQuantity) {
			accumulatedSizes = append(accumulatedSizes, size) // Append the size to accumulated sizes
			return recursive(oQuantity, packSizes, accumulatedSizes) // Recurse with the updated accumulated sizes
		}
	}
	return accumulatedSizes // Return the accumulated sizes if no conditions are met
}

// getOptimalSlice takes a slice of integer slices and returns the slice with the smallest sum.
func getOptimalSlice(arr [][]int) []int {
	var returnVal []int = arr[0] // Initialize return value with the first slice
	var optimalArrSum int = GetSum(arr[0]) // Initialize optimal sum with the sum of the first slice
	
	for _, array := range arr {
		sum := GetSum(array) // Calculate the sum of the current slice
		if sum < optimalArrSum { // If the sum is less than the optimal sum, update the return value and optimal sum
			optimalArrSum = sum
			returnVal = array
		}
	}
	return returnVal // Return the optimal slice
}

// CalculatePacks takes a order quantity and calculates the optimal packing of sizes.
// It returns a map of pack sizes and their quantities.
func CalculatePacks(oQuantity int, sizes map[int]int) map[int]int{
	packSizes := OrderPackSizes(sizes) // Order the pack sizes
	reversePackSizes := Reverse(packSizes) // Reverse the ordered pack sizes
	var allCalculations = [][]int{} // Initialize a slice to store all calculations
	for _, size := range reversePackSizes {
		acc := []int{size} // Initialize accumulated sizes with the current size
		calculation := recursive(oQuantity, packSizes, acc) // Call the recursive function
		allCalculations = append(allCalculations, calculation) // Append the calculation to all calculations
	}
	optimalSlice := getOptimalSlice(allCalculations) // Get the optimal slice from all calculations
	for _, size := range optimalSlice {
		sizes[size]++ // Increment the quantity for each size in the optimal slice
	}
	return sizes // Return the final order
}
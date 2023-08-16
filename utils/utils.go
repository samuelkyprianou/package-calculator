package utils

import "sort"

// GetSum takes an array of integers and returns the sum of all the elements.
func GetSum(arr []int) int {
	sum := 0 // Initialize the sum variable to 0
	for _, v := range arr { // Iterate through the array
		sum += v // Add each element's value to the sum
	}
	return sum // Return the total sum
}

// OrderPackSizes takes a map where the keys are pack sizes and the values are quantities.
// It returns a sorted slice of the pack sizes in ascending order.
func OrderPackSizes(order map[int]int) []int {
	var packSizes []int // Initialize an empty slice to store pack sizes
	for size := range order { // Iterate through the keys of the map (i.e., the pack sizes)
		packSizes = append(packSizes, size) // Append each pack size to the slice
	}
	sort.Ints(packSizes) // Sort the pack sizes in ascending order
	return packSizes // Return the sorted pack sizes
}

// Reverse takes an array of integers and returns a new array with the elements in reverse order.
func Reverse(arr []int) []int {
	reversed := make([]int, len(arr)) // Initialize a new array with the same length as the input
	for i, v := range arr { // Iterate through the array
		reversed[len(arr)-1-i] = v // Assign the elements in reverse order to the new array
	}
	return reversed // Return the reversed array
}
package main

import (
	"reflect"
	"testing"
)

// Function to create a new instance of the mock pack sizes
func newMockPackSizes() map[int]int {
	return map[int]int{
		250:  0,
		500:  0,
		1000: 0,
		2000: 0,
		5000: 0,
	}
}

func TestCalculatePacks1(t *testing.T) {
	mockPackSizes := newMockPackSizes()
	expected := map[int]int{
		250:  1,
		500: 0,
		1000: 0,
		2000: 0,
		5000: 0,
	}

	result := CalculatePacks(1, mockPackSizes)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %%v: %v GOT %%v: %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}

func TestCalculatePacks250(t *testing.T) {
	mockPackSizes := newMockPackSizes()
	expected := map[int]int{
		250:  1,
		500: 0,
		1000: 0,
		2000: 0,
		5000: 0,
	}

	result := CalculatePacks(250, mockPackSizes)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %v GOT %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}

func TestCalculatePacks251(t *testing.T) {
	mockPackSizes := newMockPackSizes()
	expected := map[int]int{
		250: 0,
		500:  1,
		1000: 0,
		2000: 0,
		5000: 0,
	}

	result := CalculatePacks(251, mockPackSizes)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %v GOT %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}

func TestCalculatePacks501(t *testing.T) {
	mockPackSizes := newMockPackSizes()
	expected := map[int]int{
		250:  1,
		500:  1,
		1000: 0,
		2000: 0,
		5000: 0,
	}

	result := CalculatePacks(501, mockPackSizes)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %v GOT %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}

func TestCalculatePacks8500(t *testing.T) {
	mockPackSizes := newMockPackSizes()
	expected := map[int]int{
		250: 0,
		500:  1,
		1000: 1,
		2000: 1,
		5000: 1,
	}

	result := CalculatePacks(8500, mockPackSizes)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %v GOT %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}

func TestCalculatePacks12001(t *testing.T) {
	mockPackSizes := newMockPackSizes()
	expected := map[int]int{
		250:  1,
		500: 0,
		1000: 0,
		2000: 1,
		5000: 2,
	}

	result := CalculatePacks(12001, mockPackSizes)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %v GOT %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}
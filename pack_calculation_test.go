package main

import (
	"reflect"
	"testing"
)

func TestCalculatePacks1(t *testing.T) {
	expected := map[int]int{
		250:  1,
	}

	result := CalculatePacks(1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %%v: %v GOT %%v: %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}

func TestCalculatePacks250(t *testing.T) {
	expected := map[int]int{
		250:  1,
	}

	result := CalculatePacks(250)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %v GOT %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}

func TestCalculatePacks251(t *testing.T) {
	expected := map[int]int{
		500:  1,
	}

	result := CalculatePacks(251)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %v GOT %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}

func TestCalculatePacks501(t *testing.T) {
	expected := map[int]int{
		250:  1,
		500:  1,
	}

	result := CalculatePacks(501)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %v GOT %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}

func TestCalculatePacks8500(t *testing.T) {
	expected := map[int]int{
		500:  1,
		1000: 1,
		2000: 1,
		5000: 1,
	}

	result := CalculatePacks(8500)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %v GOT %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}

func TestCalculatePacks12001(t *testing.T) {
	expected := map[int]int{
		250:  1,
		2000: 1,
		5000: 2,
	}

	result := CalculatePacks(12001)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %v GOT %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}
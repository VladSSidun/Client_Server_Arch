//  go test ./... -v
package main

import (
	"testing"
)

// ===== ТЕСТИ ЗАВДАННЯ 2 =====

func TestMultiplyMatrices_Valid(t *testing.T) {
	A := [][]int{{1, 2}, {3, 4}}
	B := [][]int{{5, 6}, {7, 8}}
	result, err := multiplyMatrices(A, B)
	if err != nil {
		t.Fatalf("очікували nil, отримали: %v", err)
	}
	// [1*5+2*7, 1*6+2*8] = [19, 22]
	// [3*5+4*7, 3*6+4*8] = [43, 50]
	expected := [][]int{{19, 22}, {43, 50}}
	for i := range expected {
		for j := range expected[i] {
			if result[i][j] != expected[i][j] {
				t.Errorf("[%d][%d]: очікували %d, отримали %d",
					i, j, expected[i][j], result[i][j])
			}
		}
	}
}

func TestMultiplyMatrices_Identity(t *testing.T) {
	A := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// Одинична матриця
	I := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	result, err := multiplyMatrices(A, I)
	if err != nil {
		t.Fatalf("очікували nil, отримали: %v", err)
	}
	// A * I = A
	for i := range A {
		for j := range A[i] {
			if result[i][j] != A[i][j] {
				t.Errorf("[%d][%d]: очікували %d, отримали %d",
					i, j, A[i][j], result[i][j])
			}
		}
	}
}

func TestMultiplyMatrices_IncompatibleSizes(t *testing.T) {
	A := [][]int{{1, 2}, {3, 4}}
	B := [][]int{{1, 2}, {3, 4}, {5, 6}}
	_, err := multiplyMatrices(A, B)
	if err == nil {
		t.Error("очікували помилку для несумісних розмірів")
	}
}

func TestMultiplyMatrices_Empty(t *testing.T) {
	_, err := multiplyMatrices([][]int{}, [][]int{{1, 2}})
	if err == nil {
		t.Error("очікували помилку для порожньої матриці")
	}
}

// ===== ТЕСТИ ЗАВДАННЯ 7 =====

func TestSumOfAverages_Valid(t *testing.T) {
	slices := [][]float64{
		{1.0, 2.0, 3.0, 4.0, 5.0}, // avg = 3.0
		{10.0, 20.0, 30.0},         // avg = 20.0
		{2.5, 7.5},                  // avg = 5.0
	}
	result, err := sumOfAverages(slices)
	if err != nil {
		t.Fatalf("очікували nil, отримали: %v", err)
	}
	expected := 28.0
	if result != expected {
		t.Errorf("очікували %.2f, отримали %.2f", expected, result)
	}
}

func TestSumOfAverages_SingleElements(t *testing.T) {
	slices := [][]float64{{5.0}, {10.0}, {15.0}}
	result, err := sumOfAverages(slices)
	if err != nil {
		t.Fatalf("очікували nil, отримали: %v", err)
	}
	if result != 30.0 {
		t.Errorf("очікували 30.00, отримали %.2f", result)
	}
}

func TestSumOfAverages_WrongCount(t *testing.T) {
	_, err := sumOfAverages([][]float64{{1, 2}, {3, 4}})
	if err == nil {
		t.Error("очікували помилку для неправильної кількості зрізів")
	}
}

// ===== ТЕСТИ ЗАВДАННЯ 8 =====

func TestSearchInSlice_Found(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	res, err := searchInSlice(slice, 50)
	if err != nil {
		t.Fatalf("очікували nil, отримали: %v", err)
	}
	if !res.Found {
		t.Error("очікували знайти 50 у зрізі")
	}
	if slice[res.Index] != 50 {
		t.Errorf("неправильний індекс: slice[%d] = %d", res.Index, slice[res.Index])
	}
}

func TestSearchInSlice_NotFound(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	res, err := searchInSlice(slice, 99)
	if err != nil {
		t.Fatalf("очікували nil, отримали: %v", err)
	}
	if res.Found {
		t.Error("не очікували знайти 99 у зрізі")
	}
}

func TestSearchInSlice_Empty(t *testing.T) {
	_, err := searchInSlice([]int{}, 42)
	if err == nil {
		t.Error("очікували помилку для порожнього зрізу")
	}
}

func TestSearchInSlice_FirstElement(t *testing.T) {
	slice := []int{42, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res, err := searchInSlice(slice, 42)
	if err != nil {
		t.Fatalf("очікували nil, отримали: %v", err)
	}
	if !res.Found || res.Index != 0 {
		t.Errorf("очікували знайти 42 на індексі 0, отримали Found=%v Index=%d",
			res.Found, res.Index)
	}
}
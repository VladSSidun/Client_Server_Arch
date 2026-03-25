package calculator

import "testing"

// testCase — структура для зберігання вхідних даних та очікуваного результату
type testCase struct {
	arg1, arg2, expected int
}

// TestAdd — табличний тест для функції Add
func TestAdd(t *testing.T) {
	cases := []testCase{
		{2, 3, 5},
		{0, 0, 0},
		{-5, 5, 0},
		{-4, -6, -10},
		{100, 200, 300},
	}
	for _, tc := range cases {
		result := Add(tc.arg1, tc.arg2)
		
		if result != tc.expected {
			t.Errorf("Add(%d, %d): отримано %d, очікувалося %d",
				tc.arg1, tc.arg2, result, tc.expected)
		}
	}
}

// TestSub — табличний тест для функції Sub
func TestSub(t *testing.T) {
	cases := []testCase{
		{10, 3, 7},
		{0, 0, 0},
		{-5, -5, 0},
		{5, 10, -5},
		{100, 1, 99},
	}
	for _, tc := range cases {
		result := Sub(tc.arg1, tc.arg2)
		if result != tc.expected {
			t.Errorf("Sub(%d, %d): отримано %d, очікувалося %d",
				tc.arg1, tc.arg2, result, tc.expected)
		}
	}
}

// TestMul — табличний тест для функції Mul
func TestMul(t *testing.T) {
	cases := []testCase{
		{3, 4, 12},
		{0, 100, 0},
		{-2, 5, -10},
		{-3, -3, 9},
		{7, 7, 49},
	}
	for _, tc := range cases {
		result := Mul(tc.arg1, tc.arg2)
		if result != tc.expected {
			t.Errorf("Mul(%d, %d): отримано %d, очікувалося %d",
				tc.arg1, tc.arg2, result, tc.expected)
		}
	}
}
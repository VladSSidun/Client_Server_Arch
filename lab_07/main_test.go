// go test ./... -v
package main

import (
	"os"
	"strings"
	"testing"
)

// ===== ТЕСТИ ЗАВДАННЯ 6 =====

func TestCompareFiles_Valid(t *testing.T) {
	os.WriteFile("t_f1.txt", []byte("apple\nbanana\ncherry"), 0644)
	os.WriteFile("t_f2.txt", []byte("banana\ncherry\ndate"), 0644)
	defer os.Remove("t_f1.txt")
	defer os.Remove("t_f2.txt")
	defer os.Remove("t_diff.txt")

	if err := compareFiles("t_f1.txt", "t_f2.txt", "t_diff.txt"); err != nil {
		t.Fatalf("очікували nil, отримали: %v", err)
	}

	data, _ := os.ReadFile("t_diff.txt")
	content := string(data)
	if !strings.Contains(content, "apple") {
		t.Error("очікували 'apple' у результаті")
	}
	if !strings.Contains(content, "date") {
		t.Error("очікували 'date' у результаті")
	}
	if strings.Contains(content, "banana") {
		t.Error("'banana' є в обох файлах — не повинна бути у результаті")
	}
}

func TestCompareFiles_MissingFile(t *testing.T) {
	err := compareFiles("nonexistent1.txt", "nonexistent2.txt", "out.txt")
	if err == nil {
		t.Error("очікували помилку для неіснуючого файлу")
	}
}

func TestCompareFiles_IdenticalFiles(t *testing.T) {
	os.WriteFile("t_same1.txt", []byte("line1\nline2"), 0644)
	os.WriteFile("t_same2.txt", []byte("line1\nline2"), 0644)
	defer os.Remove("t_same1.txt")
	defer os.Remove("t_same2.txt")
	defer os.Remove("t_same_out.txt")

	if err := compareFiles("t_same1.txt", "t_same2.txt", "t_same_out.txt"); err != nil {
		t.Fatalf("очікували nil, отримали: %v", err)
	}
	data, _ := os.ReadFile("t_same_out.txt")
	if strings.TrimSpace(string(data)) != "" {
		t.Error("ідентичні файли — результат повинен бути порожнім")
	}
}

// ===== ТЕСТИ ЗАВДАННЯ 11 =====

func TestFindMaxMin_Valid(t *testing.T) {
	os.WriteFile("t_nums.txt", []byte("3 1 4 1 5 9 2 6"), 0644)
	defer os.Remove("t_nums.txt")
	defer os.Remove("t_maxmin.txt")

	maxVal, minVal, err := findMaxMin("t_nums.txt", "t_maxmin.txt")
	if err != nil {
		t.Fatalf("очікували nil, отримали: %v", err)
	}
	if maxVal != 9 {
		t.Errorf("очікували max=9, отримали: %d", maxVal)
	}
	if minVal != 1 {
		t.Errorf("очікували min=1, отримали: %d", minVal)
	}
}

func TestFindMaxMin_NegativeNumbers(t *testing.T) {
	os.WriteFile("t_neg.txt", []byte("-5 -1 -10 -3"), 0644)
	defer os.Remove("t_neg.txt")
	defer os.Remove("t_neg_out.txt")

	maxVal, minVal, err := findMaxMin("t_neg.txt", "t_neg_out.txt")
	if err != nil {
		t.Fatalf("очікували nil, отримали: %v", err)
	}
	if maxVal != -1 {
		t.Errorf("очікували max=-1, отримали: %d", maxVal)
	}
	if minVal != -10 {
		t.Errorf("очікували min=-10, отримали: %d", minVal)
	}
}

func TestFindMaxMin_EmptyFile(t *testing.T) {
	os.WriteFile("t_empty.txt", []byte(""), 0644)
	defer os.Remove("t_empty.txt")

	_, _, err := findMaxMin("t_empty.txt", "out.txt")
	if err == nil {
		t.Error("очікували помилку для порожнього файлу")
	}
}

func TestFindMaxMin_MissingFile(t *testing.T) {
	_, _, err := findMaxMin("no_such_file.txt", "out.txt")
	if err == nil {
		t.Error("очікували помилку для неіснуючого файлу")
	}
}

// ===== ТЕСТИ ЗАВДАННЯ 13 =====

func TestMergeUniqueWords_Valid(t *testing.T) {
	os.WriteFile("t_w1.txt", []byte("hello world foo"), 0644)
	os.WriteFile("t_w2.txt", []byte("world bar baz"), 0644)
	defer os.Remove("t_w1.txt")
	defer os.Remove("t_w2.txt")
	defer os.Remove("t_unique.txt")

	if err := mergeUniqueWords([]string{"t_w1.txt", "t_w2.txt"}, "t_unique.txt"); err != nil {
		t.Fatalf("очікували nil, отримали: %v", err)
	}

	data, _ := os.ReadFile("t_unique.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	// Унікальних слів: hello, world, foo, bar, baz = 5
	if len(lines) != 5 {
		t.Errorf("очікували 5 унікальних слів, отримали: %d", len(lines))
	}
}

func TestMergeUniqueWords_EmptyList(t *testing.T) {
	err := mergeUniqueWords([]string{}, "out.txt")
	if err == nil {
		t.Error("очікували помилку для порожнього списку файлів")
	}
}

func TestMergeUniqueWords_MissingFile(t *testing.T) {
	err := mergeUniqueWords([]string{"nonexistent.txt"}, "out.txt")
	if err == nil {
		t.Error("очікували помилку для неіснуючого файлу")
	}
}

// ===== ТЕСТИ ЗАВДАННЯ 15 =====

func TestSumEvenNumbers_Valid(t *testing.T) {
	os.WriteFile("t_even.txt", []byte("1 2 3 4 5 6"), 0644)
	defer os.Remove("t_even.txt")
	defer os.Remove("t_even_out.txt")

	sum, err := sumEvenNumbers("t_even.txt", "t_even_out.txt")
	if err != nil {
		t.Fatalf("очікували nil, отримали: %v", err)
	}
	// 2 + 4 + 6 = 12
	if sum != 12 {
		t.Errorf("очікували 12, отримали: %d", sum)
	}
}

func TestSumEvenNumbers_NoEven(t *testing.T) {
	os.WriteFile("t_odd.txt", []byte("1 3 5 7 9"), 0644)
	defer os.Remove("t_odd.txt")
	defer os.Remove("t_odd_out.txt")

	sum, err := sumEvenNumbers("t_odd.txt", "t_odd_out.txt")
	if err != nil {
		t.Fatalf("очікували nil, отримали: %v", err)
	}
	if sum != 0 {
		t.Errorf("очікували 0, отримали: %d", sum)
	}
}

func TestSumEvenNumbers_MissingFile(t *testing.T) {
	_, err := sumEvenNumbers("no_file.txt", "out.txt")
	if err == nil {
		t.Error("очікували помилку для неіснуючого файлу")
	}
}

// ===== ТЕСТИ ЗАВДАННЯ 16 =====

func TestLoadPizzeria_Valid(t *testing.T) {
	input := `{"pizzeria":{"name":"Test Pizza","location":"123 St","menu":[]}}`
	os.WriteFile("t_pizzeria.json", []byte(input), 0644)
	defer os.Remove("t_pizzeria.json")

	pf, err := loadPizzeria("t_pizzeria.json")
	if err != nil {
		t.Fatalf("очікували nil, отримали: %v", err)
	}
	if pf.Pizzeria.Name != "Test Pizza" {
		t.Errorf("очікували 'Test Pizza', отримали: %s", pf.Pizzeria.Name)
	}
}

func TestLoadPizzeria_InvalidJSON(t *testing.T) {
	os.WriteFile("t_bad.json", []byte("{not valid json"), 0644)
	defer os.Remove("t_bad.json")

	_, err := loadPizzeria("t_bad.json")
	if err == nil {
		t.Error("очікували помилку для невалідного JSON")
	}
}

func TestLoadPizzeria_MissingFile(t *testing.T) {
	_, err := loadPizzeria("no_such.json")
	if err == nil {
		t.Error("очікували помилку для неіснуючого файлу")
	}
}

func TestSavePizzeria_Valid(t *testing.T) {
	pf := &PizzeriaFile{
		Pizzeria: PizzeriaData{
			Name:     "Test",
			Location: "Test St",
			Menu:     []Pizza{{Name: "TestPizza", Price: 9.99}},
		},
	}
	defer os.Remove("t_save.json")

	if err := savePizzeria("t_save.json", pf); err != nil {
		t.Fatalf("очікували nil, отримали: %v", err)
	}

	// Перевіряємо що файл існує і містить потрібні дані
	data, err := os.ReadFile("t_save.json")
	if err != nil {
		t.Fatal("файл не був створений")
	}
	if !strings.Contains(string(data), "Test") {
		t.Error("збережений файл не містить очікуваних даних")
	}
}
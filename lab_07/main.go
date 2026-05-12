package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println("       ЛАБОРАТОРНА РОБОТА №7 | ВАРІАНТ 11         ")
	fmt.Println("==================================================")

	Task6()
	divider()
	Task11()
	divider()
	Task13()
	divider()
	Task15()
	divider()
	Task16()

	fmt.Println("\n==================================================")
	fmt.Println("               ВИКОНАННЯ ЗАВЕРШЕНО                ")
	fmt.Println("==================================================")
}

func divider() {
	fmt.Println("\n--------------------------------------------------")
}

// ================================================
// ДОПОМІЖНІ ФУНКЦІЇ
// ================================================

// readLines читає файл рядок за рядком через bufio.Scanner
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// readNumbers читає всі цілі числа з файлу (розділені пробілами або новими рядками)
func readNumbers(path string) ([]int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	fields := strings.Fields(string(data))
	if len(fields) == 0 {
		return nil, fmt.Errorf("файл %q не містить чисел", path)
	}

	var nums []int
	for _, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			return nil, fmt.Errorf("некоректне значення %q у файлі %q", f, path)
		}
		nums = append(nums, n)
	}
	return nums, nil
}

// ================================================
// ЗАВДАННЯ 6: Порівняння двох файлів
// ================================================

// compareFiles знаходить рядки що є тільки в одному з двох файлів,
// виводить у термінал та записує у вихідний файл
func compareFiles(path1, path2, outputPath string) error {
	lines1, err := readLines(path1)
	if err != nil {
		return fmt.Errorf("compareFiles: %w", err)
	}
	lines2, err := readLines(path2)
	if err != nil {
		return fmt.Errorf("compareFiles: %w", err)
	}

	// Будуємо set для кожного файлу
	set1 := make(map[string]bool)
	for _, l := range lines1 {
		if strings.TrimSpace(l) != "" {
			set1[l] = true
		}
	}
	set2 := make(map[string]bool)
	for _, l := range lines2 {
		if strings.TrimSpace(l) != "" {
			set2[l] = true
		}
	}

	// Знаходимо рядки що є тільки в одному файлі
	var result []string
	for l := range set1 {
		if !set2[l] {
			result = append(result, fmt.Sprintf("[тільки у файлі 1] %s", l))
		}
	}
	for l := range set2 {
		if !set1[l] {
			result = append(result, fmt.Sprintf("[тільки у файлі 2] %s", l))
		}
	}
	sort.Strings(result)

	fmt.Println("[Унікальні рядки:]")
	for _, l := range result {
		fmt.Println(" ", l)
	}

	return os.WriteFile(outputPath, []byte(strings.Join(result, "\n")), 0644)
}

func Task6() {
	fmt.Println("---> ЗАВДАННЯ №6 (Порівняння двох файлів)")

	os.WriteFile("file1.txt", []byte("apple\nbanana\ncherry\ndate"), 0644)
	os.WriteFile("file2.txt", []byte("banana\ncherry\nelderberry\nfig"), 0644)
	defer func() {
		os.Remove("file1.txt")
		os.Remove("file2.txt")
		os.Remove("diff_output.txt")
	}()

	if err := compareFiles("file1.txt", "file2.txt", "diff_output.txt"); err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
		return
	}
	fmt.Println("[OK] Результат записано у diff_output.txt")

	// Тест з неіснуючим файлом
	fmt.Println("\n[ТЕСТ] Неіснуючий файл:")
	if err := compareFiles("missing.txt", "file2.txt", "out.txt"); err != nil {
		fmt.Printf("[ОЧІКУВАНА ПОМИЛКА] %v\n", err)
	}
}

// ================================================
// ЗАВДАННЯ 11: Максимум і мінімум з файлу
// ================================================

// findMaxMin знаходить максимальне і мінімальне значення у файлі з числами
func findMaxMin(inputPath, outputPath string) (int, int, error) {
	nums, err := readNumbers(inputPath)
	if err != nil {
		return 0, 0, fmt.Errorf("findMaxMin: %w", err)
	}

	maxVal, minVal := nums[0], nums[0]
	for _, n := range nums[1:] {
		if n > maxVal {
			maxVal = n
		}
		if n < minVal {
			minVal = n
		}
	}

	content := fmt.Sprintf("Максимум: %d\nМінімум: %d", maxVal, minVal)
	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return 0, 0, fmt.Errorf("findMaxMin: помилка запису: %w", err)
	}
	return maxVal, minVal, nil
}

func Task11() {
	fmt.Println("---> ЗАВДАННЯ №11 (Максимум і мінімум з файлу)")

	os.WriteFile("numbers.txt", []byte("5 3 8 1 9 2 7\n4 6 10 -3 15"), 0644)
	defer func() {
		os.Remove("numbers.txt")
		os.Remove("maxmin_output.txt")
	}()

	maxVal, minVal, err := findMaxMin("numbers.txt", "maxmin_output.txt")
	if err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
		return
	}
	fmt.Printf("[OK] Максимум: %d, Мінімум: %d\n", maxVal, minVal)
	fmt.Println("[OK] Результат записано у maxmin_output.txt")

	// Тест з порожнім файлом
	fmt.Println("\n[ТЕСТ] Порожній файл:")
	os.WriteFile("empty_numbers.txt", []byte(""), 0644)
	defer os.Remove("empty_numbers.txt")
	if _, _, err := findMaxMin("empty_numbers.txt", "out.txt"); err != nil {
		fmt.Printf("[ОЧІКУВАНА ПОМИЛКА] %v\n", err)
	}

	// Тест з неіснуючим файлом
	fmt.Println("\n[ТЕСТ] Неіснуючий файл:")
	if _, _, err := findMaxMin("missing.txt", "out.txt"); err != nil {
		fmt.Printf("[ОЧІКУВАНА ПОМИЛКА] %v\n", err)
	}
}

// ================================================
// ЗАВДАННЯ 13: Унікальні слова з кількох файлів
// ================================================

// mergeUniqueWords збирає унікальні слова з кількох файлів,
// записує кожне на новому рядку у вихідний файл
func mergeUniqueWords(inputPaths []string, outputPath string) error {
	if len(inputPaths) == 0 {
		return fmt.Errorf("mergeUniqueWords: список файлів порожній")
	}

	uniqueWords := make(map[string]bool)
	for _, path := range inputPaths {
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("mergeUniqueWords: помилка читання %q: %w", path, err)
		}
		for _, word := range strings.Fields(string(data)) {
			// Очищуємо пунктуацію і приводимо до нижнього регістру
			clean := strings.ToLower(strings.Trim(word, ".,!?;:\"'()"))
			if clean != "" {
				uniqueWords[clean] = true
			}
		}
	}

	words := make([]string, 0, len(uniqueWords))
	for w := range uniqueWords {
		words = append(words, w)
	}
	sort.Strings(words)

	fmt.Printf("[Унікальних слів знайдено: %d]\n", len(words))
	for _, w := range words {
		fmt.Printf("  %s\n", w)
	}

	return os.WriteFile(outputPath, []byte(strings.Join(words, "\n")), 0644)
}

func Task13() {
	fmt.Println("---> ЗАВДАННЯ №13 (Унікальні слова з кількох файлів)")

	os.WriteFile("words1.txt", []byte("hello world foo bar"), 0644)
	os.WriteFile("words2.txt", []byte("world baz hello qux"), 0644)
	os.WriteFile("words3.txt", []byte("foo new word test"), 0644)
	defer func() {
		os.Remove("words1.txt")
		os.Remove("words2.txt")
		os.Remove("words3.txt")
		os.Remove("unique_output.txt")
	}()

	err := mergeUniqueWords(
		[]string{"words1.txt", "words2.txt", "words3.txt"},
		"unique_output.txt",
	)
	if err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
		return
	}
	fmt.Println("[OK] Унікальні слова записано у unique_output.txt")

	// Тест з неіснуючим файлом у списку
	fmt.Println("\n[ТЕСТ] Неіснуючий файл у списку:")
	if err := mergeUniqueWords([]string{"words1.txt", "missing.txt"}, "out.txt"); err != nil {
		fmt.Printf("[ОЧІКУВАНА ПОМИЛКА] %v\n", err)
	}

	// Тест з порожнім списком файлів
	fmt.Println("\n[ТЕСТ] Порожній список файлів:")
	if err := mergeUniqueWords([]string{}, "out.txt"); err != nil {
		fmt.Printf("[ОЧІКУВАНА ПОМИЛКА] %v\n", err)
	}
}

// ================================================
// ЗАВДАННЯ 15: Сума парних чисел
// ================================================

// sumEvenNumbers рахує суму тільки парних чисел з файлу
func sumEvenNumbers(inputPath, outputPath string) (int, error) {
	nums, err := readNumbers(inputPath)
	if err != nil {
		return 0, fmt.Errorf("sumEvenNumbers: %w", err)
	}

	sum := 0
	count := 0
	for _, n := range nums {
		if n%2 == 0 {
			sum += n
			count++
		}
	}

	content := fmt.Sprintf("Сума парних чисел: %d (парних знайдено: %d)", sum, count)
	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return 0, fmt.Errorf("sumEvenNumbers: помилка запису: %w", err)
	}
	return sum, nil
}

func Task15() {
	fmt.Println("---> ЗАВДАННЯ №15 (Сума парних чисел)")

	os.WriteFile("even_numbers.txt", []byte("1 2 3 4 5 6 7 8 9 10\n11 12 -4 7 -6"), 0644)
	defer func() {
		os.Remove("even_numbers.txt")
		os.Remove("even_output.txt")
	}()

	sum, err := sumEvenNumbers("even_numbers.txt", "even_output.txt")
	if err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
		return
	}
	fmt.Printf("[OK] Сума парних чисел: %d\n", sum)
	fmt.Println("[OK] Результат записано у even_output.txt")

	// Тест — тільки непарні числа
	fmt.Println("\n[ТЕСТ] Файл тільки з непарними числами:")
	os.WriteFile("odd_only.txt", []byte("1 3 5 7 9 11"), 0644)
	defer os.Remove("odd_only.txt")
	defer os.Remove("even_output2.txt")
	if sum2, err := sumEvenNumbers("odd_only.txt", "even_output2.txt"); err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
	} else {
		fmt.Printf("[OK] Сума парних: %d (очікується 0)\n", sum2)
	}

	// Тест з неіснуючим файлом
	fmt.Println("\n[ТЕСТ] Неіснуючий файл:")
	if _, err := sumEvenNumbers("missing.txt", "out.txt"); err != nil {
		fmt.Printf("[ОЧІКУВАНА ПОМИЛКА] %v\n", err)
	}
}

// ================================================
// ЗАВДАННЯ 16: JSON — Pizzeria
// ================================================

type Pizza struct {
	Name                string   `json:"name"`
	Price               float64  `json:"price"`
	Ingredients         []string `json:"ingredients"`
	IsVegetarian        bool     `json:"is_vegetarian"`
	IsSpicy             bool     `json:"is_spicy"`
	SpecialInstructions *string  `json:"special_instructions"` // вказівник бо може бути null
}

type PizzeriaData struct {
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Menu     []Pizza `json:"menu"`
}

type PizzeriaFile struct {
	Pizzeria PizzeriaData `json:"pizzeria"`
}

// loadPizzeria завантажує дані піцерії з JSON-файлу
func loadPizzeria(path string) (*PizzeriaFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("loadPizzeria: %w", err)
	}

	var pf PizzeriaFile
	if err := json.Unmarshal(data, &pf); err != nil {
		return nil, fmt.Errorf("loadPizzeria: помилка розбору JSON: %w", err)
	}
	return &pf, nil
}

// savePizzeria зберігає дані піцерії у JSON-файл з відступами
func savePizzeria(path string, pf *PizzeriaFile) error {
	data, err := json.MarshalIndent(pf, "", "  ")
	if err != nil {
		return fmt.Errorf("savePizzeria: %w", err)
	}
	return os.WriteFile(path, data, 0644)
}

func Task16() {
	fmt.Println("---> ЗАВДАННЯ №16 (JSON — Pizzeria)")

	inputJSON := `{
  "pizzeria": {
    "name": "Pizza World",
    "location": "789 Pine Street",
    "menu": [
      {
        "name": "Margherita",
        "price": 10.99,
        "ingredients": ["tomato", "mozzarella", "basil"],
        "is_vegetarian": true,
        "is_spicy": false,
        "special_instructions": "Extra cheese available upon request"
      },
      {
        "name": "Pepperoni",
        "price": 12.99,
        "ingredients": ["tomato", "mozzarella", "pepperoni"],
        "is_vegetarian": false,
        "is_spicy": true,
        "special_instructions": null
      }
    ]
  }
}`

	os.WriteFile("pizzeria.json", []byte(inputJSON), 0644)
	defer func() {
		os.Remove("pizzeria.json")
		os.Remove("pizzeria_updated.json")
	}()

	// Завантажуємо
	pf, err := loadPizzeria("pizzeria.json")
	if err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
		return
	}

	fmt.Printf("[OK] Завантажено: %s (%s)\n", pf.Pizzeria.Name, pf.Pizzeria.Location)
	fmt.Println("[Меню:]")
	for _, p := range pf.Pizzeria.Menu {
		fmt.Printf("  - %-15s | $%.2f | вегетаріанська: %v | гостра: %v\n",
			p.Name, p.Price, p.IsVegetarian, p.IsSpicy)
	}

	// Модифікуємо дані
	fmt.Println("\n[ЗМІНИ] Оновлюємо дані...")
	pf.Pizzeria.Name = "Pizza World Premium"
	pf.Pizzeria.Menu[0].Price = 13.99
	pf.Pizzeria.Menu[1].Price = 14.99

	newInstruction := "Gluten-free crust available"
	pf.Pizzeria.Menu = append(pf.Pizzeria.Menu, Pizza{
		Name:                "Quattro Formaggi",
		Price:               15.99,
		Ingredients:         []string{"tomato", "mozzarella", "gorgonzola", "parmesan"},
		IsVegetarian:        true,
		IsSpicy:             false,
		SpecialInstructions: &newInstruction,
	})

	// Зберігаємо
	if err := savePizzeria("pizzeria_updated.json", pf); err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
		return
	}

	fmt.Printf("[OK] Нова назва: %s\n", pf.Pizzeria.Name)
	fmt.Println("[Оновлене меню:]")
	for _, p := range pf.Pizzeria.Menu {
		fmt.Printf("  - %-20s | $%.2f\n", p.Name, p.Price)
	}
	fmt.Println("[OK] Збережено у pizzeria_updated.json")

	// Тест з невалідним JSON
	fmt.Println("\n[ТЕСТ] Невалідний JSON:")
	os.WriteFile("bad.json", []byte("{invalid json content}"), 0644)
	defer os.Remove("bad.json")
	if _, err := loadPizzeria("bad.json"); err != nil {
		fmt.Printf("[ОЧІКУВАНА ПОМИЛКА] %v\n", err)
	}

	// Тест з неіснуючим файлом
	fmt.Println("\n[ТЕСТ] Неіснуючий файл:")
	if _, err := loadPizzeria("missing.json"); err != nil {
		fmt.Printf("[ОЧІКУВАНА ПОМИЛКА] %v\n", err)
	}
}
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println("       ЛАБОРАТОРНА РОБОТА №2 | ВАРІАНТ 11         ")
	fmt.Println("==================================================")

	Task5()
	divider()

	Task9()
	divider()

	Task11()
	divider()

	Task12()
	divider()

	Task17()
	divider()

	Task21()
	divider()

	Task27()
	divider()

	Task31()
	divider()

	Task35()
	divider()

	Task39()

	fmt.Println("\n==================================================")
	fmt.Println("               ВИКОНАННЯ ЗАВЕРШЕНО                ")
	fmt.Println("==================================================")
}

func divider() {
	fmt.Println("\n--------------------------------------------------")
}

// Завдання 5: Розрахунок ПДФО (2 версії коду)
func Task5() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №5 (Умови)")
	salary := 60000.0
	fmt.Printf("Вхідні дані: Оклад = %.2f\n", salary)

	// Версія 1: Використання стандартного if-else
	var tax1 float64
	if salary > 50000.0 {
		tax1 = salary * 0.13
	} else {
		tax1 = salary * 0.06
	}

	// Версія 2: Використання switch без виразу (аналог ланцюга if-else)
	var tax2 float64
	switch {
	case salary > 50000.0:
		tax2 = salary * 0.13
	default:
		tax2 = salary * 0.06
	}

	fmt.Printf("Результат (Версія if-else): Податок = %.2f\n", tax1)
	fmt.Printf("Результат (Версія switch): Податок = %.2f\n", tax2)
}

// Завдання 9: Визначення чверті системи координат
func Task9() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №9 (Умови)")
	x, y := -4.5, 3.2
	fmt.Printf("Вхідні дані: x = %.1f, y = %.1f\n", x, y)

	fmt.Print("Результат: Точка знаходиться у ")
	if x > 0 && y > 0 {
		fmt.Println("1-й чверті")
	} else if x < 0 && y > 0 {
		fmt.Println("2-й чверті")
	} else if x < 0 && y < 0 {
		fmt.Println("3-й чверті")
	} else if x > 0 && y < 0 {
		fmt.Println("4-й чверті")
	} else {
		fmt.Println("перетині осей (на осі або в центрі)")
	}
}

// Завдання 11: Перевірка належності точки колу (Радіус = 5, Центр = 0,0)
func Task11() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №11 (Умови)")
	x, y := 3.0, 4.0
	radius := 5.0
	fmt.Printf("Вхідні дані: координати (%.1f, %.1f), радіус = %.1f\n", x, y, radius)

	// Рівняння кола: x^2 + y^2 <= R^2
	if (x*x)+(y*y) <= (radius * radius) {
		fmt.Println("Результат: YES")
	} else {
		fmt.Println("Результат: NO")
	}
}

// Завдання 12: Перевірка, чи є введене значення числом
func Task12() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №12 (Умови)")
	input := "123.45" // Імітація вводу з клавіатури як рядка
	fmt.Printf("Вхідні дані: '%s'\n", input)

	// Спроба конвертувати рядок у число
	_, err := strconv.ParseFloat(input, 64)

	if err == nil {
		fmt.Println("Результат: Number")
	} else {
		fmt.Println("Результат: Other")
	}
}

// Завдання 17: Конвертація температур (C <-> F)
func Task17() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №17 (Умови)")
	tempInput := "98.6f"
	fmt.Printf("Вхідні дані: '%s'\n", tempInput)

	tempInput = strings.ToLower(tempInput)
	lastChar := tempInput[len(tempInput)-1:]
	numStr := tempInput[:len(tempInput)-1]

	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		fmt.Println("Результат: Помилка формату числа")
		return
	}

	if lastChar == "c" {
		fahrenheit := (num * 9 / 5) + 32
		fmt.Printf("Результат: %.1fF\n", fahrenheit)
	} else if lastChar == "f" {
		celsius := (num - 32) * 5 / 9
		fmt.Printf("Результат: %.1fC\n", celsius)
	} else {
		fmt.Println("Результат: Невідомий формат вимірювання")
	}
}

// Завдання 21: Добуток елементів зрізу через цикл
func Task21() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №21 (Цикли)")
	slice := []int{2, 3, 4, 5}
	fmt.Printf("Вхідні дані: %v\n", slice)

	product := 1
	for _, val := range slice {
		product *= val
	}
	fmt.Printf("Результат (добуток): %d\n", product)
}

// Завдання 27: Підрахунок входжень кожного символу в рядок
func Task27() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №27 (Цикли)")
	text := "hello world"
	fmt.Printf("Вхідні дані: '%s'\n", text)

	charCount := make(map[rune]int)
	for _, char := range text {
		charCount[char]++
	}

	fmt.Println("Результат:")
	for char, count := range charCount {
		fmt.Printf(" - '%c': %d разів\n", char, count)
	}
}

// Завдання 31: Сума математичного виразу з параметром z
func Task31() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №31 (Математичний цикл)")
	z := 3
	fmt.Printf("Вхідні дані: z = %d\n", z)

	sum := 0.0
	for n := 1; n <= z; n++ {
		nf := float64(n)
		inner := nf + 2.5*nf
		numerator := math.Sqrt(math.Pow(inner, 3))
		sum += numerator / 4.0
	}
	fmt.Printf("Результат (сума): %.4f\n", sum)
}

// Завдання 35: Сума математичного виразу з параметром z
func Task35() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №35 (Математичний цикл)")
	z := 3
	fmt.Printf("Вхідні дані: z = %d\n", z)

	sum := 0.0
	for n := 1; n <= z; n++ {
		nf := float64(n)
		numerator := math.Tan(nf - 2.0*nf)
		denominator := math.Sqrt(10.0 + 0.6*nf)
		sum += numerator / denominator
	}
	fmt.Printf("Результат (сума): %.4f\n", sum)
}

// Завдання 39: Сума математичного виразу з параметром z
func Task39() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №39 (Математичний цикл)")
	z := 3
	fmt.Printf("Вхідні дані: z = %d\n", z)

	sum := 0.0
	for n := 1; n <= z; n++ {
		nf := float64(n)
		numerator := 2.0*math.Pow(nf, 2) - 4.0*nf + 10.0
		denominator := 2.0 * nf
		sum += numerator / denominator
	}
	fmt.Printf("Результат (сума): %.4f\n", sum)
}
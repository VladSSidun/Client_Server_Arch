package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println("       ЛАБОРАТОРНА РОБОТА №1 | ВАРІАНТ 11         ")
	fmt.Println("==================================================")

	Task2()
	divider()

	Task7()
	divider()

	Task8()
	divider()

	Task12()
	divider()

	Task23()
	divider()

	Task26()
	divider()

	Task28()
	divider()

	Task38()
	divider()

	Task40()
	divider()

	Task42()
	divider()

	Task46()
	divider()

	TaskMathVar13()
	divider()

	TaskMathVar16()
	divider()

	TaskMathVar17()
	divider()

	TaskMathVar19()

	fmt.Println("\n==================================================")
	fmt.Println("               ВИКОНАННЯ ЗАВЕРШЕНО                ")
	fmt.Println("==================================================")
}

func divider() {
	fmt.Println("\n--------------------------------------------------")
}

// Завдання 2: Виводить довжину рядка, а також індекс першого та останнього входження букви.
func Task2() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №2 (Рядки)")
	str := "Programming in Go"
	char := "m"

	fmt.Printf("Вхідні дані:\n - Рядок: '%s'\n - Шукана буква: '%s'\n", str, char)
	fmt.Printf("Результат:\n")
	fmt.Printf(" - Довжина рядка (у байтах): %d\n", len(str))
	fmt.Printf(" - Перше входження: %d\n", strings.Index(str, char))
	fmt.Printf(" - Останнє входження: %d\n", strings.LastIndex(str, char))
}

// Завдання 7: Створює новий рядок str3 шляхом додавання str2 в середину str1.
func Task7() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №7 (Рядки)")
	str1 := "Golang lang"
	str2 := " - is awesome - "

	middle := len(str1) / 2
	str3 := str1[:middle] + str2 + str1[middle:]

	fmt.Printf("Вхідні дані:\n - str1: '%s'\n - str2: '%s'\n", str1, str2)
	fmt.Printf("Результат (str3):\n - '%s'\n", str3)
}

// Завдання 8: Формує нову змінну шляхом видалення символів із рядка (з 0 по 3-й).
func Task8() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №8 (Рядки)")
	word := "University"

	result := ""
	if len(word) > 3 {
		result = word[4:]
	} else {
		result = "Помилка: слово занадто коротке"
	}

	fmt.Printf("Вхідне слово: '%s'\n", word)
	fmt.Printf("Результат (без перших 4-х символів): '%s'\n", result)
}

// Завдання 12: Зріз дійсних чисел. Виводить розмір, значення першого та останнього елементів.
func Task12() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №12 (Зрізи)")
	slice := []float64{10.5, -3.2, 8.9, 4.1, 0.0}

	fmt.Printf("Вхідний зріз: %v\n", slice)
	fmt.Printf("Результат:\n")
	fmt.Printf(" - Розмір зрізу: %d\n", len(slice))

	if len(slice) > 0 {
		fmt.Printf(" - Перший елемент: %.2f\n", slice[0])
		fmt.Printf(" - Останній елемент: %.2f\n", slice[len(slice)-1])
	} else {
		fmt.Println(" - Зріз порожній.")
	}
}

// Завдання 23: Заокруглює всі елементи зрізу до найближчого максимального цілого.
func Task23() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №23 (Зрізи)")
	floats := []float64{1.2, 2.8, -3.5, 4.0}

	ints := make([]int, len(floats))

	for i, v := range floats {
		ints[i] = int(math.Ceil(v))
	}

	fmt.Printf("Вхідний зріз дійсних чисел: %v\n", floats)
	fmt.Printf("Результат (новий зріз цілих чисел): %v\n", ints)
}

// Завдання 26: Видаляє із зрізу всі елементи, які дорівнюють А, і знаходить суму решти.
func Task26() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №26 (Зрізи)")
	slice := []int{5, 10, 15, 5, 20}
	A := 5
	var newSlice []int
	sum := 0

	for _, v := range slice {
		if v != A {
			newSlice = append(newSlice, v)
			sum += v
		}
	}

	fmt.Printf("Вхідні дані:\n - Зріз: %v\n - Число А: %d\n", slice, A)
	fmt.Printf("Результат:\n - Зріз після видалення: %v\n - Сума решти: %d\n", newSlice, sum)
}

// Завдання 28: Видаляє із зрізу елементи, які лежать в діапазоні від індексу 'a' по 'b'.
func Task28() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №28 (Зрізи)")
	slice := []string{"A", "B", "C", "D", "E", "F"}
	a, b := 1, 3

	fmt.Printf("Вхідні дані:\n - Зріз: %v\n - Видалити від індексу %d по %d\n", slice, a, b)

	if a >= 0 && b < len(slice) && a <= b {
		newSlice := make([]string, 0, len(slice)-(b-a+1))
		newSlice = append(newSlice, slice[:a]...)
		newSlice = append(newSlice, slice[b+1:]...)
		
		fmt.Printf("Результат:\n - Новий зріз: %v\n", newSlice)
	} else {
		fmt.Println("Результат:\n - Помилка: Некоректні індекси.")
	}
}

// Завдання 38: Об'єднує два об'єкти map і виводить результат та кількість елементів.
func Task38() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №38 (Карти/Maps)")
	map1 := map[int]int{1: 100, 2: 200}
	map2 := map[int]int{3: 300, 2: 999}

	merged := make(map[int]int)
	for k, v := range map1 {
		merged[k] = v
	}
	for k, v := range map2 {
		merged[k] = v
	}

	fmt.Printf("Вхідні дані:\n - Map 1: %v\n - Map 2: %v\n", map1, map2)
	fmt.Printf("Результат:\n - Об'єднана Map: %v\n - Кількість елементів: %d\n", merged, len(merged))
}

// Завдання 40: Видаляє всі елементи, значення яких закінчуються підрядком А.
func Task40() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №40 (Карти/Maps)")
	myMap := map[int]string{1: "apple_tree", 2: "banana", 3: "red_apple"}
	A := "apple"

	fmt.Printf("Вхідні дані:\n - Map: %v\n - Підрядок А: '%s'\n", myMap, A)

	for k, v := range myMap {
		if strings.HasSuffix(v, A) {
			delete(myMap, k)
		}
	}

	fmt.Printf("Результат:\n - Map після очищення: %v\n", myMap)
}

// Завдання 42: Видаляє елементи, значення і ключі яких кратні А, та знаходить добуток решти.
func Task42() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №42 (Карти/Maps)")
	myMap := map[int]int{2: 4, 3: 9, 4: 8, 5: 10}
	A := 2
	product := 1
	hasElements := false

	fmt.Printf("Вхідні дані:\n - Map: %v\n - Число А: %d\n", myMap, A)

	for k, v := range myMap {
		if k%A == 0 && v%A == 0 {
			delete(myMap, k)
		} else {
			product *= k * v
			hasElements = true
		}
	}

	if !hasElements {
		product = 0
	}
	fmt.Printf("Результат:\n - Map після видалення кратних: %v\n", myMap)
	fmt.Printf(" - Добуток ключів та значень, що залишилися: %d\n", product)
}

// Завдання 46: Розраховує результат виразу: (val1 * 3 + val1) / 4 * val2.
func Task46() {
	fmt.Println("---> ВИКОНАННЯ ЗАВДАННЯ №46 (Арифметика)")
	val1, val2 := 4.0, 2.0

	result := ((val1*3 + val1) / 4.0) * val2

	fmt.Printf("Вхідні дані: val1 = %.1f, val2 = %.1f\n", val1, val2)
	fmt.Printf("Результат виразу: %.2f\n", result)
}

// TaskMathVar13: Варіант 13
func TaskMathVar13() {
	fmt.Println("---> ВИКОНАННЯ: Варіант 13")
	n := 2.0 

	numerator := 5.0 * n * math.Cos(n)
	denominator := math.Sqrt(math.Pow(n, 3))
	result := numerator / denominator

	fmt.Printf("Вхідне число n: %.2f\n", n)
	fmt.Printf("Результат формули: %.4f\n", result)
}

// TaskMathVar16: Варіант 16
func TaskMathVar16() {
	fmt.Println("---> ВИКОНАННЯ: Варіант 16")
	n := 2.0

	numerator := 3.0*math.Sin(n) - 15.0
	denominator := math.Sqrt(math.Pow(n, 5))
	result := numerator / denominator

	fmt.Printf("Вхідне число n: %.2f\n", n)
	fmt.Printf("Результат формули: %.4f\n", result)
}

// TaskMathVar17: Варіант 17
func TaskMathVar17() {
	fmt.Println("---> ВИКОНАННЯ: Варіант 17")
	n := 2.0

	numerator := 10.0 + 2.0*math.Cos(n)
	denominator := 5.0 - math.Sqrt(math.Pow(n, 5))
	result := numerator / denominator

	fmt.Printf("Вхідне число n: %.2f\n", n)
	fmt.Printf("Результат формули: %.4f\n", result)
}

// TaskMathVar19: Варіант 19
func TaskMathVar19() {
	fmt.Println("---> ВИКОНАННЯ: Варіант 19")
	n := 2.0

	numerator := 2.0*math.Pow(n, 2) - 4.0*n + 10.0
	denominator := 2.0 * n
	result := numerator / denominator

	fmt.Printf("Вхідне число n: %.2f\n", n)
	fmt.Printf("Результат формули: %.4f\n", result)
}
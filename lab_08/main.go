package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println("       ЛАБОРАТОРНА РОБОТА №8 | ВАРІАНТ 11(2, 7, 8)       ")
	fmt.Println("==================================================")

	Task2()
	divider()
	Task7()
	divider()
	Task8()

	fmt.Println("\n==================================================")
	fmt.Println("               ВИКОНАННЯ ЗАВЕРШЕНО                ")
	fmt.Println("==================================================")
}

func divider() {
	fmt.Println("\n--------------------------------------------------")
}

// ================================================
// ЗАВДАННЯ 2: Множення матриць через горутини
// Кожен рядок результуючої матриці — окрема горутина
// ================================================

// multiplyMatrices множить матриці A (n×m) на B (m×k)
// кожен рядок результату обчислюється в окремій горутині
func multiplyMatrices(A, B [][]int) ([][]int, error) {
	rowsA := len(A)
	if rowsA == 0 {
		return nil, fmt.Errorf("multiplyMatrices: матриця A порожня")
	}
	colsA := len(A[0])
	rowsB := len(B)
	if rowsB == 0 {
		return nil, fmt.Errorf("multiplyMatrices: матриця B порожня")
	}
	colsB := len(B[0])

	// Перевірка сумісності — кількість стовпців A = кількість рядків B
	if colsA != rowsB {
		return nil, fmt.Errorf("multiplyMatrices: несумісні розміри (%dx%d) * (%dx%d)",
			rowsA, colsA, rowsB, colsB)
	}

	// Створюємо результуючу матрицю rowsA x colsB
	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	// WaitGroup — чекаємо поки всі горутини завершаться
	var wg sync.WaitGroup

	// Запускаємо горутину для кожного рядка результату
	for i := 0; i < rowsA; i++ {
		wg.Add(1) // повідомляємо WaitGroup що є ще одна горутина

		go func(row int) {
			defer wg.Done() // коли горутина завершиться — зменшуємо лічильник

			// Обчислюємо всі елементи рядка row
			for j := 0; j < colsB; j++ {
				sum := 0
				for k := 0; k < colsA; k++ {
					sum += A[row][k] * B[k][j]
				}
				result[row][j] = sum
			}
		}(i) // передаємо i як аргумент — інакше всі горутини захоплять одне і те ж i
	}

	wg.Wait() // чекаємо поки всі горутини викличуть Done()
	return result, nil
}

// printMatrix виводить матрицю у зручному форматі
func printMatrix(name string, m [][]int) {
	fmt.Printf("%s:\n", name)
	for _, row := range m {
		fmt.Print("  [")
		for j, val := range row {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%4d", val)
		}
		fmt.Println("]")
	}
}

func Task2() {
	fmt.Println("---> ЗАВДАННЯ №2 (Множення матриць через горутини)")

	// Тест 1: стандартне множення 3x3 * 3x3
	fmt.Println("\n[ТЕСТ 1] Множення 3x3 * 3x3:")
	A := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	B := [][]int{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	}
	printMatrix("Матриця A", A)
	printMatrix("Матриця B", B)

	result, err := multiplyMatrices(A, B)
	if err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
	} else {
		printMatrix("Результат A*B", result)
	}

	// Тест 2: множення 2x3 * 3x2
	fmt.Println("\n[ТЕСТ 2] Множення 2x3 * 3x2:")
	C := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	D := [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}
	printMatrix("Матриця C", C)
	printMatrix("Матриця D", D)

	result2, err := multiplyMatrices(C, D)
	if err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
	} else {
		printMatrix("Результат C*D", result2)
	}

	// Тест 3 (виключення): несумісні розміри
	fmt.Println("\n[ТЕСТ 3] Несумісні розміри (2x2 * 3x2):")
	E := [][]int{{1, 2}, {3, 4}}
	F := [][]int{{1, 2}, {3, 4}, {5, 6}}
	if _, err := multiplyMatrices(E, F); err != nil {
		fmt.Printf("[ОЧІКУВАНА ПОМИЛКА] %v\n", err)
	}

	// Тест 4 (виключення): порожня матриця
	fmt.Println("\n[ТЕСТ 4] Порожня матриця:")
	if _, err := multiplyMatrices([][]int{}, B); err != nil {
		fmt.Printf("[ОЧІКУВАНА ПОМИЛКА] %v\n", err)
	}
}

// ================================================
// ЗАВДАННЯ 7: Сума середніх арифметичних трьох зрізів
// Кожен середній обчислюється в окремій горутині
// ================================================

// averageWorker обчислює середнє арифметичне зрізу
// і відправляє результат у канал
func averageWorker(id int, slice []float64, ch chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()

	if len(slice) == 0 {
		fmt.Printf("[Горутина %d] порожній зріз — відправляємо 0\n", id)
		ch <- 0
		return
	}

	sum := 0.0
	for _, v := range slice {
		sum += v
	}
	avg := sum / float64(len(slice))
	fmt.Printf("[Горутина %d] зріз %v → середнє = %.2f\n", id, slice, avg)
	ch <- avg
}

// sumOfAverages рахує суму середніх арифметичних трьох зрізів
func sumOfAverages(slices [][]float64) (float64, error) {
	if len(slices) != 3 {
		return 0, fmt.Errorf("sumOfAverages: очікується 3 зрізи, отримано %d", len(slices))
	}

	// Буферизований канал на 3 місця — горутини не блокуватимуться при відправці
	ch := make(chan float64, 3)
	var wg sync.WaitGroup

	// Запускаємо 3 горутини паралельно
	for i, slice := range slices {
		wg.Add(1)
		go averageWorker(i+1, slice, ch, &wg)
	}

	// Чекаємо всіх горутин і закриваємо канал
	wg.Wait()
	close(ch)

	// Збираємо результати з каналу
	total := 0.0
	for avg := range ch {
		total += avg
	}
	return total, nil
}

func Task7() {
	fmt.Println("---> ЗАВДАННЯ №7 (Сума середніх арифметичних через горутини)")

	// Тест 1: стандартний випадок
	fmt.Println("\n[ТЕСТ 1] Три зрізи з числами:")
	slices := [][]float64{
		{1.0, 2.0, 3.0, 4.0, 5.0},   // середнє = 3.0
		{10.0, 20.0, 30.0},           // середнє = 20.0
		{2.5, 7.5, 5.0, 10.0, 15.0}, // середнє = 8.0
	}
	total, err := sumOfAverages(slices)
	if err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
	} else {
		fmt.Printf("[OK] Сума середніх: %.2f (очікується 31.00)\n", total)
	}

	// Тест 2: зріз з одного елемента
	fmt.Println("\n[ТЕСТ 2] Зрізи з одного елемента:")
	slices2 := [][]float64{
		{5.0},
		{10.0},
		{15.0},
	}
	total2, err := sumOfAverages(slices2)
	if err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
	} else {
		fmt.Printf("[OK] Сума середніх: %.2f (очікується 30.00)\n", total2)
	}

	// Тест 3 (виключення): неправильна кількість зрізів
	fmt.Println("\n[ТЕСТ 3] Неправильна кількість зрізів:")
	_, err = sumOfAverages([][]float64{{1, 2}, {3, 4}})
	if err != nil {
		fmt.Printf("[ОЧІКУВАНА ПОМИЛКА] %v\n", err)
	}
}

// ================================================
// ЗАВДАННЯ 8: Пошук значення у зрізі через 3 горутини
// Горутини повністю перекривають довжину зрізу
// ================================================

// SearchResult містить результат пошуку від однієї горутини
type SearchResult struct {
	Found bool
	Index int
	Value int
}

// searchWorker шукає target у підзрізі slice[start:end]
// відправляє результат у канал
func searchWorker(id int, slice []int, start, end, target int, ch chan<- SearchResult, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("[Горутина %d] шукаємо %d у діапазоні [%d:%d]\n", id, target, start, end)

	for i := start; i < end; i++ {
		if slice[i] == target {
			fmt.Printf("[Горутина %d] знайдено на індексі %d\n", id, i)
			ch <- SearchResult{Found: true, Index: i, Value: slice[i]}
			return
		}
	}

	fmt.Printf("[Горутина %d] не знайдено\n", id)
	ch <- SearchResult{Found: false}
}

// searchInSlice шукає target у зрізі через 3 горутини
func searchInSlice(slice []int, target int) (SearchResult, error) {
	if len(slice) == 0 {
		return SearchResult{}, fmt.Errorf("searchInSlice: зріз порожній")
	}

	n := len(slice)
	numWorkers := 3

	// Рівномірно ділимо зріз на 3 частини
	chunkSize := (n + numWorkers - 1) / numWorkers

	ch := make(chan SearchResult, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > n {
			end = n // останній воркер не виходить за межі
		}
		if start >= n {
			break // якщо зріз менший за 3 елементи
		}

		wg.Add(1)
		go searchWorker(i+1, slice, start, end, target, ch, &wg)
	}

	wg.Wait()
	close(ch)

	// Перевіряємо результати всіх горутин
	for res := range ch {
		if res.Found {
			return res, nil
		}
	}
	return SearchResult{Found: false}, nil
}

func Task8() {
	fmt.Println("---> ЗАВДАННЯ №8 (Пошук значення через 3 горутини)")

	slice := []int{15, 3, 42, 8, 77, 19, 56, 4, 91, 33, 67, 25}
	fmt.Printf("[Зріз] %v\n", slice)

	// Тест 1: елемент існує в зрізі
	fmt.Println("\n[ТЕСТ 1] Шукаємо 56:")
	res, err := searchInSlice(slice, 56)
	if err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
	} else if res.Found {
		fmt.Printf("[OK] Знайдено значення %d на індексі %d\n", res.Value, res.Index)
	} else {
		fmt.Println("[OK] Не знайдено")
	}

	// Тест 2: елемент на початку зрізу
	fmt.Println("\n[ТЕСТ 2] Шукаємо 15 (перший елемент):")
	res2, err := searchInSlice(slice, 15)
	if err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
	} else if res2.Found {
		fmt.Printf("[OK] Знайдено значення %d на індексі %d\n", res2.Value, res2.Index)
	} else {
		fmt.Println("[OK] Не знайдено")
	}

	// Тест 3: елемент не існує
	fmt.Println("\n[ТЕСТ 3] Шукаємо 100 (відсутнє):")
	res3, err := searchInSlice(slice, 100)
	if err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
	} else if res3.Found {
		fmt.Printf("[OK] Знайдено значення %d на індексі %d\n", res3.Value, res3.Index)
	} else {
		fmt.Println("[OK] Значення 100 не знайдено")
	}

	// Тест 4 (виключення): порожній зріз
	fmt.Println("\n[ТЕСТ 4] Порожній зріз:")
	_, err = searchInSlice([]int{}, 42)
	if err != nil {
		fmt.Printf("[ОЧІКУВАНА ПОМИЛКА] %v\n", err)
	}
}
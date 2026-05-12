package main

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println("       ЛАБОРАТОРНА РОБОТА №6 | ВАРІАНТ 11         ")
	fmt.Println("==================================================")

	Task11()
	divider()
	Task23()
	divider()
	Task31()
	divider()
	Task41()
	divider()
	Task56()

	fmt.Println("\n==================================================")
	fmt.Println("               ВИКОНАННЯ ЗАВЕРШЕНО                ")
	fmt.Println("==================================================")
}

func divider() {
	fmt.Println("\n--------------------------------------------------")
}

// ================================================
// ЗАВДАННЯ 11: Власний тип ValidationError
// ================================================

// ValidationError — власний тип помилки з полями Field та Message
type ValidationError struct {
	Field   string
	Message string
}

// Error() реалізує інтерфейс error
func (e *ValidationError) Error() string {
	return fmt.Sprintf("помилка валідації: поле %q — %s", e.Field, e.Message)
}

// User — структура користувача для валідації
type User struct {
	Name  string
	Age   int
	Email string
}

// validateUser перевіряє дані користувача і повертає ValidationError при помилці
func validateUser(u User) error {
	if strings.TrimSpace(u.Name) == "" {
		return &ValidationError{
			Field:   "name",
			Message: "ім'я не може бути порожнім",
		}
	}
	if u.Age < 0 || u.Age > 150 {
		return &ValidationError{
			Field:   "age",
			Message: fmt.Sprintf("вік %d виходить за межі діапазону [0, 150]", u.Age),
		}
	}
	if !strings.Contains(u.Email, "@") {
		return &ValidationError{
			Field:   "email",
			Message: fmt.Sprintf("значення %q не є коректним email", u.Email),
		}
	}
	return nil
}

func Task11() {
	fmt.Println("---> ЗАВДАННЯ №11 (Власний тип ValidationError)")

	// Тест 1: коректний користувач
	validUser := User{Name: "Влад", Age: 20, Email: "vlad@gmail.com"}
	if err := validateUser(validUser); err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
	} else {
		fmt.Printf("[OK] Користувач '%s' пройшов валідацію\n", validUser.Name)
	}

	// Тест 2: порожнє ім'я
	noName := User{Name: "", Age: 20, Email: "vlad@gmail.com"}
	if err := validateUser(noName); err != nil {
		// errors.As — витягуємо конкретний тип для доступу до полів
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("[ПОМИЛКА] Поле: %q | Причина: %s\n", ve.Field, ve.Message)
		}
	}

	// Тест 3: некоректний вік
	badAge := User{Name: "Влад", Age: -5, Email: "vlad@gmail.com"}
	if err := validateUser(badAge); err != nil {
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("[ПОМИЛКА] Поле: %q | Причина: %s\n", ve.Field, ve.Message)
		}
	}

	// Тест 4: некоректний email
	badEmail := User{Name: "Влад", Age: 20, Email: "not-an-email"}
	if err := validateUser(badEmail); err != nil {
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("[ПОМИЛКА] Поле: %q | Причина: %s\n", ve.Field, ve.Message)
		}
	}
}

// ================================================
// ЗАВДАННЯ 23: Загортання *os.PathError + errors.As
// ================================================

// readConfig зчитує файл конфігурації і загортає помилку os у власний контекст
func readConfig(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		// %w — загортаємо оригінальну помилку, зберігаючи її для errors.As
		return "", fmt.Errorf("readConfig: не вдалося прочитати файл %q: %w", path, err)
	}
	return string(data), nil
}

func Task23() {
	fmt.Println("---> ЗАВДАННЯ №23 (Загортання *os.PathError + errors.As)")

	// Тест 1: файл не існує — отримаємо загорнуту *os.PathError
	_, err := readConfig("non_existent_config.json")
	if err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)

		// errors.As рекурсивно проходить ланцюжок загортань
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			fmt.Printf("[ДЕТАЛІ] Операція: %q | Шлях: %q | Причина: %v\n",
				pathErr.Op, pathErr.Path, pathErr.Err)
		}
	}

	// Тест 2: створюємо тимчасовий файл і читаємо успішно
	tmpFile := "temp_config_test.txt"
	os.WriteFile(tmpFile, []byte("config data"), 0644)
	defer os.Remove(tmpFile)

	content, err := readConfig(tmpFile)
	if err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
	} else {
		fmt.Printf("[OK] Файл прочитано успішно. Вміст: %q\n", content)
	}

	// Тест 3: демонстрація errors.Unwrap для ручного розгортання
	_, err = readConfig("another_missing.txt")
	if err != nil {
		fmt.Printf("[UNWRAP] Загорнута помилка: %v\n", errors.Unwrap(err))
	}
}

// ================================================
// ЗАВДАННЯ 31: panic + recover через defer
// ================================================

// riskyOperation викликає panic з рядковим повідомленням
func riskyOperation(input string) {
	if input == "" {
		panic("riskyOperation: отримано порожній рядок — це неприпустимо")
	}
	fmt.Printf("[OK] Операція виконана успішно з вхідними даними: %q\n", input)
}

// safeOperation відновлює виконання після паніки через defer + recover
func safeOperation(input string) (err error) {
	// defer з recover — ЄДИНИЙ спосіб перехопити паніку
	defer func() {
		if r := recover(); r != nil {
			// Конвертуємо паніку у звичайну помилку
			err = fmt.Errorf("перехоплено паніку: %v", r)
		}
	}()

	riskyOperation(input)
	return nil
}

func Task31() {
	fmt.Println("---> ЗАВДАННЯ №31 (panic + recover через defer)")

	// Тест 1: нормальне виконання — паніки немає
	if err := safeOperation("hello"); err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
	}

	// Тест 2: порожній рядок — спричинить паніку, яку перехопить recover
	if err := safeOperation(""); err != nil {
		fmt.Printf("[ВІДНОВЛЕНО] %v\n", err)
	}

	// Тест 3: демонстрація порядку виконання defer при паніці
	fmt.Println("[DEFER] Демонстрація порядку LIFO при паніці:")
	func() {
		defer fmt.Println("[DEFER] Виконано 1-й defer (останній зареєстрований)")
		defer fmt.Println("[DEFER] Виконано 2-й defer")
		defer fmt.Println("[DEFER] Виконано 3-й defer (перший зареєстрований)")
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("[RECOVER] Паніку перехоплено: %v\n", r)
			}
		}()
		panic("тестова паніка для демонстрації LIFO")
	}()

	fmt.Println("[OK] Програма продовжує роботу після паніки")
}

// ================================================
// ЗАВДАННЯ 41: defer для гарантованого закриття файлу
// ================================================

// processFile зчитує файл і обробляє його вміст
// defer гарантує закриття файлу навіть при виникненні помилки
func processFile(path string) (result string, err error) {
	fmt.Printf("[ВІДКРИТТЯ] Спроба відкрити файл: %q\n", path)

	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("processFile: %w", err)
	}

	// defer спрацює при будь-якому завершенні функції:
	// нормальному поверненні, помилці або паніці
	defer func() {
		fmt.Printf("[DEFER] Закриваємо файл: %q\n", path)
		if cerr := f.Close(); cerr != nil && err == nil {
			// Якщо основної помилки не було — повертаємо помилку закриття
			err = fmt.Errorf("processFile: помилка закриття файлу: %w", cerr)
		}
	}()

	// Вимірюємо час обробки через defer
	start := time.Now()
	defer func() {
		fmt.Printf("[DEFER] Час обробки файлу: %v\n", time.Since(start))
	}()

	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	if err != nil {
		return "", fmt.Errorf("processFile: помилка читання: %w", err)
	}

	return string(buf[:n]), nil
}

func Task41() {
	fmt.Println("---> ЗАВДАННЯ №41 (defer для закриття файлу)")

	// Тест 1: файл не існує — defer все одно виконається (але f.Close не буде)
	_, err := processFile("missing_file.txt")
	if err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
	}

	// Тест 2: створюємо реальний файл і читаємо його
	tmpPath := "test_defer_file.txt"
	os.WriteFile(tmpPath, []byte("Привіт з defer!"), 0644)
	defer os.Remove(tmpPath)

	content, err := processFile(tmpPath)
	if err != nil {
		fmt.Printf("[ПОМИЛКА] %v\n", err)
	} else {
		fmt.Printf("[OK] Вміст файлу: %q\n", content)
	}

	// Тест 3: демонстрація defer з вимірюванням часу окремо
	fmt.Println("[DEFER] Демонстрація вимірювання часу:")
	func() {
		start := time.Now()
		defer func() {
			fmt.Printf("[DEFER] Функція виконувалась: %v\n", time.Since(start))
		}()
		time.Sleep(10 * time.Millisecond)
		fmt.Println("[OK] Робота функції завершена")
	}()
}

// ================================================
// ЗАВДАННЯ 56: Дженерик тип Result[T]
// ================================================

// Result — універсальний тип для представлення результату або помилки
type Result[T any] struct {
	value T
	err   error
}

// Ok створює успішний Result
func Ok[T any](v T) Result[T] {
	return Result[T]{value: v}
}

// Err створює Result з помилкою
func Err[T any](e error) Result[T] {
	return Result[T]{err: e}
}

// IsOk повертає true якщо помилки немає
func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// Unwrap повертає значення та помилку
func (r Result[T]) Unwrap() (T, error) {
	return r.value, r.err
}

// MapErr трансформує помилку якщо вона є, значення залишається незмінним
func (r Result[T]) MapErr(fn func(error) error) Result[T] {
	if r.err != nil {
		return Result[T]{err: fn(r.err)}
	}
	return r
}

// String для зручного виводу
func (r Result[T]) String() string {
	if r.IsOk() {
		return fmt.Sprintf("Ok(%v)", r.value)
	}
	return fmt.Sprintf("Err(%v)", r.err)
}

// parsePositiveInt парсить рядок у додатнє число, повертає Result
func parsePositiveInt(s string) Result[int] {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	if err != nil {
		return Err[int](fmt.Errorf("parsePositiveInt: не вдалося розпарсити %q: %w", s, err))
	}
	if n < 0 {
		return Err[int](fmt.Errorf("parsePositiveInt: очікується додатнє число, отримано %d", n))
	}
	return Ok(n)
}

// divide ділить два числа, повертає Result
func divideResult(a, b float64) Result[float64] {
	if b == 0 {
		return Err[float64](errors.New("divide: ділення на нуль"))
	}
	return Ok(a / b)
}

func Task56() {
	fmt.Println("---> ЗАВДАННЯ №56 (Result[T] — функціональний стиль)")

	// Тест 1: успішне парсування
	r1 := parsePositiveInt("42")
	fmt.Printf("[Result] %s\n", r1)
	if val, err := r1.Unwrap(); err == nil {
		fmt.Printf("[OK] Значення: %d\n", val)
	}

	// Тест 2: некоректний рядок
	r2 := parsePositiveInt("abc")
	fmt.Printf("[Result] %s\n", r2)

	// Тест 3: від'ємне число
	r3 := parsePositiveInt("-10")
	fmt.Printf("[Result] %s\n", r3)

	// Тест 4: MapErr — додаємо контекст до помилки
	r4 := parsePositiveInt("xyz").MapErr(func(err error) error {
		return fmt.Errorf("обробка форми: %w", err)
	})
	fmt.Printf("[MapErr] %s\n", r4)

	// Тест 5: ділення через Result
	fmt.Println()
	cases := [][2]float64{{10, 2}, {7, 0}, {100, 4}}
	for _, c := range cases {
		res := divideResult(c[0], c[1])
		if val, err := res.Unwrap(); err == nil {
			fmt.Printf("[OK] %.0f / %.0f = %.2f\n", c[0], c[1], val)
		} else {
			fmt.Printf("[ПОМИЛКА] %.0f / %.0f → %v\n", c[0], c[1], err)
		}
	}

	// Тест 6: Result зі string
	getGreeting := func(name string) Result[string] {
		if name == "" {
			return Err[string](errors.New("ім'я не може бути порожнім"))
		}
		return Ok("Привіт, " + name + "!")
	}

	fmt.Println()
	for _, name := range []string{"Влад", "", "Anna"} {
		r := getGreeting(name)
		fmt.Printf("[Result<string>] %s\n", r)
	}

	// Тест 7: runtime/debug — отримання stack trace при паніці
	fmt.Println("\n[STACK] Демонстрація отримання stack trace:")
	func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				stack := debug.Stack()
				// Виводимо лише перші 3 рядки стеку для читабельності
				lines := strings.Split(string(stack), "\n")
				if len(lines) > 3 {
					lines = lines[:3]
				}
				fmt.Printf("[RECOVER] Паніка: %v\n[STACK]  %s\n", r, strings.Join(lines, "\n         "))
				err = fmt.Errorf("відновлено після паніки: %v", r)
			}
		}()
		panic("демонстраційна паніка зі stack trace")
	}()
}
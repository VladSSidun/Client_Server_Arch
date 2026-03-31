package wallet

import (
	"testing"
)

// Тест 1: Додавання коректних купюр
func TestAddBill_Valid(t *testing.T) {
	w := &Wallet{}

	err := w.AddBill(UAH, 500)
	if err != nil {
		t.Errorf("Очікували nil, отримали: %v", err)
	}

	err = w.AddBill(USD, 20)
	if err != nil {
		t.Errorf("Очікували nil, отримали: %v", err)
	}

	if len(w.Bills) != 2 {
		t.Errorf("Очікували 2 купюри, отримали: %d", len(w.Bills))
	}
}

// Тест 2: Додавання некоректної купюри
func TestAddBill_Invalid(t *testing.T) {
	w := &Wallet{}

	err := w.AddBill(UAH, -100)
	if err == nil {
		t.Error("Очікували помилку для від'ємного номіналу")
	}

	err = w.AddBill(USD, 0)
	if err == nil {
		t.Error("Очікували помилку для нульового номіналу")
	}

	if len(w.Bills) != 0 {
		t.Errorf("Гаманець повинен бути порожнім, але містить: %d купюр", len(w.Bills))
	}
}

// Тест 3: Підрахунок суми по валюті
func TestTotalByCurrency(t *testing.T) {
	w := &Wallet{}
	w.AddBill(UAH, 500)
	w.AddBill(UAH, 200)
	w.AddBill(USD, 10)

	totalUAH := w.TotalByCurrency(UAH)
	if totalUAH != 700 {
		t.Errorf("Очікували 700 UAH, отримали: %.2f", totalUAH)
	}

	totalUSD := w.TotalByCurrency(USD)
	if totalUSD != 10 {
		t.Errorf("Очікували 10 USD, отримали: %.2f", totalUSD)
	}
}

// Тест 4: Конвертація всього у гривні
func TestTotalInCurrency_UAH(t *testing.T) {
	w := &Wallet{}
	w.AddBill(UAH, 500)
	w.AddBill(USD, 10) // 10 * 41 = 410 UAH

	total := w.TotalInCurrency(UAH, ExchangeRate)
	expected := 910.0
	if total != expected {
		t.Errorf("Очікували %.2f UAH, отримали: %.2f", expected, total)
	}
}

// Тест 5: Конвертація всього у долари
func TestTotalInCurrency_USD(t *testing.T) {
	w := &Wallet{}
	w.AddBill(UAH, 410) // 410 / 41 = 10 USD
	w.AddBill(USD, 5)

	total := w.TotalInCurrency(USD, ExchangeRate)
	expected := 15.0
	if total != expected {
		t.Errorf("Очікували %.2f USD, отримали: %.2f", expected, total)
	}
}
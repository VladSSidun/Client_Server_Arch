package order

import (
	"testing"
)

// Тест 1: Створення коректного продукту
func TestNewProduct_Valid(t *testing.T) {
	p, err := NewProduct("Ноутбук", Electronics, 35000)
	if err != nil {
		t.Errorf("Очікували nil, отримали: %v", err)
	}
	if p.Name != "Ноутбук" {
		t.Errorf("Очікували 'Ноутбук', отримали: %s", p.Name)
	}
	if p.Price != 35000 {
		t.Errorf("Очікували 35000, отримали: %.2f", p.Price)
	}
}

// Тест 2: Створення продукту з некоректною ціною
func TestNewProduct_Invalid(t *testing.T) {
	p, err := NewProduct("Зламаний", Electronics, -500)
	if err == nil {
		t.Error("Очікували помилку для від'ємної ціни")
	}
	if p != nil {
		t.Error("Продукт повинен бути nil при некоректній ціні")
	}
}

// Тест 3: Знижка на товар
func TestSetDiscount(t *testing.T) {
	p, _ := NewProduct("Телефон", Electronics, 10000)

	err := p.SetDiscount(10)
	if err != nil {
		t.Errorf("Очікували nil, отримали: %v", err)
	}

	expected := 9000.0
	if p.FinalPrice() != expected {
		t.Errorf("Очікували %.2f, отримали: %.2f", expected, p.FinalPrice())
	}

	// Некоректна знижка
	err = p.SetDiscount(150)
	if err == nil {
		t.Error("Очікували помилку для знижки > 100%")
	}
}

// Тест 4: Загальна вартість замовлення
func TestTotalPrice(t *testing.T) {
	o := NewOrder(1)

	p1, _ := NewProduct("Ноутбук", Electronics, 35000)
	p2, _ := NewProduct("Футболка", Clothing, 800)
	p1.SetDiscount(10) // 35000 - 10% = 31500
	// p2 без знижки = 800

	o.AddProduct(p1)
	o.AddProduct(p2)

	expected := 32300.0
	if o.TotalPrice() != expected {
		t.Errorf("Очікували %.2f, отримали: %.2f", expected, o.TotalPrice())
	}
}

// Тест 5: Вартість за типом товару
func TestTotalByType(t *testing.T) {
	o := NewOrder(2)

	p1, _ := NewProduct("Ноутбук", Electronics, 35000)
	p2, _ := NewProduct("Телефон", Electronics, 15000)
	p3, _ := NewProduct("Хліб", Food, 45)

	o.AddProduct(p1)
	o.AddProduct(p2)
	o.AddProduct(p3)

	expectedElectronics := 50000.0
	if o.TotalByType(Electronics) != expectedElectronics {
		t.Errorf("Очікували %.2f, отримали: %.2f", expectedElectronics, o.TotalByType(Electronics))
	}

	expectedFood := 45.0
	if o.TotalByType(Food) != expectedFood {
		t.Errorf("Очікували %.2f, отримали: %.2f", expectedFood, o.TotalByType(Food))
	}
}
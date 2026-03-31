package main

import (
	"fmt"
	"lab5/order"
	"lab5/wallet"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println("       ЛАБОРАТОРНА РОБОТА №5 | ВАРІАНТ 11         ")
	fmt.Println("==================================================")

	// ---- ЗАВДАННЯ 4 ----
	fmt.Println("\n==================== ЗАВДАННЯ 4 ====================")

	w := &wallet.Wallet{}
	w.AddBill(wallet.UAH, 500)
	w.AddBill(wallet.UAH, 200)
	w.AddBill(wallet.USD, 10)
	w.AddBill(wallet.USD, 20)
	w.Print()

	fmt.Printf("\nВсього UAH: %.2f\n", w.TotalByCurrency(wallet.UAH))
	fmt.Printf("Всього USD: %.2f\n", w.TotalByCurrency(wallet.USD))
	fmt.Printf("Все у гривнях (курс %.0f): %.2f UAH\n", wallet.ExchangeRate, w.TotalInCurrency(wallet.UAH, wallet.ExchangeRate))
	fmt.Printf("Все у доларах (курс %.0f): %.2f USD\n", wallet.ExchangeRate, w.TotalInCurrency(wallet.USD, wallet.ExchangeRate))

	fmt.Println("\n--- Некоректні дані ---")
	if err := w.AddBill(wallet.UAH, -100); err != nil {
		fmt.Println("[Помилка]", err)
	}

	// ---- ЗАВДАННЯ 7 ----
	fmt.Println("\n==================== ЗАВДАННЯ 7 ====================")

	o := order.NewOrder(1)

	laptop, _ := order.NewProduct("Ноутбук", order.Electronics, 35000)
	phone, _ := order.NewProduct("Смартфон", order.Electronics, 15000)
	tshirt, _ := order.NewProduct("Футболка", order.Clothing, 800)
	bread, _ := order.NewProduct("Хліб", order.Food, 45)

	o.AddProduct(laptop)
	o.AddProduct(phone)
	o.AddProduct(tshirt)
	o.AddProduct(bread)
	o.Print()

	laptop.SetDiscount(10)
	phone.SetDiscount(15)
	tshirt.SetDiscount(50)
	o.Print()

	fmt.Printf("Сума електроніки: %.2f грн\n", o.TotalByType(order.Electronics))
	fmt.Printf("Сума одягу:       %.2f грн\n", o.TotalByType(order.Clothing))
	fmt.Printf("Сума їжі:         %.2f грн\n", o.TotalByType(order.Food))

	fmt.Println("\n--- Некоректні дані ---")
	if _, err := order.NewProduct("Зламаний", order.Electronics, -500); err != nil {
		fmt.Println("[Помилка]", err)
	}
	if err := tshirt.SetDiscount(150); err != nil {
		fmt.Println("[Помилка]", err)
	}

	fmt.Println("\n==================================================")
	fmt.Println("               ВИКОНАННЯ ЗАВЕРШЕНО                ")
	fmt.Println("==================================================")
}
/*


go test ./... -v -cover
go test lab5/wallet -v -cover
go test lab5/order -v -cover
*/ 
package wallet

import "fmt"

type CurrencyType string

const (
	UAH CurrencyType = "UAH"
	USD CurrencyType = "USD"
)

const ExchangeRate = 41.0

type Bill struct {
	Currency CurrencyType
	Amount   float64
}

func (b Bill) String() string {
	return fmt.Sprintf("%.2f %s", b.Amount, b.Currency)
}

type Wallet struct {
	Bills []Bill
}

func (w *Wallet) AddBill(currency CurrencyType, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("номінал купюри повинен бути більше 0")
	}
	w.Bills = append(w.Bills, Bill{Currency: currency, Amount: amount})
	fmt.Printf("[Гаманець] Додано: %.2f %s\n", amount, currency)
	return nil
}

func (w *Wallet) TotalByCurrency(currency CurrencyType) float64 {
	total := 0.0
	for _, bill := range w.Bills {
		if bill.Currency == currency {
			total += bill.Amount
		}
	}
	return total
}

func (w *Wallet) TotalInCurrency(target CurrencyType, rate float64) float64 {
	total := 0.0
	for _, bill := range w.Bills {
		if bill.Currency == target {
			total += bill.Amount
		} else {
			if target == UAH {
				total += bill.Amount * rate
			} else {
				total += bill.Amount / rate
			}
		}
	}
	return total
}

func (w *Wallet) Print() {
	fmt.Println("\n===== СТАН ГАМАНЦЯ =====")
	if len(w.Bills) == 0 {
		fmt.Println("Гаманець порожній")
		return
	}
	for i, bill := range w.Bills {
		fmt.Printf("  [%d] %s\n", i+1, bill)
	}
	fmt.Printf("  Всього UAH: %.2f\n", w.TotalByCurrency(UAH))
	fmt.Printf("  Всього USD: %.2f\n", w.TotalByCurrency(USD))
	fmt.Println("========================")
}